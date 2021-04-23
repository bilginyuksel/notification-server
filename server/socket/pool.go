package socket

import (
	"errors"
)

var singletonPool = &connectionPoolImpl{
	connections: make(map[string]Connection),
}

type connectionPoolImpl struct {
	connections map[string]Connection
}

type ConnectionPool interface {
	Add(id string, conn Connection) error
	Get(id string) (Connection, error)
}

func GetPoolInstance() ConnectionPool {
	return singletonPool
}

// Add ...
func (cp *connectionPoolImpl) Add(id string, conn Connection) error {
	if _, err := cp.Get(id); err == nil {
		return errors.New("connection already exist")
	}

	cp.connections[id] = conn
	return nil
}

func (cp *connectionPoolImpl) Get(id string) (Connection, error) {
	conn, ok := cp.connections[id]
	if !ok {
		return nil, errors.New("no socket connection found")
	}
	return conn, nil
}
