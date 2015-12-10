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
	LoadAuth()

	if os.Getenv("TUTUM_STREAM_HOST") != "" {
		u, _ := url.Parse(os.Getenv("TUTUM_STREAM_HOST"))
		_, port, _ := net.SplitHostPort(u.Host)
		if port == "" {
			u.Host = u.Host + ":443"
		}
		StreamUrl = u.Scheme + "://" + u.Host + "/v1/"
	} else if os.Getenv("TUTUM_STREAM_URL") != "" {
		u, _ := url.Parse(os.Getenv("TUTUM_STREAM_URL"))
		_, port, _ := net.SplitHostPort(u.Host)
		if port == "" {
			u.Host = u.Host + ":443"
		}
		StreamUrl = u.Scheme + "://" + u.Host + "/v1/"
	}

	Url := StreamUrl + "events/"

	header := http.Header{}
	header.Add("Authorization", AuthHeader)
	header.Add("User-Agent", customUserAgent)

	var Dialer websocket.Dialer
	ws, _, err := Dialer.Dial(Url, header)
	if err != nil {
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
				log.Println("[DIAL ERROR]: " + err.Error())
				e <- err
			}
		} else {
			return ws
		}
	}
}

func messagesHandler(ws *websocket.Conn, ticker *time.Ticker, msg Event, c chan Event, e chan error, e2 chan error) {
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(PONG_WAIT))
		return nil
	})
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			e <- err
			e2 <- err
			time.Sleep(4 * time.Second)
		} else {
			if reflect.TypeOf(msg).String() == "tutum.Event" {
				log.Println(msg)
				c <- msg
			}
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

	e2 := make(chan error)

	defer func() {
		close(c)
		close(e)
		ws.Close()
	}()
	go messagesHandler(ws, ticker, msg, c, e, e2)

Loop:
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				ticker.Stop()
				log.Println("Ping Timeout")
				e <- err
				break Loop
			}
		case <-e2:
			ticker.Stop()
		}
	}
}
