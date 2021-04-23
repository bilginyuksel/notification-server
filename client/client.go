package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	hostAddress   = "localhost:8888"
	handshakePath = "/handshake"
	scheme        = "ws"
)

func main() {
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	uuid, _ := uuid.NewUUID()

	u := url.URL{Scheme: scheme,
		Host:     hostAddress,
		Path:     handshakePath,
		RawQuery: fmt.Sprintf("clientId=" + uuid.String())}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}

	defer c.Close()

	for {
		_, bytes, _ := c.ReadMessage()
		log.Println(string(bytes))
	}
}
