.PHONY: up down build

# Default target
all: build

# Build the Docker image
build:
	docker build -t leboncoin-test .

# Start the container
up: build
	docker run -d --name leboncoin-test -p 8080:8080 leboncoin-test

# Stop and remove the container
down:
	docker stop leboncoin-test || true
	docker rm leboncoin-test || true 