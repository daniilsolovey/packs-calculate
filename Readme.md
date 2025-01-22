# Shipping Pack Calculator API

## Description
This Go-based API calculates the number of packs needed to fulfill an order based on the number of items ordered. The available pack sizes are dynamic and can be changed without modifying the code.


## Installation
1. Clone the repository.
2. Build the Go application using `go build`.
3. Run the application with `go run main.go`.


## Docker
To containerize the application:
1. Build the Docker image: `docker build -t packs .`
2. Run the container: `docker run -p 8080:8080 packs`

## Unit Tests
Run tests using `go test`.


