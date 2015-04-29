package tutum

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

/*
	func dial()
	Returns : a websocket connection
*/
func dial() (websocket.Conn, error) {
	var StreamUrl = "wss://stream.tutum.co:443/v1/events?token=" + ApiKey + "&user=" + User
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
func TutumStreamCall(c chan string) error {

	ws, err := dial()

	var msg = make([]byte, 512)
	var n int
	for {
		if n, err = ws.Read(msg); err != nil {
			return err
		}

		c <- fmt.Sprintf("%s", msg[:n])
	}
}

func OnEvent(f func()) {
	c := make(chan string)
	go TutumStreamCall(c)
	for {
		select {
		case stream := <-c:
			fmt.Println(stream)
			f()
		}
	}
}

func Stream() {
	c := make(chan string)
	go TutumStreamCall(c)
	for {
		select {
		case stream := <-c:
			fmt.Println(stream)
		}
	}
}
