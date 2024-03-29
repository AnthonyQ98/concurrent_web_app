package handlers

import (
	"concurrent-web-app/utilities"
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Name  string
	Price float32
}

func NewProduct(name string, price float32) Product {
	if name == "" {
		name = "Foobar"
	}
	if price < 0 {
		price = 10.0
	}
	return Product{
		Name:  name,
		Price: price,
	}
}

func productBuilder() []Product {
	products := []Product{NewProduct("Jim", 12.0), NewProduct("Bob", -10)}
	return products
}

func Home(w http.ResponseWriter, r *http.Request) {
	products := productBuilder()
	res, err := template.ParseFiles("templates/home.html")
	utilities.Check(err)
	err = res.Execute(w, products)
	utilities.Check(err)
	log.Println("Successful request!")
}
