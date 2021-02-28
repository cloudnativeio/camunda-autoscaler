package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	basePath  = "engine-rest"
	countPath = "history/process-instance/count"
)

func GetProcess() int {
	client := http.Client{}

	host := os.Getenv("CAMUNDA_SERVICE_SERVICE_HOST")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:8080/%s/%s", host, basePath, countPath), nil)
	if err != nil {
		logger.Fatalf("error creating the request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autoscaler")

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatalf("server error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Fatalf("server returns: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalf("error reading response body: %s", err)
	}

	var data map[string]int
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Fatalf("error handling the payload: %s", err)
	}

	return data["count"]
}