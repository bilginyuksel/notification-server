package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	hostname string = "localhost"
	port     int    = 8888
)

func StartServerWithHostnameAndPort(hostname string, port int) {
	http.HandleFunc("/handshake", handshake)
	http.HandleFunc("/handshake/stop", handshakeStop)
	http.HandleFunc("/handshake/status", handshakeStatus)

	// send notification in debug mode
	http.HandleFunc("/notification", sendNotification)

	url := fmt.Sprintf("%s:%d", hostname, port)

	log.Println("up and running, ", url)

	log.Fatal(http.ListenAndServe(url, nil))
}

// StartServer Starts server and initializes endpoints on localhost with port 8888
func StartServer() {
	StartServerWithHostnameAndPort(hostname, port)
}

func handshake(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	clientId := params.Get("clientId")

	acceptConnection(clientId, w, r)
}

func handshakeStop(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	clientId := params.Get("clientId")

	if err := closeConnection(clientId); err != nil {
		w.Write([]byte("handshake couldn't stop"))
	} else {
		w.Write([]byte("handshake completely stopped"))
	}

}

func handshakeStatus(w http.ResponseWriter, r *http.Request) {

	clientId := r.URL.Query().Get("clientId")

	if status, err := getStatus(clientId); err == nil {
		res, _ := json.Marshal(status)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no connection found"))
	}
}

func sendNotification(w http.ResponseWriter, r *http.Request) {

	clientId := r.URL.Query().Get("clientId")

	notification, _ := pushNotification(clientId)

	bytes, err := json.Marshal(notification)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("notification send operation failed"))
	}

	w.Write(bytes)
}
