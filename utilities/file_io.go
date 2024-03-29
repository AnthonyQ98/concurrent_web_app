package utilities

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Product struct {
	Name  string
	Price float64
}

func NewProduct(name string, price float64) Product {
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
func Add(product *Product) {
	// Open file in append mode or create it if it doesn't exist
	file, err := os.OpenFile("products.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Convert product fields to string and write to CSV
	err = writer.Write([]string{product.Name, strconv.FormatFloat(product.Price, 'f', 2, 64)})
	if err != nil {
		panic(err)
	}
}

func Remove() {
	// Open file in truncate mode or create it if it doesn't exist
	file, err := os.OpenFile("products.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}

func Load() ([]Product, error) {
	// Open the CSV file
	file, err := os.Open("products.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Initialize slice of products
	var products []Product

	// Iterate over records and create product structs
	for _, record := range records {
		// Convert price string to float64
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		// Create product struct
		product := Product{
			Name:  record[0],
			Price: price,
		}

		// Append product to slice
		products = append(products, product)
	}

	return products, nil
}
