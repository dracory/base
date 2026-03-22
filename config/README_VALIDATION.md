# Configuration Validation Utilities

Generic configuration validation utilities for environment variable loading and error accumulation.

## Overview

This package provides reusable validation infrastructure for loading and validating configuration from environment variables. It includes error types, validation functions, and an accumulator pattern for collecting multiple validation errors.

## Components

### 1. MissingEnvError

Custom error type for missing required environment variables.

```go
type MissingEnvError struct {
    Key     string
    Context string
}
```

**Features:**
- Structured error with key and context
- Formatted error messages
- Type-safe error handling

**Example:**
```go
err := config.MissingEnvError{
    Key:     "DATABASE_URL",
    Context: "required for database connection",
}
// Error: config: required env "DATABASE_URL" is missing: required for database connection
```

### 2. Validation Functions

#### RequireString

Retrieves and validates a required environment variable.

```go
func RequireString(key, context string) (string, error)
```

**Features:**
- Reads from environment variables
- Trims whitespace
- Returns error if empty or missing

**Example:**
```go
dbHost, err := config.RequireString("DB_HOST", "database connection")
if err != nil {
    // Handle missing environment variable
}
```

#### RequireWhen

Conditionally validates an environment variable.

```go
func RequireWhen(condition bool, key, context, value string) error
```

**Features:**
- Only validates when condition is true
- Useful for optional features
- Returns nil when condition is false

**Example:**
```go
err := config.RequireWhen(
    usePostgres,
    "DB_PASSWORD",
    "required for PostgreSQL",
    dbPassword,
)
```

#### EnsureRequired

Low-level validation that a value is present.

```go
func EnsureRequired(value, key, context string) error
```

**Features:**
- Validates any string value
- Trims whitespace
- Returns MissingEnvError if empty

**Example:**
```go
err := config.EnsureRequired(apiKey, "API_KEY", "authentication")
```

### 3. LoadAccumulator

Error accumulator for collecting multiple validation errors during configuration loading.

```go
type LoadAccumulator struct {
    // internal fields
}
```

**Methods:**

- `Add(err error)` - Add an error to the accumulator
- `MustString(key, context string) string` - Load required string, recording errors
- `MustWhen(condition bool, key, context, value string)` - Conditional validation
- `Err() error` - Get accumulated errors as ValidationError

**Example:**
```go
acc := &config.LoadAccumulator{}

// Collect multiple validation errors
dbHost := acc.MustString("DB_HOST", "database connection")
dbPort := acc.MustString("DB_PORT", "database connection")
acc.MustWhen(useSSL, "DB_SSL_CERT", "SSL connection", sslCert)

// Check for any errors
if err := acc.Err(); err != nil {
    // Handle all validation errors at once
    return nil, err
}
```

### 4. ValidationError

Aggregates multiple validation errors into a single error.

```go
type ValidationError struct {
    // internal fields
}
```

**Methods:**

- `Error() string` - Formatted error message with all errors
- `Errors() []error` - Get all accumulated errors

**Example:**
```go
if err := acc.Err(); err != nil {
    if verr, ok := err.(config.ValidationError); ok {
        for _, e := range verr.Errors() {
            log.Printf("Validation error: %v", e)
        }
    }
}
```

## Usage Patterns

### Basic Validation

```go
import "github.com/dracory/base/config"

// Validate single required value
apiKey, err := config.RequireString("API_KEY", "API authentication")
if err != nil {
    return err
}
```

### Conditional Validation

```go
// Only require when feature is enabled
err := config.RequireWhen(
    stripeEnabled,
    "STRIPE_SECRET_KEY",
    "Stripe payment processing",
    stripeKey,
)
```

### Batch Validation with Accumulator

```go
acc := &config.LoadAccumulator{}

// Load multiple required values
appHost := acc.MustString("APP_HOST", "application server")
appPort := acc.MustString("APP_PORT", "application server")
dbDriver := acc.MustString("DB_DRIVER", "database connection")

// Conditional requirements
acc.MustWhen(
    dbDriver == "postgres",
    "DB_PASSWORD",
    "PostgreSQL requires password",
    dbPassword,
)

// Check all errors at once
if err := acc.Err(); err != nil {
    return nil, err
}
```

### Error Handling

```go
if err := acc.Err(); err != nil {
    // Check if it's a ValidationError
    if verr, ok := err.(config.ValidationError); ok {
        // Access individual errors
        for _, e := range verr.Errors() {
            // Check if it's a MissingEnvError
            if merr, ok := e.(config.MissingEnvError); ok {
                log.Printf("Missing: %s (%s)", merr.Key, merr.Context)
            }
        }
    }
}
```

## Benefits

1. **Reusable**: Generic validation logic shared across all Dracory projects
2. **Type-Safe**: Structured error types for better error handling
3. **Accumulation**: Collect all validation errors before failing
4. **Context**: Rich error messages with context for debugging
5. **Tested**: Comprehensive test coverage ensures reliability
6. **Consistent**: Standardized validation patterns across projects

## Integration

### Blueprint Integration

Blueprint uses these utilities directly via type aliases:

```go
// blueprint/internal/config/functions.go
type MissingEnvError = baseConfig.MissingEnvError

func requireString(key, context string) (string, error) {
    return baseConfig.RequireString(key, context)
}

// blueprint/internal/config/loader_accumulator.go
type loadAccumulator = baseConfig.LoadAccumulator
type validationError = baseConfig.ValidationError
```

Blueprint code uses the base package directly with capitalized method names:
```go
acc := &loadAccumulator{}
host := acc.MustString("APP_HOST", "application server")
acc.MustWhen(condition, "DB_PASSWORD", "required for PostgreSQL", password)
if err := acc.Err(); err != nil {
    return nil, err
}
```

## Testing

Comprehensive tests cover:
- Error message formatting
- Environment variable loading
- Whitespace trimming
- Conditional validation
- Error accumulation
- Defensive copying
- Integration scenarios

Run tests:
```bash
go test ./config -v -run "TestMissingEnvError|TestEnsureRequired|TestRequireWhen|TestRequireString|TestLoadAccumulator|TestValidationError"
```

## Migration Notes

When migrating from project-specific validation to base package:

1. Import the base config package
2. Replace local types with base types
3. Update function calls to use exported names (RequireString vs requireString)
4. Update tests to use public methods (Errors() vs errs field)
5. Maintain backward compatibility with wrapper functions if needed
