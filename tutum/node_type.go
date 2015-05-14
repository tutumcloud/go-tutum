package tutum

import "encoding/json"

/*
func ListNodeTypes
Returns : Array of NodeType objects
*/
func ListNodeTypes() (NodeTypeListResponse, error) {

	url := "nodetype/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeTypeListResponse

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
func GetNodeType
Argument : provider name and type name
Returns : NodeType JSON object
*/
func GetNodeType(provider string, name string) (NodeType, error) {
	url := "nodetype/" + provider + "/" + name + "/"
	request := "GET"
	body := []byte(`{}`)
	var response NodeType

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
