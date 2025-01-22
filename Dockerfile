FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port the app will run on
EXPOSE 8080

# Run the app when the container starts
CMD ["./app"]
