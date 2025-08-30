package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

func writeEndpoint(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func TestClient(t *testing.T) {
	defer cancel()
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, "ws://127.0.0.1:8080/conn", nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	writeEndpoint(conn)

	fmt.Println("Connection opened")
}
