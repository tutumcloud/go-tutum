package tutum

import (
	"encoding/json"
	"time"
)

type AListResponse struct {
	Objects []Application `json:"objects"`
}

type Application struct {
	Autodestroy            string    `json:"autodestroy"`
	Autoreplace            string    `json:"autoreplace"`
	Autorestart            string    `json:"autorestart"`
	Container_ports        []ACPInfo `json:"container_ports"`
	Container_size         string    `json:"container_size"`
	Current_num_containers int       `json:"current_num_containers"`
	Deployed_datetime      time.Time    `json:"deployed_datetime"`
	Destroyed_datetime     time.Time     `json:"destroyed_datetime"`
	Entrypoint             string    `json:"entrypoint"`
	Image_name             string    `json:"image_name"`
	Image_tag              string    `json:"image_tag"`
	Name                   string    `json:"name"`
	Sequential_deployment  bool      `json:"sequential_deployment"`
	Public_dns             string    `json:"public_dns"`
	Resource_uri           string    `json:"resource_uri"`
	Run_command            string    `json:"run_command"`
	Running_num_containers int       `json:"running_num_containers"`
	Started_datetime       time.Time     `json:"started_datetime"`
	State                  string    `json:"state"`
	Stopped_datetime       time.Time     `json:"stopped_datetime"`
	Stopped_num_containers int       `json:"stopped_num_containers"`
	Target_num_containers  int       `json:"target_num_containers"`
	Unique_name            string    `json:"unique_name"`
	Uuid                   string    `json:"uuid"`
	Web_public_dns         string    `json:"web_public_dns"`
}

type ACPInfo struct {
	Application string `json:"application"`
	Inner_port  int    `json:"inner_port"`
	Outer_port  int    `json:"outer_port"`
	Protocol    string `json:"protocol"`
}

func ListApplications() ([]Application, error) {

	url := "application/"
	request := "GET"
	var response AListResponse
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
