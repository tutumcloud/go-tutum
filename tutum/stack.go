package tutum

import "encoding/json"

type StackListResponse struct {
	Objects []StackShort `json: "objects"`
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

func ListStacks() ([]StackShort, error) {
	url := "stack/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response StackListResponse

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
