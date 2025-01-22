package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds the configuration values for the application
type Config struct {
	Host      string // Host of the API server
	Port      string // Port of the API server
	PackSizes []int  // List of pack sizes
}

// LoadConfig loads the configuration from environment variables or from a .env file
func LoadConfig(envFilePath string) (*Config, error) {
	// Load the .env file if it exists and envFilePath is provided
	if envFilePath != "" {
		err := godotenv.Load(envFilePath)
		if err != nil {
			return nil, fmt.Errorf("Error loading .env file: %v", err)
		}
	}

	// Load environment variables
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "localhost" // default value
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // default value
	}

	packSizesEnv := os.Getenv("PACK_SIZES")
	if packSizesEnv == "" {
		packSizesEnv = "5000,2000,1000,500,250" // default pack sizes
	}
	packSizes := parsePackSizes(packSizesEnv)

	config := &Config{
		Host:      host,
		Port:      port,
		PackSizes: packSizes,
	}
	return config, nil
}

// parsePackSizes parses the PACK_SIZES environment variable and converts it to a slice of integers
func parsePackSizes(packSizesEnv string) []int {
	var packSizes []int
	for _, s := range strings.Split(packSizesEnv, ",") {
		if size, err := strconv.Atoi(s); err == nil {
			packSizes = append(packSizes, size)
		}
	}
	return packSizes
}
