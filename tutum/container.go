package tutum

import (
	"encoding/json"
	"log"

	"code.google.com/p/go.net/websocket"
)

/*
func ListContainers
Returns : Array of Container objects
*/
func ListContainers() (CListResponse, error) {

	url := "container/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response CListResponse
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
func GetContainer
Argument : uuid
Returns : Container JSON object
*/
func GetContainer(uuid string) (Container, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "container/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

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
func GetContainerLogs
Argument : uuid
Returns : A string containing the logs of the container
*/
func (self *Container) Logs(c chan string) {

	endpoint := "container/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
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
func StartContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Start() error {

	url := "container/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

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
func StopContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Stop() error {

	url := "container/" + self.Uuid + "/stop/"
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
func RedeployContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Redeploy() error {

	url := "container/" + self.Uuid + "/redeploy/"
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
func TerminateContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Terminate() error {

	url := "container/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}
