package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavo-DCosta/server/router"
)

func main() {
	router.Router()

	fmt.Println("Server running...")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
