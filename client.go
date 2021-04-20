package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

const (
	hostAddress   = "localhost:8888"
	handshakePath = "/handshake"
	scheme        = "ws"
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: scheme, Host: hostAddress, Path: handshakePath, RawQuery: "key=mykey&another=myAnother"}
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
