package tutum

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
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
	var finalResponse CListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	finalResponse = response

Loop:
	for {
		if response.Meta.Next != "" {
			var nextResponse CListResponse
			data, err := TutumCall(response.Meta.Next[8:], request, body)
			if err != nil {
				return nextResponse, err
			}
			err = json.Unmarshal(data, &nextResponse)
			if err != nil {
				return nextResponse, err
			}
			finalResponse.Objects = append(finalResponse.Objects, nextResponse.Objects...)
			response = nextResponse

		} else {
			break Loop
		}
	}

	return finalResponse, nil
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

func (self *Container) Logs(c chan Logs) {

	endpoint := "container/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
	url := StreamUrl + endpoint

	header := http.Header{}
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(url, header)
	if err != nil {
		log.Println(err)
	}

	var msg Logs
	for {
		if err = ws.ReadJSON(&msg); err != nil {
			if err != nil && err.Error() != "EOF" {
				log.Println(err)
			} else {
				break
			}
		}
		c <- msg
	}
}

/*
func Exec
Arguments : the command to execute, a channel of type string for the output
*/

func (self *Container) Exec(command string, c chan Exec) {
	go self.Run(command, c)
Loop:
	for {
		select {
		case s := <-c:
			if s.Output != "EOF" {
				fmt.Printf("%s", s.Output)
			} else {
				break Loop
			}
		}
	}
}

func (self *Container) Run(command string, c chan Exec) {

	endpoint := "container/" + self.Uuid + "/exec/?user=" + User + "&token=" + ApiKey + "&command=" + url.QueryEscape(command)
	url := "wss://stream.tutum.co:443/v1/" + endpoint

	header := http.Header{}
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(url, header)
	if err != nil {
		log.Println(err)
	}

	var msg Exec
	for {
		if err = ws.ReadJSON(&msg); err != nil {
			if err != nil && err.Error() != "EOF" {
				log.Println(err)
			} else {
				break
			}
		}
		c <- msg
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
