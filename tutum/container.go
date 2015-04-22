package tutum

import (
	"encoding/json"
	"fmt"
)

type CListResponse struct {
	Objects []Container `json:"objects"`
}

type Container struct {
	Application            string    `json:"application"`
	Autodestroy            string    `json:"autodestroy"`
	Autoreplace            string    `json:"autoreplace"`
	Autorestart            string    `json:"autorestart"`
	Container_ports        []CCPInfo `json:"container_ports"`
	Container_size         string    `json:"container_size"`
	Current_num_containers int       `json:"current_num_containers"`
	Deployed_datetime      string    `json:"deployed_datetime"`
	Destroyed_datetime     string    `json:"destroyed_datetime"`
	Entrypoint             string    `json:"entrypoint"`
	Exit_code              int       `json:"exit_code"`
	Exit_code_message      string    `json:"exit_code_message"`
	Image_name             string    `json:"image_name"`
	Image_tag              string    `json:"image_tag"`
	Name                   string    `json:"name"`
	Public_dns             string    `json:"public_dns"`
	Resource_uri           string    `json:"resource_uri"`
	Run_command            string    `json:"run_command"`
	Started_datetime       string    `json:"started_datetime"`
	State                  string    `json:"state"`
	Stopped_datetime       string    `json:"stopped_datetime"`
	Unique_name            string    `json:"unique_name"`
	Uuid                   string    `json:"uuid"`
}

type CCPInfo struct {
	Container  string `json:"container"`
	Inner_port int    `json:"inner_port"`
	Outer_port int    `json:"outer_port"`
	Protocol   string `json:"protocol"`
}

func ListContainers() ([]Container, error) {

	url := "container/"
	request := "GET"

	var response CListResponse
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

func GetContainer(uuid string) (Container, error) {

	url := "container/" + uuid + "/"
	request := "GET"

	var response Container

	data, err := TutumCall(url, request)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response, nil

}

func GetContainerLogs(uuid string) (string, error) {

	url := "container/" + uuid + "/logs/"
	request := "GET"

	data, err := TutumCall(url, request)
	if err != nil {
		return "", err
	}

	s := string(data)

	return s, nil

}

func StartContainer(uuid string) {

}

func StopContainer(uuid string) ([]byte, error) {

	url := "container/" + uuid + "/stop/"
	request := "POST"

	data, err := TutumCall(url, request)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func RedeployContainer(uuid string) {

}
