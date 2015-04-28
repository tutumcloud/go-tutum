package tutum

import "encoding/json"

type StackListResponse struct {
	Objects []StackShort `json:"objects"`
}

type StackShort struct {
	Deployed_datetime  string   `json:"deployed_datetime"`
	Destroyed_datetime string   `json:"destroyed_datetime"`
	Name               string   `json:"name"`
	Resource_uri       string   `json:"resource_uri"`
	Service            []string `json:"services"`
	State              string   `json:"state"`
	Synchronized       bool     `json:"synchronized"`
	Uuid               string   `json:"uuid"`
}

type Stack struct {
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime"`
	Name               string    `json:"name"`
	Resource_uri       string    `json:"resource_uri"`
	Service            []Service `json:"services"`
	State              string    `json:"state"`
	Synchronized       bool      `json:"synchronized"`
	Uuid               string    `json:"uuid"`
}

/*
func ListStacks
Returns : Array of Stack objects
*/
func ListStacks() (StackListResponse, error) {
	url := "stack/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response StackListResponse

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
func GetStack
Argument : uuid
Returns : Stack JSON object
*/
func GetStack(uuid string) (Stack, error) {

	url := "stack/" + uuid + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

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
func Export
Argument : uuid
Returns : String that contains the Stack details
*/
func ExportStack(uuid string) (string, error) {

	url := "stack/" + uuid + "/export/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	s := string(data)

	return s, nil
}

/*
func CreateStack
Argument : Stack JSON object (see documentation)
Returns : Stack JSON object
*/
func CreateStack(newStack []byte) (Stack, error) {

	url := "stack/"
	request := "POST"

	var response Stack

	data, err := TutumCall(url, request, newStack)
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
func UpdateStack
Argument : uuid and Stack JSON object (see documentation)
Returns : Stack JSON object
*/
func UpdateStack(uuid string, updatedStack []byte) (Stack, error) {

	url := "stack/" + uuid + "/"
	request := "PATCH"
	var response Stack

	data, err := TutumCall(url, request, updatedStack)
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
func StartStack
Argument : uuid
Returns : Stack JSON object
*/
func StartStack(uuid string) (Stack, error) {

	url := "stack/" + uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

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
func StopStack
Argument : uuid
Returns : Stack JSON object
*/
func StopStack(uuid string) (Stack, error) {

	url := "stack/" + uuid + "/stop/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

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
func RedeployStack
Argument : uuid
Returns : Stack JSON object
*/
func RedeployStack(uuid string) (Stack, error) {

	url := "stack/" + uuid + "/redeploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

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
func TerminateStack
Argument : uuid
Returns : Stack JSON object
*/
func TerminateStack(uuid string) (Stack, error) {

	url := "stack/" + uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

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
