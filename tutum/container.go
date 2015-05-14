package tutum

import (
	"encoding/json"
	"log"
	"net/url"

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
Argument : a channel of type string for the output
*/
func (self *Container) Logs(c chan string) {

	endpoint := "container/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
	origin := "http://localhost/"
	url := "wss://live-test.tutum.co:443/v1/" + endpoint
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Println(err)
	}

	msg := make([]byte, 512)

	for {
		ws.Request()
		if _, err = ws.Read(msg); err != nil {
			log.Println(err)
		}
		c <- string(msg)
	}
}

/*
func Exec
Arguments : the command to execute, a channel of type string for the output
*/

func (self *Container) Exec(command string, c chan string) {

	endpoint := "container/" + self.Uuid + "/exec/?user=" + User + "&token=" + ApiKey + "&command=" + url.QueryEscape(command)
	log.Println(endpoint)
	origin := "http://localhost/"
	url := "wss://live-test.tutum.co:443/v1/" + endpoint
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		if err.Error() != "EOF" {
			log.Println(err)
		}
	}

	msg := make([]byte, 1024)

	for {
		if _, err = ws.Read(msg); err != nil {
			if err.Error() == "EOF" {
				c <- err.Error()
			}
		}
		c <- string(msg)
	}
}

/*
func StartContainer
Returns : Error
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
Returns : Error
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
Returns : Error
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
Returns : Error
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
