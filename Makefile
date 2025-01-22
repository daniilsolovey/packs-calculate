# Makefile for Packs-Calculate

# Define variables
GO = go
GO_BUILD = $(GO) build
GO_TEST = $(GO) test
GO_MOD_TIDY = $(GO) mod tidy
DOCKER_BUILD = docker build -t packs-calculate .
DOCKER_RUN = docker run --rm -p 8080:8080 packs-calculate
BUILD_DIR = cmd/calculate

# Default target (when no target is specified)
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	$(GO_MOD_TIDY)
	$(GO_BUILD) -C $(BUILD_DIR) -o ../../packs-calculate

# Run tests
.PHONY: test
test:
	$(GO_TEST) ./... -v

# Build Docker image
.PHONY: docker-build
docker-build:
	$(DOCKER_BUILD)

# Run Docker container
.PHONY: docker-run
docker-run:
	$(DOCKER_RUN)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f packs-calculate
