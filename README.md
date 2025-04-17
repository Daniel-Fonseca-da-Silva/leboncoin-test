# FizzBuzz REST API

A production-ready REST API implementation of the FizzBuzz problem using Go, Chi router, and clean architecture principles.

## Features

- FizzBuzz calculation with customizable parameters
- Request statistics tracking
- Clean architecture with separate layers (handlers, services, repositories, models)
- Input validation
- Graceful shutdown
- Production-ready middleware

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

## Testing the API

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

## Production Considerations

- The server includes middleware for:
  - Request logging
  - Panic recovery
  - Real IP detection
  - Request ID tracking
  - Request timeout (60 seconds)
- Graceful shutdown on SIGINT and SIGTERM signals
- Input validation for all requests
- Thread-safe statistics tracking
- Health check endpoint for monitoring 