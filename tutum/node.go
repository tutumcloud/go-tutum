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

/*
func ListNodes
Returns : Array of Node objects
*/
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

/*
func GetNode
Argument : uuid
Returns : Node JSON object
*/
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

/*
func UpdateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Update(requestBody string) {

	url := "node/" + self.Uuid + "/"
	request := "PATCH"

	updatedNode := []byte(requestBody)

	_, err := TutumCall(url, request, updatedNode)
	if err != nil {
		panic(err)
	}
}

/*
func UpgradeDaemon
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Upgrade() {

	url := "node/" + self.Uuid + "/docker-upgrade/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}
}

/*
func TerminateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Terminate() {

	url := "node/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}
}
