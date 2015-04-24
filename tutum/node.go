package tutum

import "encoding/json"

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

func ListNodes() ([]Node, error) {
	url := "node/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Objects, nil
}
