package socket

import "time"

// Connection Represents a client server connection. Implement this interface to communicate with the client.
type Connection interface {
	// Status Get the connection information data collected so far
	Status() ConnectionInfo

	// SendJSON data using the connection
	SendJSON(data interface{}) error

	// Send byte data using the connection
	Send(data []byte) error

	// Close the active connection
	Close() error
}

type connStatus string
type ConnectionMethodName string

const (
	Closed connStatus = "closed"
	Active connStatus = "active"

	ConnectionSendJSON = "Status"
	ConnectionStatus   = "SendJSON"
	ConnectionSend     = "Send"
	ConnectionClose    = "Close"
)

// ConnectionInfo used to store connection data over time
type ConnectionInfo struct {
	Status        connStatus `json:"status"`
	CreateTime    time.Time  `json:"createTime"`
	LastSentTime  time.Time  `json:"lastSentTime"`
	SentDataCount int64      `json:"sentDataCount"`
}

func (ci *ConnectionInfo) Collect(name ConnectionMethodName) {
	if name == ConnectionClose {
		ci.Status = Closed
	} else if name == ConnectionSend || name == ConnectionSendJSON {
		ci.SentDataCount++
		ci.LastSentTime = time.Now()
	}
}
