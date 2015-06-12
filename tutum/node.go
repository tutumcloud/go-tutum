package tutum

import (
	"encoding/json"
	"log"

	"code.google.com/p/go.net/websocket"
)

/*
func ListNodes
Returns : Array of Node objects
*/
func ListNodes() (NodeListResponse, error) {

	url := "node/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response NodeListResponse

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
func GetNode
Argument : uuid
Returns : Node JSON object
*/
func GetNode(uuid string) (Node, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "node/" + uuid + "/"
	}

	request := "GET"
	body := []byte(`{}`)
	var response Node

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
func UpdateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Update(createRequest Node) error {

	url := "node/" + self.Uuid + "/"
	request := "PATCH"

	updatedNode, err := json.Marshal(createRequest)
	if err != nil {
		return err
	}

	_, errr := TutumCall(url, request, updatedNode)
	if err != nil {
		return errr
	}

	return nil
}

/*
func UpgradeDaemon
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Upgrade() error {

	url := "node/" + self.Uuid + "/docker-upgrade/"
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
func TerminateNode
Argument : uuid
Returns : Node JSON object
*/
func (self *Node) Terminate() error {

	url := "node/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Node) Events(c chan NodeEvent) {
	endpoint := "node/" + self.Uuid + "/events/?user=" + User + "&token=" + ApiKey
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
	var response NodeEvent

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
