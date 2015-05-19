package tutum

import (
	"log"
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
	var streamWebsocket *websocket.Conn
Loop:
	for {
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
			return nil, err
		} else {
			streamWebsocket = ws
			break Loop
		}
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
	for {
		webSocket, err := dial()
		if err != nil {
			log.Print("Error")
			tries++
			if tries > 3 {
				log.Print("Returning")
				e <- err
				return
			}
		} else {
			ws = webSocket
			break
		}
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
