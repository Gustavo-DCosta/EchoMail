package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	reader(conn)

}

func TestServer(t *testing.T) {
	http.HandleFunc("/conn", wsEndpoint)

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}

	fmt.Println("Server listening on port 8080:")
}
