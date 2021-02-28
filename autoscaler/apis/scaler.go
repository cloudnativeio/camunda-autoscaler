package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetReplica gets the number of deployment replica
func (c *Client) GetReplica(token, kubeHost, ns, name string) (interface{}, error) {
	client, err := c.client("kube: getReplica")
	if err != nil {
		return nil, fmt.Errorf("error instantiating client: %s", err)
	}

	requestUrl := fmt.Sprintf("https://%s/apis/apps/v1/namespaces/%s/deployments/%s", kubeHost, ns, name)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to construct request to kubernetes api: %s", requestUrl)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("User-Agent", "autoscaler")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to kubernetes api: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body")
	}
	var deployment *Deployment
	err = json.Unmarshal(body, &deployment)
	if err != nil {
		return nil, fmt.Errorf("error handling the payload")
	}

	return deployment.Spec.Replicas, nil
}

// SetReplica patches deployment and set the desired replica
func (c *Client) SetReplica(token, kubeHost, ns, name string, payload []byte) (interface{}, error) {
	client, err := c.client("kube: patchReplica")
	if err != nil {
		return nil, fmt.Errorf("error instantiating client: %s", err)
	}
	requestUrl := fmt.Sprintf("https://%s/apis/apps/v1/namespaces/%s/deployments/%s", kubeHost, ns, name)

	req, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to construct request to kubernetes api: %s", requestUrl)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/strategic-merge-patch+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("User-Agent", "autoscaler")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to kubernetes api: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body")
	}

	var deployment *Deployment
	err = json.Unmarshal(body, &deployment)
	if err != nil {
		return nil, fmt.Errorf("error handling the payload")
	}

	return deployment.Spec.Replicas, nil
}
