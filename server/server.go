package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	conn "github.com/bilginyuksel/notification-server/server/conn"
	dto "github.com/bilginyuksel/notification-server/server/dto"
	entity "github.com/bilginyuksel/notification-server/server/entity"
)

var pool = conn.GetPoolInstance()

type (
	customEndpointHandler func(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{}

	endpointHandler func(w http.ResponseWriter, r *http.Request)
)

func pushNotification(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if hasKey(query) {
		log.Println(dto.IllegalArgument)
		w.Write(json2byte(dto.IllegalArgument))
		return
	}

	uniqueKey := query.Get("key")

	notificationMessage := entity.Notification{
		Title:       "Golang Push Server Test",
		Description: "First push notification try"}

	if err := pool.Write(uniqueKey, notificationMessage); err != nil {
		w.Write(json2byte(dto.Failed))
	}

	w.Write(json2byte(notificationMessage))

}

func handshake(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{} {

	if pool.Has(uniqueKey) {
		return dto.AlreadyConnected
	}

	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return dto.Failed
	}

	pool.Add(uniqueKey, connection)

	return dto.Success
}

func closeConnection(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if hasKey(query) {
		w.Write(json2byte(dto.IllegalArgument))
	}

	uniqueKey := query.Get("key")

	if pool.Has(uniqueKey) {
		pool.Close(uniqueKey)
		w.Write(json2byte(dto.Success))
	}
}

func middleware(callback customEndpointHandler) endpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if hasKey(query) {
			w.Write(json2byte(dto.IllegalArgument))
		}

		json := callback(w, r, query.Get("key"))
		log.Println(json)
		// w.Write(saveConvertJSON(json))
	}
}

func main() {
	log.SetFlags(0)
	http.HandleFunc("/handshake", middleware(handshake))
	http.HandleFunc("/notification", pushNotification)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
