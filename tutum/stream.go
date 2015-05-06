package tutum

import (
	"encoding/json"
	"net/url"
	"os"
	"reflect"

	"golang.org/x/net/websocket"
)

type Event struct {
	Type         string   `json:"type"`
	Action       string   `json:"action"`
	Parents      []string `json:"parents"`
	Resource_uri string   `json:"resource_uri"`
	State        string   `json:"state"`
}

/*
	func dial()
	Returns : a websocket connection
*/
func dial() (websocket.Conn, error) {
	var StreamUrl = ""
	if os.Getenv("TUTUM_AUTH") != "" {
		endpoint := url.QueryEscape(os.Getenv("TUTUM_AUTH"))
		StreamUrl = "wss://stream.tutum.co:443/v1/events?auth=" + endpoint
	}
	if User != "" && ApiKey != "" {
		StreamUrl = "wss://stream.tutum.co:443/v1/events?token=" + ApiKey + "&user=" + User
	}
	var origin = "http://localhost"
	ws, err := websocket.Dial(StreamUrl, "", origin)

	if os.Getenv("TUTUM_STREAM_URL") != "" {
		StreamUrl = os.Getenv("TUTUM_STREAM_URL")
	}

	if err != nil {
		dial()
	}
	return *ws, nil
}

/*
	func TutumStreamCall
	Returns : The stream of all events from your NodeClusters, Containers, Services, Stack, Actions, ...
*/
func TutumEvents(c chan Event) (Event, error) {

	ws, err := dial()

	var msg = make([]byte, 512)
	var event Event
	var n int
	for {
		if n, err = ws.Read(msg); err != nil {
			return event, err
		}
		err := json.Unmarshal(msg[:n], &event)
		if err != nil {
			return event, err
		}
		if reflect.TypeOf(event).String() == "tutum.Event" {
			c <- event
		}
	}
}
