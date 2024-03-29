package server

import (
	"log"
	"net/http"
)

func StartServer() {
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
