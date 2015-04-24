package tutum

type NodeListResponse struct {
	Objects []Node `json:"objects"`
}

type Node struct {
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime"`
	Docker_version     string    `json:"docker_version"`
	Last_seen          string    `json:"last_seen"`
	Node_cluster       string    `json:"node_cluster"`
	Public_ip          string    `json:"public_ip"`
	Region             string    `json:"region"`
	Resource_uri       string    `json:"resource_uri"`
	State              string    `json:"state"`
	Tags               []NodeTag `json:"tags"`
	Uuid               string    `json:"uuid"`
}

type NodeTag struct {
	Name string `json:"name"`
}
