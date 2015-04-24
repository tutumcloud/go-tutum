package tutum

import (
	"fmt"
	"os"

	"github.com/docker/docker/vendor/src/code.google.com/p/go.net/websocket"
)

func TutumStreamCall() error {

	var StreamUrl = "wss://stream.tutum.co:443/v1/events?token=" + ApiKey + "&user=" + User
	var origin = "http://localhost"

	if os.Getenv("TUTUM_STREAM_URL") != "" {
		StreamUrl = os.Getenv("TUTUM_STREAM_URL")
	}

	if !IsAuthenticated() {
		return fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
	}

	ws, err := websocket.Dial(StreamUrl, "", origin)
	if err != nil {
		return err
	}

	var msg = make([]byte, 512)
	var n int
	for {
		if n, err = ws.Read(msg); err != nil {
			return err
		}

		fmt.Printf("Received: %s.\n", msg[:n])
	}
	return nil
}
