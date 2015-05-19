package tutum

import (
	"encoding/json"
	"log"

	"code.google.com/p/go.net/websocket"
)

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
Argument : a channel of type string for the output
*/
func (self *Service) Logs(c chan string) {

	endpoint := "service/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
	origin := "http://localhost/"
	url := "wss://stream.tutum.co:443/v1/" + endpoint
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	for {
		ws.Request()
		if _, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		c <- string(msg)
	}
}

/*
func UpdateService
Argument : updatedService JSON object
Returns : Error
*/
func (self *Service) Scale() error {

	url := "service/" + self.Uuid + "/scale/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Service

	data, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return nil
}

/*
func UpdateService
Argument : updatedService JSON object
Returns : Error
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
Returns : Error
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
Returns : Error
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
Returns : Error
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
Returns : Error
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
