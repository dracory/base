# Config Encryption Loader

This package provides functionality to load environment variables from encrypted vault files using the `envenc` package.

## Overview

The Config Encryption Loader supports two distinct approaches for loading encrypted environment variables:

1. **File-based loading**: Load vault files from the local filesystem
2. **Resource-based loading**: Load vault content from embedded resources

## Functions

### File-based Loading

```go
func InitializeEnvEncVariablesFromFile(appEnvironment, publicKey, privateKey string) error
```

Loads environment variables from encrypted vault files on the filesystem.

**Parameters:**
- `appEnvironment`: The application environment (e.g., "development", "production", "staging")
- `publicKey`: The public encryption key (typically from ENV_ENCRYPTION_KEY_PUBLIC)
- `privateKey`: The private encryption key (typically from ENV_ENCRYPTION_KEY_PRIVATE)

**Behavior:**
- Looks for vault file named `.env.<app_environment>.vault` in the local filesystem
- Returns error if file is not found
- Uses `envenc.HydrateEnvFromFile()` for hydration

### Resource-based Loading

```go
func InitializeEnvEncVariablesFromResources(appEnvironment, publicKey, privateKey string, resourceLoader func(string) (string, error)) error
```

Loads environment variables from encrypted vault content from embedded resources.

**Parameters:**
- `appEnvironment`: The application environment (e.g., "development", "production", "staging")
- `publicKey`: The public encryption key (typically from ENV_ENCRYPTION_KEY_PUBLIC)
- `privateKey`: The private encryption key (typically from ENV_ENCRYPTION_KEY_PRIVATE)
- `resourceLoader`: Function to load embedded resources, returns vault content

**Behavior:**
- Looks for vault resource named `.env.<app_environment>.vault` in embedded resources
- Returns error if resource is not found or empty
- Uses `envenc.HydrateEnvFromString()` for hydration

## Usage Examples

### Basic File-based Usage

```go
package main

import (
    "log"
    "os"
    "github.com/dracory/base/config"
)

func main() {
    // Load from filesystem
    err := config.InitializeEnvEncVariablesFromFile(
        "production",
        os.Getenv("ENV_ENCRYPTION_KEY_PUBLIC"),
        os.Getenv("ENV_ENCRYPTION_KEY_PRIVATE"),
    )
    if err != nil {
        log.Fatalf("Failed to load environment variables: %v", err)
    }
    
    // Environment variables are now available via os.Getenv()
    dbHost := os.Getenv("DB_HOST")
    dbPassword := os.Getenv("DB_PASSWORD")
    
    log.Printf("Database host: %s", dbHost)
}
```

### Embedded Resources Usage

```go
package main

import (
    "log"
    "embed"
    "os"
    "github.com/dracory/base/config"
)

//go:embed *.vault
var vaultFS embed.FS

func main() {
    // Create resource loader
    resourceLoader := func(name string) (string, error) {
        content, err := vaultFS.ReadFile(name)
        if err != nil {
            return "", err
        }
        return string(content), nil
    }
    
    // Load from embedded resources
    err := config.InitializeEnvEncVariablesFromResources(
        "production",
        os.Getenv("ENV_ENCRYPTION_KEY_PUBLIC"),
        os.Getenv("ENV_ENCRYPTION_KEY_PRIVATE"),
        resourceLoader,
    )
    if err != nil {
        log.Fatalf("Failed to load environment variables: %v", err)
    }
}
```

### Caller-Controlled Testing

Testing environment detection is completely up to the caller:

```go
// Skip loading in testing environments
if os.Getenv("APP_ENV") != "testing" {
    err := config.InitializeEnvEncVariablesFromFile(
        appEnv,
        publicKey,
        privateKey,
    )
    if err != nil {
        return err
    }
}

// Or with more complex logic
shouldLoadConfig := !isTestingEnvironment() || config.ForceLoadInTests()
if shouldLoadConfig {
    err := config.InitializeEnvEncVariablesFromResources(
        appEnv,
        publicKey,
        privateKey,
        resourceLoader,
    )
}
```

## Environment Variables

The loader expects these environment variables to be set:

- `APP_ENVIRONMENT`: Application environment (development, production, staging, etc.)
- `ENV_ENCRYPTION_KEY_PUBLIC`: Public encryption key
- `ENV_ENCRYPTION_KEY_PRIVATE`: Private encryption key

## Vault Files

Vault files should be named following the pattern:
`.env.{environment}.vault`

For example:
- `.env.development.vault`
- `.env.production.vault`
- `.env.staging.vault`

The loader will look for vault files in this order:
1. Local filesystem (e.g., `./.env.development.vault`)
2. Embedded resources (via resource loader function)

## Error Types

### MissingEnvError

Raised when required environment variables are missing:

```go
type MissingEnvError struct {
    Key     string // Missing environment variable name
    Context string // Context of why it's required
}
```

### EnvEncError

Raised during vault operations:

```go
type EnvEncError struct {
    Operation string // Operation that failed (derive_key, hydrate_from_file, etc.)
    Message   string // Detailed error message
}
```

## Security Considerations

1. **Key Management**: Store encryption keys securely (e.g., in secure environment variables, key management services)
2. **Vault Files**: Commit encrypted vault files to your repository, but never commit the keys
3. **Environment Separation**: Use different vault files for different environments
4. **Access Control**: Limit access to vault files and encryption keys

## Integration with Blueprint

To use this in your Blueprint application, replace the `initializeEnvEncVariables` function in `internal/config/load.go`:

```go
// Old code
if err := initializeEnvEncVariables(app.env, ENV_ENCRYPTION_KEY_PUBLIC, envEnc.privateKey); err != nil {
    acc.add(err)
}

// New code (file-based)
if err := config.InitializeEnvEncVariablesFromFile(
    app.env, 
    ENV_ENCRYPTION_KEY_PUBLIC, 
    envEnc.privateKey,
); err != nil {
    acc.add(err)
}

// Or new code (resource-based)
if err := config.InitializeEnvEncVariablesFromResources(
    app.env, 
    ENV_ENCRYPTION_KEY_PUBLIC, 
    envEnc.privateKey,
    resources.Resource,
); err != nil {
    acc.add(err)
}
```

## Testing

The package includes comprehensive tests. Run them with:

```bash
go test ./config -v
```

For integration tests that use the file system:

```bash
go test ./config -v -run Integration
```

Skip integration tests in short mode:

```bash
go test ./config -v -short
```
