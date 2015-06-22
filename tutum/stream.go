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

	config, err := websocket.NewConfig(Url, origin)
	if err != nil {
		return nil, err
	}

	config.Header.Add("User-Agent", customUserAgent)

	ws, err := websocket.DialConfig(config)
	if err != nil {
		return nil, err
	} else {
		streamWebsocket = ws
	}
	return streamWebsocket, nil
}

func dialHandler(e chan error) *websocket.Conn {
	tries := 0
	for {
		ws, err := dial()
		if err != nil {
			tries++
			time.Sleep(3 * time.Second)
			if tries > 3 {
				e <- err
				return nil
			}
		} else {
			return ws
		}
	}
}

func messagesHandler(ws *websocket.Conn, msg Event, c chan Event, e chan error) {
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

/*
	func TutumStreamCall
	Returns : The stream of all events from your NodeClusters, Containers, Services, Stack, Actions, ...
*/
func TutumEvents(c chan Event, e chan error) {

	defer close(c)
	defer close(e)

	var msg Event
	ws := dialHandler(e)
	messagesHandler(ws, msg, c, e)
	ws.Close()
}
