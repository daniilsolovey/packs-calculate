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

# Set environment variables (can also be set in a .env file)
ENV API_HOST=0.0.0.0
ENV API_PORT=8080
ENV PACK_SIZES=5000,2000,1000,500,250

# Build the Go application
RUN go build -o packs-calculate-app .

# Expose the port the app will run on
EXPOSE 8080

# Run the application
CMD ["./packs-calculate-app"]
