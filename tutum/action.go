package tutum

import "encoding/json"

type ActionListResponse struct {
	Objects []Action `json:"objects"`
}

type Action struct {
	Action       string `json:"action"`
	Body         string `json:"body"`
	End_date     string `json:"end_date"`
	Ip           string `json:"ip"`
	Location     string `json:"location"`
	Logs         string `json:"logs"`
	Method       string `json:"method"`
	Object       string `json:"object"`
	Path         string `json:"path"`
	Resource_uri string `json:"resource_uri"`
	Start_date   string `json:"start_date"`
	State        string `json:"state"`
	Uuid         string `json:"uuid"`
}

func ListActions() (ActionListResponse, error) {
	url := "action/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ActionListResponse

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

func GetAction(uuid string) (Action, error) {
	url := "action/" + uuid + "/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Action

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
