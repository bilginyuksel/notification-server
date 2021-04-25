package main

import (
	"flag"
	"log"

	"github.com/bilginyuksel/notification-server/server/endpoint"
)

var (
	debug    = false
	brokers  = ""
	version  = ""
	topics   = ""
	group    = ""
	assignor = ""
	oldest   = false
	verbose  = false
)

func main() {
	log.SetFlags(0)
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.StringVar(&brokers, "brokers", "localhost:9092", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "example", "Kafka consumer group definition")
	flag.StringVar(&version, "version", "2.1.1", "Kafka cluster version")
	flag.StringVar(&topics, "topics", "sarama", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if debug {
		go setNotificationConsumer()
	}

	endpoint.StartServer()
}
