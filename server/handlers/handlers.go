package handlers

import (
	"concurrent-web-app/utilities"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/home.html")
	utilities.Check(err)
	err = res.Execute(w, nil)
	utilities.Check(err)
	log.Println("Successful request!")
}
