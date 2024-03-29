package server

import (
	"concurrent-web-app/server/handlers"
	"concurrent-web-app/utilities"
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Starting webserver on port 8080...")
	http.HandleFunc("/home", handlers.Home)
	err := http.ListenAndServe("localhost:8080", nil)
	utilities.Check(err)
}
