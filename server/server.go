package main

import (
	"log"

	"github.com/bilginyuksel/notification-server/server/endpoint"
)

func main() {
	log.SetFlags(0)
	endpoint.StartServer()
}
