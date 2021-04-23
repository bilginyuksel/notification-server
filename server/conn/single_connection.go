package conn

// Connection Represents a client server connection. Implement this interface to communicate with the client.
type Connection interface {
	SendJSON(data interface{}) error
	Send(data []byte) error
	Close() error
}
