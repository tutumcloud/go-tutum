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

func ListNodes() (NodeListResponse, error) {

	url := "node/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

func GetNode(uuid string) (Node, error) {
	url := "node/" + uuid + "/"
	request := "GET"
	body := []byte(`{}`)
	var response Node

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

func UpdateNode(uuid string, updatedNode []byte) (Node, error) {

	url := "node/" + uuid + "/"
	request := "PATCH"

	var response Node

	data, err := TutumCall(url, request, updatedNode)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

func UpgradeDaemon(uuid string) (Node, error) {

	url := "node/" + uuid + "/docker-upgrade/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Node

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

func TerminateNode(uuid string) (Node, error) {

	url := "node/" + uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)
	var response Node

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil

}
