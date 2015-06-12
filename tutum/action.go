package tutum

import (
	"encoding/json"
	"log"

	"code.google.com/p/go.net/websocket"
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
	origin := "http://localhost/"
	url := StreamUrl + endpoint

	config, err := websocket.NewConfig(url, origin)
	if err != nil {
		log.Println(err)
	}

	config.Header.Add("User-Agent", customUserAgent)

	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Println(err)
	}
	var response Logs

	var n int
	var msg = make([]byte, 1024)
	for {
		if n, err = ws.Read(msg); err != nil {
			if err != nil && err.Error() != "EOF" {
				log.Println(err)
			} else {
				break
			}
		}
		err = json.Unmarshal(msg[:n], &response)
		if err != nil {
			log.Println(err)
		}

		c <- response
	}
}
