package socket

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var (
	wsTestMessage = []byte("ws connection established")
)

func testServerCallback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	clientId := query.Get("clientId")
	c, err := NewWSConn(clientId, w, r)

	if err != nil {
		log.Fatal("connection couldn't established, err: ", err)
	}

	c.Send(wsTestMessage)
}

func TestEstablishNewConnection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testServerCallback))

	defer server.Close()

	url := url.URL{Scheme: "ws", Host: strings.ReplaceAll(server.URL, "http://", ""), Path: "", RawQuery: "key=mykey"}

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		t.Fatal("client connection couldn't established, dial: ", err)
	}

	defer c.Close()

	_, bytes, _ := c.ReadMessage()

	assert.Equal(t, string(wsTestMessage), string(bytes),
		"the messages are not same, expected= %v, given= %v", string(wsTestMessage), string(bytes))

	log.Println(url)
	log.Println(string(bytes))
}
