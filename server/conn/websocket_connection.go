package conn

import (
	"time"

	"github.com/gorilla/websocket"
)

type websocketConn struct {
	clientId                string
	conn                    *websocket.Conn
	connectionEstablishTime time.Time
}

func (wc *websocketConn) SendJSON(data interface{}) error {
	return wc.conn.WriteJSON(data)
}

func (wc *websocketConn) Send(data []byte) error {
	return wc.conn.WriteMessage(websocket.TextMessage, data)
}

func (wc *websocketConn) Close() error {
	return wc.conn.Close()
}

func NewWebsocketConnection(clientId string, conn *websocket.Conn) Connection {
	return &websocketConn{
		clientId:                clientId,
		conn:                    conn,
		connectionEstablishTime: time.Now(),
	}
}
