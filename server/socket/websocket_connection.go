package socket

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type websocketConn struct {
	clientId                string
	conn                    *websocket.Conn
	connectionEstablishTime time.Time
	*ConnectionInfo
}

func (wc *websocketConn) SendJSON(data interface{}) error {
	wc.Collect(ConnectionSendJSON)
	return wc.conn.WriteJSON(data)
}

func (wc *websocketConn) Send(data []byte) error {
	wc.Collect(ConnectionSend)
	return wc.conn.WriteMessage(websocket.TextMessage, data)
}

func (wc *websocketConn) Close() error {
	wc.Collect(ConnectionClose)
	return wc.conn.Close()
}

func (wc *websocketConn) Status() ConnectionInfo {
	wc.Collect(ConnectionStatus)
	return *wc.ConnectionInfo
}

// NewWSConn Creates a special connection of gorilla websocket
func NewWSConn(clientId string, w http.ResponseWriter, r *http.Request) (Connection, error) {
	wsUpgrader := websocket.Upgrader{}
	ws, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil, err
	}

	return &websocketConn{
		clientId:                clientId,
		conn:                    ws,
		connectionEstablishTime: time.Now(),
		ConnectionInfo: &ConnectionInfo{
			CreateTime: time.Now(),
			Status:     Active,
		},
	}, nil
}
