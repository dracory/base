# HTTP Utilities

The http package provides HTTP-related utility functions for Go applications.

## Installation

```bash
go get github.com/dracory/base/http
```

## Usage

### SafeCloseResponseBody

Safely closes an HTTP response body with proper error handling and logging.

```go
import "github.com/dracory/base/http"

resp, err := http.Get("https://example.com")
if err != nil {
    // handle error
}
defer http.SafeCloseResponseBody(resp.Body)

// Use response body...
```

This utility function ensures consistent error handling and logging across the application. It safely handles nil responses and logs any closing errors without panicking.

### Redirect

Performs an HTTP redirect to the specified URL using temporary redirect (307).

```go
import "github.com/dracory/base/http"

func handler(w http.ResponseWriter, r *http.Request) string {
    // Redirect to login page
    return http.Redirect(w, r, "/login")
}
```

The function wraps the standard `http.Redirect` for consistency across the codebase and returns an empty string for compatibility with controller return patterns.
