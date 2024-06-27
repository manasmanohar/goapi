package main

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Create a sample nested JSON response
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
		},
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the user struct to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/user", userHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
