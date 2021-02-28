package autoscaler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/trx35479/camunda-autoscaler/autoscaler/apis"
	"github.com/trx35479/camunda-autoscaler/autoscaler/log"
)

type Deploy struct {
	Spec interface{} `json:"spec"`
}

type Specification struct {
	Replicas int `json:"replicas"`
}

const (
	ServiceAccountPath = "/var/run/secrets/kubernetes.io/serviceaccount"
)

var (
	kubernetesServiceHost = os.Getenv("KUBERNETES_SERVICE_HOST")
	logger                = log.NewLogger()
	name                  = os.Getenv("BPM_DEPLOYMENT_NAME")
)

func Handler() error {
	count := apis.GetProcess()

	logger.Printf("count: %v", count)

	// Let's get the replica here
	serviceAcctToken, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", ServiceAccountPath, "token"))
	if err != nil {
		return fmt.Errorf("cannot read kubernetes token: %v", err)
	}
	namespace, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", ServiceAccountPath, "namespace"))
	if err != nil {
		return fmt.Errorf("cannot read kuberneres namespace: %v", err)
	}
	cacrt, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", ServiceAccountPath, "ca.crt"))
	if err != nil {
		return fmt.Errorf("cannot read ca certificate: %v", err)
	}

	client := &apis.Client{
		Certificate: cacrt,
	}

	replicas, err := client.GetReplica(string(serviceAcctToken), kubernetesServiceHost, string(namespace), name)
	if err != nil {
		return err
	}

	if count >= 50 {
		if replicas.(int) < 4 {
			number := replicas.(int) + 1
			spec := &Deploy{
				Spec: Specification{Replicas: number},
			}
			payload, _ := json.Marshal(spec)
			logger.Printf("scaling replica to %d: ", number)
			scale, err := client.SetReplica(string(serviceAcctToken), kubernetesServiceHost, string(namespace), name, payload)
			if err != nil {
				return err
			}
			logger.Printf("set replicas to: %d", scale)
		} else {
			logger.Printf("not scaling up as the current replica is equal to: %v", replicas)
		}

	}

	if count <= 20 {
		if replicas.(int) > 1 {
			number := replicas.(int) - 1
			spec := &Deploy{
				Spec: Specification{Replicas: number},
			}
			payload, _ := json.Marshal(spec)
			logger.Printf("scaling replica to %d: ", number)
			scale, err := client.SetReplica(string(serviceAcctToken), kubernetesServiceHost, string(namespace), name, payload)
			if err != nil {
				return err
			}
			logger.Printf("set replicas to: %d", scale)
		}
	} else {
		logger.Printf("not scaling down as the process count is more that: %v", count)
	}

	return nil
}
