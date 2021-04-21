package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type baseResponse struct {
	code    int
	message string
}

var connections = make(map[string]*websocket.Conn)

var (
	success          = baseResponse{code: 0, message: "OK"}
	alreadyConnected = baseResponse{code: 204, message: "Handshake already completed"}
	illegalArgument  = baseResponse{code: 403, message: "Illeagal argument exceptiono"}
	failed           = baseResponse{code: 404, message: "Unknown Error"}
)

func validateKey(values url.Values) bool {
	return values.Get("key") == ""
}

func sendSampleNotification(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if validateKey(query) {
		log.Println(saveConvertJSON(illegalArgument))
		w.Write(saveConvertJSON(illegalArgument))
		return
	}

	uniqueKey := query.Get("key")
	notificationMessage := `{"message": "Sample notification"}`
	if potentialNotificationMessage := query.Get("notificationMessage"); potentialNotificationMessage != "" {
		notificationMessage = potentialNotificationMessage
	}

	if conn, ok := connections[uniqueKey]; ok {
		log.Println(saveConvertJSON(success))
		log.Println(notificationMessage)
		conn.WriteJSON(notificationMessage)
		w.Write(saveConvertJSON(success))
		return
	}

	log.Println(saveConvertJSON(failed))
	w.Write(saveConvertJSON(failed))
}

func saveConvertJSON(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}

func handshake(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{} {

	if _, ok := connections[uniqueKey]; ok {
		return alreadyConnected
	}

	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return failed
	}

	connections[uniqueKey] = connection

	return success
}

func closeConnection(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if validateKey(query) {
		w.Write(saveConvertJSON(illegalArgument))
	}

	uniqueKey := query.Get("key")

	if connection, ok := connections[uniqueKey]; ok {
		connection.Close()
		w.Write(saveConvertJSON(success))
	}
}

type customEndpointHandler func(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{}
type endpointHandler func(w http.ResponseWriter, r *http.Request)

func middleware(callback customEndpointHandler) endpointHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if validateKey(query) {
			w.Write(saveConvertJSON(illegalArgument))
		}

		json := callback(w, r, query.Get("key"))
		log.Println(json)
		// w.Write(saveConvertJSON(json))
	}
}

func main() {
	log.SetFlags(0)
	http.HandleFunc("/handshake", middleware(handshake))
	http.HandleFunc("/notification", sendSampleNotification)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
