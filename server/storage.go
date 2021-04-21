package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

var singletonPool = &connectionPoolImpl{
	socketConnections: make(map[string]*websocket.Conn),
}

type connectionPoolImpl struct {
	socketConnections map[string]*websocket.Conn
}

type ConnectionPool interface {
	get(id string) (*websocket.Conn, error)
	add(id string, conn *websocket.Conn) error
}

func GetPoolInstance() ConnectionPool {
	return singletonPool
}

func (cp *connectionPoolImpl) get(id string) (*websocket.Conn, error) {
	conn, ok := cp.socketConnections[id]
	if !ok {
		return nil, errors.New("no socket connection found")
	}
	return conn, nil
}

func (cp *connectionPoolImpl) add(id string, conn *websocket.Conn) error {
	if _, err := cp.get(id); err == nil {
		return errors.New("connection already exist")
	}

	cp.socketConnections[id] = conn
	return nil
}
