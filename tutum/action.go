package tutum

import "encoding/json"

/*
func ListActions
Returns : Array of Action objects
*/
func ListActions() (ActionListResponse, error) {
	url := "action/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ActionListResponse

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
func GetAction
Argument : uuid
Returns : Action JSON object
*/
func GetAction(uuid string) (Action, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "action/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Action

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
