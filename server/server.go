package server

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/home.html")
	check(err)
	err = res.Execute(w, nil)
	check(err)
	log.Println("Successful request!")
}
func StartServer() {
	log.Println("Starting webserver on port 8080...")
	http.HandleFunc("/home", home)
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}
