package tutum

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
func ListActions
Returns : Array of Action objects
*/
func ListActions() (ActionListResponse, error) {
	url := "action/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response ActionListResponse

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
func GetAction
Argument : uuid
Returns : Action JSON object
*/
func GetAction(uuid string) (Action, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "action/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Action

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
func GetLogs
Argument : a channel of type string for the output
*/

func (self *Action) GetLogs(c chan Logs) {
	endpoint := "action/" + self.Uuid + "/logs/?user=" + User + "&token=" + ApiKey
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
