package tutum

import "encoding/json"

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
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func GetNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func GetNodeCluster(uuid string) (NodeCluster, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:] + "/"
	} else {
		url = "nodecluster/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response NodeCluster

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func CreateNodeCluster
Argument : NodeCluster JSON object (see documentation)
Returns : NodeCluster JSON object
*/
func CreateNodeCluster(createRequest NodeCreateRequest) (NodeCluster, error) {

	url := "nodecluster/"
	request := "POST"
	var response NodeCluster

	newCluster, err := json.Marshal(createRequest)
	if err != nil {
		return response, err
	}

	data, err := TutumCall(url, request, newCluster)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func DeployNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Deploy() error {

	url := "nodecluster/" + self.Uuid + "/deploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func UpdateNodeCluster
Argument : uuid and nodecluster JSON object (see documentation)
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Update(createRequest NodeCreateRequest) error {

	url := "nodecluster/" + self.Uuid + "/"
	request := "PATCH"

	updatedNodeCluster, err := json.Marshal(createRequest)
	if err != nil {
		return err
	}

	_, errr := TutumCall(url, request, updatedNodeCluster)
	if errr != nil {
		return errr
	}

	return nil
}

/*
func UpgradeNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Upgrade() error {

	url := "nodecluster/" + self.Uuid + "/docker-upgrade/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func TerminateNodeCluster
Argument : uuid
Returns : NodeCluster JSON object
*/
func (self *NodeCluster) Terminate() error {

	url := "nodecluster/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
