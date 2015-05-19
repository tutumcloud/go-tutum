package tutum

import (
	"net/url"
	"os"
	"reflect"
	"time"

	"code.google.com/p/go.net/websocket"
)

/*
	func dial()
	Returns : a websocket connection
*/
func dial() (*websocket.Conn, error) {
	var streamWebsocket *websocket.Conn

	var Url = ""

	if os.Getenv("TUTUM_STREAM_URL") != "" {
		u, _ := url.Parse(os.Getenv("TUTUM_STREAM_URL"))
		u.Host = u.Host + ":443"
		StreamUrl = u.Scheme + "://" + u.Host + u.Path
	}

	if os.Getenv("TUTUM_AUTH") != "" {
		endpoint := ""
		endpoint = url.QueryEscape(os.Getenv("TUTUM_AUTH"))
		Url = StreamUrl + "events?auth=" + endpoint
	}
	if User != "" && ApiKey != "" {
		Url = StreamUrl + "events?token=" + ApiKey + "&user=" + User
	}

	var origin = "http://localhost"
	ws, err := websocket.Dial(Url, "", origin)
	if err != nil {
		return nil, err
	} else {
		streamWebsocket = ws
	}
	return streamWebsocket, nil
}

/*
	func TutumStreamCall
	Returns : The stream of all events from your NodeClusters, Containers, Services, Stack, Actions, ...
*/
func TutumEvents(c chan Event, e chan error) {
	var ws *websocket.Conn
	tries := 0
Loop:
	for {
		webSocket, err := dial()
		if err != nil {
			tries++
			time.Sleep(3 * time.Second)
			if tries > 3 {
				e <- err
				return
			}
		} else {
			ws = webSocket
			break Loop
		}
	}

	defer close(c)
	defer close(e)

	var msg Event
Loop2:
	for {
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			e <- err
			break Loop2
		}
		if reflect.TypeOf(msg).String() == "tutum.Event" {
			c <- msg
		}
	}
	ws.Close()
}
