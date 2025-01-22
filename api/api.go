package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daniilsolovey/packs-calculate/config"
	"github.com/daniilsolovey/packs-calculate/internal/pack"
)

// API handler to calculate optimal packs
func CalculatePacksHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body for the number of items
	var request struct {
		Number int `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Calculate the optimal number of packs using the loaded pack sizes
	result, err := pack.CalculatePacks(request.Number, cfg.PackSizes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error calculating packs: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the result as a JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// StartServer starts the HTTP server with all routes
func StartServer(cfg *config.Config) {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		CalculatePacksHandler(w, r, cfg)
	})

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Println("Server is running on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
