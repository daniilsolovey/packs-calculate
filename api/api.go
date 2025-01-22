package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daniilsolovey/packs-calculate/internal/pack"
)

// The available pack sizes, used across the application
var packSizes = []int{5000, 2000, 1000, 500, 250}

// API handler to calculate optimal packs
func CalculatePacksHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(
			w, "Invalid request method", http.StatusMethodNotAllowed,
		)
		return
	}

	// Parse request body for the number of items
	var request struct {
		Number int `json:"number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w, "Invalid input", http.StatusBadRequest,
		)
		return
	}

	// Calculate the optimal number of packs
	result, err := pack.CalculatePacks(request.Number, packSizes)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("Error calculating packs: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Return the result as a JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(
			w, "Error encoding response", http.StatusInternalServerError,
		)
	}
}

// StartServer starts the HTTP server with all routes
func StartServer() {
	http.HandleFunc("/calculate", CalculatePacksHandler)

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
