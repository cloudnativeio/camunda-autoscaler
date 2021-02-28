package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "http://localhost:8080/engine-rest"
	path    = "history/process-instance/count"
)

func GetProcess() (interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", baseURL, path), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating the request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autoscaler")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returns: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error handling the payload: %s", err)
	}

	return data["count"], nil
}
