package tutum

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	WRITE_WAIT = 5 * time.Second
	// Time allowed to read the next pong message from the peer.
	PONG_WAIT = 10 * time.Second
	// Send pings to client with this period. Must be less than PONG_WAIT.
	PING_PERIOD = PONG_WAIT / 2
)

/*
	func dial()
	Returns : a websocket connection
*/

func dial() (*websocket.Conn, error) {
	var Url = ""

	if os.Getenv("TUTUM_STREAM_HOST") != "" {
		u, _ := url.Parse(os.Getenv("TUTUM_STREAM_HOST"))
		_, port, _ := net.SplitHostPort(u.Host)
		if port == "" {
			u.Host = u.Host + ":443"
		}
		StreamUrl = u.Scheme + "://" + u.Host + "/v1/"
	} else {
		StreamUrl = "wss://stream.tutum.co:443/v1/"
	}

	log.Println(StreamUrl)

	if os.Getenv("TUTUM_AUTH") != "" {
		endpoint := ""
		endpoint = url.QueryEscape(os.Getenv("TUTUM_AUTH"))
		Url = StreamUrl + "events?auth=" + endpoint
	}
	if User != "" && ApiKey != "" {
		Url = StreamUrl + "events?token=" + ApiKey + "&user=" + User
		log.Println(Url)
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
			}
		} else {
			return ws
		}
	}
}

func messagesHandler(ws *websocket.Conn, ticker *time.Ticker, msg Event, c chan Event, e chan error) {
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(PONG_WAIT))
		return nil
	})
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("READ ERR")
			ticker.Stop()
			e <- err
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
	ticker := time.NewTicker(PING_PERIOD)
	ws := dialHandler(e)

	defer func() {
		close(c)
		close(e)
		ws.Close()
	}()
	go messagesHandler(ws, ticker, msg, c, e)

Loop:
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				ticker.Stop()
				e <- err
				break Loop
			}
		case <-e:
			ticker.Stop()
		}
	}
}
