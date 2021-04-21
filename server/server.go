package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var pool = GetPoolInstance()

func pushNotification(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if hasKey(query) {
		log.Println(illegalArgument)
		w.Write(json2byte(illegalArgument))
		return
	}

	uniqueKey := query.Get("key")

	notificationMessage := Notification{
		Title:       "Golang Push Server Test",
		Description: "First push notification try"}

	if conn, err := pool.get(uniqueKey); err == nil {
		log.Println(success, notificationMessage)
		conn.WriteJSON(notificationMessage)
		w.Write(json2byte(success))
		return
	}

	log.Println(failed)
	w.Write(json2byte(failed))
}

func handshake(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{} {

	if _, err := pool.get(uniqueKey); err == nil {
		return alreadyConnected
	}

	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return failed
	}

	pool.add(uniqueKey, connection)

	return success
}

func closeConnection(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if hasKey(query) {
		w.Write(json2byte(illegalArgument))
	}

	uniqueKey := query.Get("key")

	if connection, err := pool.get(uniqueKey); err != nil {
		connection.Close()
		w.Write(json2byte(success))
	}
}

func middleware(callback customEndpointHandler) endpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if hasKey(query) {
			w.Write(json2byte(illegalArgument))
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
