package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (c *Client)GetProcess() (interface{}, error){
	url := os.Getenv("URL")
	client, err := c.client("Getting process instance: ")
	if err != nil {
		return nil, errors.New("failed to instantiate client")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url, "engine-rest/history/process-instance/count"), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating the request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

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