package tutum

import (
	"net/url"
	"os"
	"reflect"

	"code.google.com/p/go.net/websocket"
)

/*
	func dial()
	Returns : a websocket connection
*/
func dial() (*websocket.Conn, error) {

	if os.Getenv("TUTUM_STREAM_URL") != "" {
		u, err := url.Parse(os.Getenv("TUTUM_STREAM_URL"))
		if err != nil {
			return nil, err
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

	return ws, nil
}

/*
	func TutumStreamCall
	Returns : The stream of all events from your NodeClusters, Containers, Services, Stack, Actions, ...
*/
func TutumEvents(c chan Event, e chan error) {
	ws, err := dial()
	if err != nil {
		e <- err
	}

	defer ws.Close()

	defer close(c)
	defer close(e)

	var msg Event
	for {

		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			e <- err
			return
		}
		if reflect.TypeOf(msg).String() == "tutum.Event" {
			c <- msg
		}

	}
}
