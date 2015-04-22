package tutum

import "encoding/json"

type SListResponse struct {
	Objects []Service `json: "objects"`
}

type Service struct {
	Autodestroy            string   `json:"autodestroy"`
	Autoredeploy           bool     `json:"autoredeploy"`
	Autorestart            string   `json:"autorestart"`
	Containers             []string `json:"containers"`
	Current_num_containers int      `json:"current_num_containers"`
	Uuid                   string   `json:"uuid"`
}

func ListServices() ([]Service, error) {
	url := "service/"
	request := "GET"

	var response SListResponse
	data, err := TutumCall(url, request)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Objects, nil

}
