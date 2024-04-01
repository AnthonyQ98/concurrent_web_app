package server

import (
	"concurrent-web-app/server/handlers"
	"concurrent-web-app/utilities"
	"log"
	"net/http"
)

func StartServer() {
    host, port := "0.0.0.0", "8080"
	log.Printf("Starting webserver on port %s & host %s...\n", port, host)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/new", handlers.Form)
	http.HandleFunc("/remove", handlers.Remove)
	http.HandleFunc("/addProduct", handlers.AddProduct)
	// Start a goroutine to serve incoming HTTP requests concurrently
	go func() {
		err := http.ListenAndServe(host+":"+port, nil)
		utilities.Check(err)
	}()

	// Using blocking select statement for simplicity
	select {}
}
