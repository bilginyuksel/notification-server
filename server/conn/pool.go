package conn

import (
	"errors"

	"github.com/bilginyuksel/notification-server/server/entity"
	"github.com/gorilla/websocket"
)

var singletonPool = &connectionPoolImpl{
	socketConnections: make(map[string]*websocket.Conn),
}

type connectionPoolImpl struct {
	socketConnections map[string]*websocket.Conn
}

type ConnectionPool interface {
	Add(id string, conn *websocket.Conn) error
	Write(id string, notification entity.Notification) error
	Has(id string) bool
	Close(id string) error
}

func GetPoolInstance() ConnectionPool {
	return singletonPool
}

// Add ...
func (cp *connectionPoolImpl) Add(id string, conn *websocket.Conn) error {
	if _, err := cp.get(id); err == nil {
		return errors.New("connection already exist")
	}

	cp.socketConnections[id] = conn
	return nil
}

// Write ...
func (cp *connectionPoolImpl) Write(id string, notification entity.Notification) error {
	conn, err := cp.get(id)

	if err != nil {
		return err
	}

	return conn.WriteJSON(notification)
}

// Has ...
func (cp *connectionPoolImpl) Has(id string) bool {
	_, ok := cp.socketConnections[id]
	return ok
}

func (cp *connectionPoolImpl) Close(id string) error {
	conn, err := cp.get(id)

	if err != nil {
		return err
	}

	return conn.Close()
}

func (cp *connectionPoolImpl) get(id string) (*websocket.Conn, error) {
	conn, ok := cp.socketConnections[id]
	if !ok {
		return nil, errors.New("no socket connection found")
	}
	return conn, nil
}
