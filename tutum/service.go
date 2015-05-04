package tutum

import "encoding/json"

type SListResponse struct {
	Objects []Service `json: "objects"`
}

type Service struct {
	Autodestroy            string       `json:"autodestroy`
	Autoredeploy           bool         `json:"autoredeploy`
	Autorestart            string       `json:"autorestart`
	Containers             []string     `json:"containers`
	Container_ports        []SCPInfo    `json:"container_ports`
	Container_size         string       `json:"container_size`
	Current_num_containers int          `json:"current_num_containers`
	Deployed_datetime      string       `json:"deployed_datetime`
	Destroyed_datetime     string       `json:"destroyed_datetime`
	Entrypoint             string       `json:"entrypoint`
	Exit_code              int          `json:"exit_code`
	Exit_code_message      string       `json:"exit_code_message`
	Image_name             string       `json:"image_name"`
	Image_tag              string       `json:"image_tag`
	Linked_to_service      []LinkToInfo `json:"linked_to_service`
	Name                   string       `json:"name"`
	Public_dns             string       `json:"public_dns"`
	Resource_uri           string       `json:"resource_uri"`
	Run_command            string       `json:"run_command"`
	Started_datetime       string       `json:"started_datetime"`
	State                  string       `json:"state"`
	Stack                  string       `json:"stack"`
	Stopped_datetime       string       `json:"stopped_datetime"`
	Target_num_containers  int          `json:"target_num_containers"`
	Unique_name            string       `json:"unique_name"`
	Uuid                   string       `json:"uuid"`
}

type SCPInfo struct {
	Container  string `json:"container"`
	Inner_port int    `json:"inner_port"`
	Outer_port int    `json:"outer_port"`
	Protocol   string `json:"protocol"`
}

//Basic information from linked services
type LinkToInfo struct {
	From_service string `json:"from_service"`
	Name         string `json:"name"`
	To_service   string `json:"to_service"`
}

/*
func ListServices
Returns : Array of Service objects
*/
func ListServices() (SListResponse, error) {
	url := "service/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response SListResponse

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
func GetService
Argument : uuid
Returns : Service JSON object
*/
func GetService(uuid string) (Service, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "service/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Service

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
func CreateService
Argument : Service JSON object (see documentation)
Returns : Service JSON object
*/
func CreateService(requestBody string) (Service, error) {

	url := "service/"
	request := "POST"

	newService := []byte(requestBody)
	var response Service

	data, err := TutumCall(url, request, newService)
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
func GetServiceLogs
Argument : uuid
Returns : A string containing the logs of the service
*/
func (self *Service) Logs() (string, error) {

	url := "service/" + self.Uuid + "/logs/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)

	data, err := TutumCall(url, request, body)
	if err != nil {
		return "", err
	}

	s := string(data)

	return s, nil
}

/*
func UpdateService
Argument : uuid, updatedService JSON object
Returns : Service JSON object
*/
func (self *Service) Update(requestBody string) error {

	url := "service/" + self.Uuid + "/"
	request := "PATCH"

	updatedService := []byte(requestBody)

	_, err := TutumCall(url, request, updatedService)
	if err != nil {
		return err
	}

	return nil
}

/*
func StartService
Argument : uuid
Returns : Service JSON object
*/
func (self *Service) Start() error {
	url := "service/" + self.Uuid + "/start/"
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
func StopService
Argument : uuid
Returns : Service JSON object
*/
func (self *Service) StopService() error {

	url := "service/" + self.Uuid + "/stop/"
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
func RedeployService
Argument : uuid
Returns : Service JSON object
*/
func (self *Service) Redeploy() error {

	url := "service/" + self.Uuid + "/redeploy/"
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
func TerminateService
Argument : uuid
Returns : Service JSON object
*/
func (self *Service) TerminateService() error {
	url := "service/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
