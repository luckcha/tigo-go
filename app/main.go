package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Interface
type Describable interface {
	GetInfo() string
}

type Book struct {
	BookName    string  `json:"book_name"`
	Publication string  `json:"publication"`
	Chapter     int     `json:"chapter"`
	Price       float64 `json:"price"`
}

type Dealer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Sale     float64 `json:"sale"`
}

// Implementing interface methods
func (b Book) GetInfo() string {
	return fmt.Sprintf("Book: %s (%s) - %d chapters, Price: %.2f",
		b.BookName, b.Publication, b.Chapter, b.Price)
}

func (d Dealer) GetInfo() string {
	return fmt.Sprintf("Dealer: %s (ID: %d), Location: %s, Sales: %.2f",
		d.Name, d.ID, d.Location, d.Sale)
}

var Product []Book
var Production []Dealer

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Hey there Lucky 👋, welcome to my Go API with Interfaces!")
}

func returnAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProduct")
	json.NewEncoder(w).Encode(Product)
}

func returnAllProduction(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProduction")
	json.NewEncoder(w).Encode(Production)
}

// New endpoint using interface
func describeAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: describeAll")

	var descriptions []string
	for _, p := range Product {
		descriptions = append(descriptions, p.GetInfo())
	}
	for _, d := range Production {
		descriptions = append(descriptions, d.GetInfo())
	}

	json.NewEncoder(w).Encode(descriptions)
}

func handleRequests() {
	http.HandleFunc("/product", returnAllProduct)
	http.HandleFunc("/production", returnAllProduction)
	http.HandleFunc("/describe", describeAll)
	http.HandleFunc("/", homepage)
	log.Println("Server starting on :10000...")
	http.ListenAndServe(":10000", nil)
}

func main() {
	Product = []Book{
		{BookName: "Science", Publication: "Mittal", Chapter: 15, Price: 390.00},
		{BookName: "Mathematics", Publication: "Oxford", Chapter: 12, Price: 450.00},
		{BookName: "History", Publication: "Penguin", Chapter: 20, Price: 550.00},
		{BookName: "Hindi", Publication: "Mittal", Chapter: 10, Price: 350.00},
		{BookName: "Economy", Publication: "Oxford", Chapter: 17, Price: 300.00},
		{BookName: "Geography", Publication: "Mittal", Chapter: 19, Price: 250.00},
	}

	Production = []Dealer{
		{ID: 101, Name: "Jack", Location: "Minnisota", Sale: 455658.5},
		{ID: 102, Name: "Luther", Location: "New York", Sale: 358858.8},
		{ID: 103, Name: "Garry", Location: "London", Sale: 625895.8},
		{ID: 104, Name: "Adrian", Location: "Paris", Sale: 9427158.5},
		{ID: 105, Name: "Mickel", Location: "Washington", Sale: 575754.4},
		{ID: 106, Name: "Virat", Location: "Delhi", Sale: 84584.86},
	}

	handleRequests()
}
