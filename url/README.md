# URL Building Utilities

The url package provides utilities for building and manipulating URLs in Go applications using dependency injection.

## Features

- **URLBuilder**: Configurable URL builder with dependency injection
- **RootURL**: Get the configured root URL
- **BuildURL**: Create full URLs with paths and query parameters
- **BuildQuery**: Convert parameter maps to query strings
- **HttpBuildQuery**: Convert url.Values to query strings

## Installation

```go
import "github.com/dracory/base/url"
```

## Usage

### Using URLBuilder (Recommended)

```go
// Create a URL builder with your application's root URL
builder := url.NewURLBuilder("https://example.com")

// Build URLs
fullURL := builder.BuildURL("api/users", nil)
// Returns: "https://example.com/api/users"

// With query parameters
params := map[string]string{
    "page": "1",
    "limit": "10",
}
fullURL = builder.BuildURL("api/users", params)
// Returns: "https://example.com/api/users?page=1&limit=10"
```

### Using Default Functions

For simple use cases, you can set a default URL and use the convenience functions:

```go
// Set the default URL (typically done during application initialization)
url.SetDefaultURL("https://example.com")

// Use default functions
root := url.RootURL()
// Returns: "https://example.com"

fullURL := url.BuildURL("api/users", map[string]string{"active": "true"})
// Returns: "https://example.com/api/users?active=true"
```

### Dependency Injection Pattern

The recommended approach is to inject the URLBuilder into your services:

```go
type UserService struct {
    urlBuilder *url.URLBuilder
}

func NewUserService(urlBuilder *url.URLBuilder) *UserService {
    return &UserService{urlBuilder: urlBuilder}
}

func (s *UserService) GetUserURL(id string) string {
    return s.urlBuilder.BuildURL("users/"+id, nil)
}

// In your main.go
func main() {
    appURL := os.Getenv("APP_URL")
    urlBuilder := url.NewURLBuilder(appURL)
    
    userService := NewUserService(urlBuilder)
    // ...
}
```

## Functions

### NewURLBuilder(rootURL string) *URLBuilder

Creates a new URLBuilder with the given root URL.

**Parameters:**
- `rootURL`: The base URL for all URLs built by this builder

**Returns:**
- `*URLBuilder`: A new URLBuilder instance

### URLBuilder Methods

#### RootURL() string

Returns the configured root URL.

#### BuildURL(path string, params map[string]string) string

Builds a complete URL by combining the root URL with a path and optional query parameters.

**Parameters:**
- `path`: The path to append to the root URL
- `params`: Optional map of query parameters

**Returns:**
- `string`: The complete URL (empty string if root URL is not set)

#### BuildQuery(queryData map[string]string) string

Converts a map of string parameters to a URL query string.

**Parameters:**
- `queryData`: Map of parameter key-value pairs

**Returns:**
- `string`: Query string (including "?" if parameters exist)

#### HttpBuildQuery(queryData url.Values) string

Converts url.Values to a URL-encoded query string.

**Parameters:**
- `queryData`: url.Values to encode

**Returns:**
- `string`: URL-encoded query string

### Default Functions

#### SetDefaultURL(rootURL string)

Sets the default URL builder's root URL for convenience functions.

#### RootURL() string

Returns the default root URL.

#### BuildURL(path string, params map[string]string) string

Builds a URL using the default builder.

#### BuildQuery(queryData map[string]string) string

Builds a query string using the default builder.

#### HttpBuildQuery(queryData url.Values) string

Builds a query string using the default builder.

## Examples

```go
// Dependency injection approach
func SetupServices(appURL string) {
    urlBuilder := url.NewURLBuilder(appURL)
    
    // Inject into services
    userService := NewUserService(urlBuilder)
    apiService := NewAPIService(urlBuilder)
    
    // Services can now build URLs without environment dependencies
    userURL := userService.GetUserURL("123")
    // "https://myapp.com/users/123"
}

// Default function approach
func InitializeApp() {
    appURL := os.Getenv("APP_URL")
    url.SetDefaultURL(appURL)
    
    // Throughout the application
    apiURL := url.BuildURL("api/v1/users", map[string]string{"active": "true"})
    // "https://myapp.com/api/v1/users?active=true"
}
```

## Testing

The package includes comprehensive tests:

```bash
go test ./url
```

## Dependencies

- Go 1.16+ 
- No external dependencies
- No environment variable dependencies (when using URLBuilder)
