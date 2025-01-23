package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/daniilsolovey/packs-calculate/api"
	"github.com/daniilsolovey/packs-calculate/config"
	"github.com/daniilsolovey/packs-calculate/internal/pack"
)

func main() {
	// Define a flag for the .env file location
	envFilePath := flag.String("env", "", "Path to the .env file")
	flag.Parse()

	// Load the configuration
	cfg, err := config.LoadConfig(*envFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Check if .env file was used or default values were used
	if *envFilePath != "" {
		fmt.Printf("Using configuration from file, path to file: %s\n", *envFilePath)
	} else {
		fmt.Println("No .env file provided, using default environment variables.")
	}

	// Create a packCalculator instance using the constructor
	packCalculator := pack.NewPackCalculator()

	// Start the API server with the loaded configuration
	api.StartServer(cfg, packCalculator)

	// Optionally print something indicating the server is up
	fmt.Println("API server started")
}
