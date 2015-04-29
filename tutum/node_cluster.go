package tutum

import (
	"encoding/json"
	"fmt"
)

type NodeClusterListResponse struct {
	Objects []NodeCluster `json:"objects"`
}

type NodeCluster struct {
	Current_num_nodes  int       `json:"current_num_nodes"`
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime"`
	Disk               int       `json:"disk"`
	Nodes              []string  `json:"nodes"`
	Region             string    `json:"region"`
	Resource_uri       string    `json:"resource_uri"`
	State              string    `json:"state"`
	Tags               []NodeTag `json:"tags"`
	Target_num_nodes   int       `json:"target_num_nodes"`
	Uuid               string    `json:"uuid"`
}

/*
func ListNodeClusters
Returns : Array of NodeCluster objects
*/
func ListNodeClusters() (NodeClusterListResponse, error) {

	url := "nodecluster/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeClusterListResponse

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
func GetNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func GetNodeCluster(uuid string) (NodeCluster, error) {

	url := "nodecluster/" + uuid + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response NodeCluster

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
func CreateNodeCluster
Argument : NodeCluster JSON object (see documentation)
Returns : NodeCluster JSON object
*/
func CreateNodeCluster(requestBody string) (NodeCluster, error) {

	url := "nodecluster/"
	request := "POST"

	newCluster := []byte(requestBody)
	var response NodeCluster

	data, err := TutumCall(url, request, newCluster)
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
func DeployNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Deploy() {

	url := "nodecluster/" + self.Uuid + "/deploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}
}

/*
func UpdateNodeCluster
Argument : uuid and nodecluster JSON object (see documentation)
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Update(requestBody string) {

	url := "nodecluster/" + self.Uuid + "/"
	request := "PATCH"

	updatedNodeCluster := []byte(requestBody)
	fmt.Println(string(updatedNodeCluster))

	_, err := TutumCall(url, request, updatedNodeCluster)
	if err != nil {
		panic(err)
	}
}

/*
func UpgradeNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Upgrade() {

	url := "nodecluster/" + self.Uuid + "/docker-upgrade/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}
}

/*
func TerminateNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Terminate() {

	url := "nodecluster/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}
}
