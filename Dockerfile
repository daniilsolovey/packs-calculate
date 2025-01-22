# Use the official Go image as a base
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port the app will run on (if needed)
EXPOSE 8080

# Run the application
CMD ["./main"]
