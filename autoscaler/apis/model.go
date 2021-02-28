package apis

type Deployment struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	Spec       struct {
		Replicas int `json:"replicas"`
	} `json:"spec"`
}
