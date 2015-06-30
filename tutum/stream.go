package tutum

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait = (5 * time.Second) / 2
)

/*
	func dial()
	Returns : a websocket connection
*/

func dial() (*websocket.Conn, error) {
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

	header := http.Header{}
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(Url, header)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ws, nil
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
		err := ws.ReadJSON(&msg)
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
	var msg Event

	ticker := time.NewTicker(5 * time.Second)
	ws := dialHandler(e)

	defer func() {
		close(c)
		close(e)
		ws.Close()
	}()

	go messagesHandler(ws, msg, c, e)

	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pongWait)); err != nil {
				e <- err
				return
			}
		}
	}
}
