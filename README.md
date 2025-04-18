# FizzBuzz REST API

A production-ready REST API implementation of the FizzBuzz problem using Go, Chi router, and clean architecture principles.

## Features

- FizzBuzz calculation with customizable parameters
- Request statistics tracking
- Clean architecture with separate layers (handlers, services, repositories, models)
- Input validation
- Graceful shutdown
- Production-ready middleware
- Docker support
- Structured JSON logging system

## API Endpoints

### Health Check

```
GET /api/health
```

Response:
```json
{
    "status": "ok"
}
```

### Calculate FizzBuzz

```
POST /fizzbuzz
```

Request body:
```json
{
    "int1": 3,
    "int2": 5,
    "limit": 100,
    "str1": "fizz",
    "str2": "buzz"
}
```

Response:
```json
{
    "result": ["1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", ...]
}
```

### Get Statistics

```
GET /stats
```

Response:
```json
{
    "request": {
        "int1": 3,
        "int2": 5,
        "limit": 100,
        "str1": "fizz",
        "str2": "buzz"
    },
    "hits": 42
}
```

## Setup and Running

### Local Development

1. Make sure you have Go 1.21 or later installed
2. Clone the repository
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the server:
   ```bash
   go run main.go
   ```

The server will start on port 8080.

### Docker Deployment

The project includes Docker support with a Makefile for easy management:

```bash
# Build the Docker image
make build

# Start the container
make up

# Stop and remove the container
make down
```

## Logging System

The application uses a structured JSON logging system with the following features:

- Log levels: DEBUG, INFO, WARN, ERROR
- JSON-formatted log entries
- Timestamp in RFC3339 format
- Request ID tracking
- Additional context fields for each log entry

Example log entry:
```json
{
    "timestamp": "2024-04-18T14:30:00Z",
    "level": "INFO",
    "message": "Request started",
    "fields": {
        "request_id": "req-123",
        "method": "POST",
        "path": "/fizzbuzz",
        "remote_ip": "127.0.0.1"
    }
}
```

Logs are written to stdout by default and can be redirected to a file or log management system in production.

## Testing

### Running Tests

To run the tests, use the following command:

```bash
# Run all tests
go test -v ./...

# Run specific package tests
go test -v ./internal/service/...
```

### API Testing

You can test the API using curl:

```bash
# Health Check
curl http://localhost:8080/api/health

# Calculate FizzBuzz
curl -X POST http://localhost:8080/fizzbuzz \
  -H "Content-Type: application/json" \
  -d '{"int1": 3, "int2": 5, "limit": 100, "str1": "fizz", "str2": "buzz"}'

# Get Statistics
curl http://localhost:8080/stats
```

## Architecture

The application follows clean architecture principles with the following layers:

- **Models**: Data structures and validation rules
- **Repository**: Data storage and retrieval
- **Service**: Business logic
- **Handler**: HTTP request handling and routing
- **Config**: Application configuration and logging

## Production Considerations

- The server includes middleware for:
  - Request logging with structured JSON format
  - Panic recovery
  - Real IP detection
  - Request ID tracking
  - Request timeout (60 seconds)
- Graceful shutdown on SIGINT and SIGTERM signals
- Input validation for all requests
- Thread-safe statistics tracking
- Health check endpoint for monitoring
- Docker support for containerized deployment
- Configurable logging levels (DEBUG, INFO, WARN, ERROR)