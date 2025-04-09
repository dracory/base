# Server Package

The `server` package provides a simple and configurable HTTP server implementation with support for graceful shutdown, different operating modes, and configurable logging levels.

## Features

- **Configurable Server Options**: Set host, port, URL, handler, mode, and log level
- **Multiple Operating Modes**: Production and testing modes
- **Configurable Logging**: Debug, info, error, and none log levels
- **Graceful Shutdown**: Handles OS signals (SIGINT, SIGTERM) for clean server termination
- **Colorized Logging**: Uses `cfmt` for better visibility of server status

## Usage

### Basic Example

```go
package main

import (
    "net/http"
    "your-project/server"
)

func main() {
    // Define a simple handler
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    }

    // Start the server
    server.Start(server.Options{
        Host:    "localhost",
        Port:    "8080",
        URL:     "http://localhost:8080",
        Handler: handler,
        Mode:    server.ProductionMode,
        LogLevel: server.LogLevelInfo,
    })
}
```

### Configuration Options

The `Options` struct provides the following configuration options:

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| Host | string | The host to bind the server to | Required |
| Port | string | The port to bind the server to | Required |
| URL | string | The URL displayed in logs | Optional |
| Handler | http.HandlerFunc | The HTTP handler function | Required |
| Mode | string | Server mode (production/testing) | "production" |
| LogLevel | LogLevel | Logging level | "info" |

### Log Levels

The package supports the following log levels:

- `LogLevelDebug`: Detailed debugging information
- `LogLevelInfo`: General information about server operations
- `LogLevelError`: Error messages only
- `LogLevelNone`: No logging

### Operating Modes

- `ProductionMode`: Standard production mode with normal error handling
- `TestingMode`: Special mode for testing with different error handling

## Testing

The package includes a test file (`start_test.go`) that demonstrates how to test the server functionality, including:

- Starting the server
- Making requests to verify it's running
- Gracefully shutting down the server
- Verifying the server has shut down

## Dependencies

- `github.com/mingrammer/cfmt`: Colorized formatting for logs

## License

This package is part of the main project and subject to its license terms. 