package main

import (
	"fmt"

	"github.com/daniilsolovey/packs-calculate/api"
)

// func main() {
// 	number := 800
// 	packs := []int{5000, 2000, 1000, 500, 250}

// 	result, err := pack.CalculatePacks(number, packs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Result:", result)
// }

func main() {
	// Start the API server
	api.StartServer()

	// Optionally print something indicating the server is up
	fmt.Println("API server started on http://localhost:8080")
}
