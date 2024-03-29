package handlers

import (
	"concurrent-web-app/utilities"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	products, err := utilities.Load()
	if err != nil {
		log.Println("Error loading products:", err)
	}
	res, err := template.ParseFiles("templates/home.html")
	utilities.Check(err)
	err = res.Execute(w, products)
	utilities.Check(err)
	log.Println("Successful request!")
}

func Form(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/form.html")
	utilities.Check(err)
	err = res.Execute(w, nil)
	utilities.Check(err)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	price, err := strconv.ParseFloat(r.FormValue("productPrice"), 64)
	utilities.Check(err)
	product := utilities.NewProduct(r.FormValue("productName"), price)

	utilities.Add(&product)
	log.Println("New product successfully added to file")

	http.Redirect(w, r, "/home", http.StatusFound)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	utilities.Remove()
	log.Println("Removed all products!")

	http.Redirect(w, r, "/home", http.StatusFound)
}
