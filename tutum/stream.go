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

var StreamUrl = ""

/*
	func dial()
	Returns : a websocket connection
*/
func dial() (websocket.Conn, error) {

	if os.Getenv("TUTUM_STREAM_URL") != "" {
		u, err := url.Parse(os.Getenv("TUTUM_STREAM_URL"))
		if err != nil {
			panic(err)
		}
		u.Host = u.Host + ":443"
		StreamUrl = u.Scheme + "://" + u.Host + u.Path
	}

	if os.Getenv("TUTUM_AUTH") != "" {
		endpoint := url.QueryEscape(os.Getenv("TUTUM_AUTH"))

		StreamUrl = StreamUrl + "events?auth=" + endpoint
	}
	if User != "" && ApiKey != "" {
		StreamUrl = StreamUrl + "events?token=" + ApiKey + "&user=" + User
	}

	var origin = "http://localhost"
	ws, err := websocket.Dial(StreamUrl, "", origin)

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
