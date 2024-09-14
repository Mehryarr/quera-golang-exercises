package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	ID    int64
	Name  string
	Price float32
	Count int32
}

func main() {
	http.HandleFunc("/sayHelloWorld", handleHelloWorld)
	http.HandleFunc("/get-product/", getProduct)
	http.HandleFunc("/add-product", addProduct)
	http.ListenAndServe(":8080", nil)

}
func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello, world!",
		"status":  "success",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	productID := parts[2]
	fmt.Fprintf(w, "Requested product ID: %s\n", productID)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	/*
	   ...
	   Add product to the store
	   ...
	*/

	fmt.Fprintf(w, "Product added successfully:\n%+v\n", newProduct)
}
