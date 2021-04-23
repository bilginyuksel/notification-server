package endpoint

import (
	"log"
	"net/http"

	"github.com/bilginyuksel/notification-server/server/entity"
	"github.com/bilginyuksel/notification-server/server/socket"
)

func acceptConnection(clientId string, w http.ResponseWriter, r *http.Request) {
	wsConn, err := socket.NewWSConn(clientId, w, r)

	if err != nil {
		log.Println("handshake failed, err= ", err)
	}

	socket.GetPoolInstance().Add(clientId, wsConn)
	log.Println("accepted connection with client= ", clientId)
}

func closeConnection(clientId string) error {
	conn, err := socket.GetPoolInstance().Get(clientId)

	if err != nil {
		log.Println("connection not exists")
		return err
	}

	return conn.Close()
}

func getStatus(clientId string) (socket.ConnectionInfo, error) {

	wscon, err := socket.GetPoolInstance().Get(clientId)

	if err != nil {
		log.Println("connection not exists")
		return socket.ConnectionInfo{}, err
	}

	return wscon.Status(), nil
}

func pushNotification(clientId string) (entity.Notification, error) {
	conn, err := socket.GetPoolInstance().Get(clientId)

	if err != nil {
		log.Println("connection not exists")
		return entity.Notification{}, err
	}

	notification := entity.Notification{
		Title:       "Go Sample Push Notification Server",
		Description: "Hello this is a notification.",
	}

	err = conn.SendJSON(notification)

	return notification, err
}
