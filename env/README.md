## Environment Variables

The Dracory framework utilizes environment variables for configuration. The `env` package provides utilities for loading and accessing these variables.

### Loading Environment Variables

Environment variables can be loaded from a `.env` file using the `env.Initialize()` function. This function attempts to load variables from a `.env` file in the current directory.

```go
import "github.com/dracory/base/env"

func main() {
    env.Initialize() // Loads from .env file
}
```

### Accessing Environment Variables

The `env` package provides two functions for accessing environment variables:

*   `env.Value(key string) string`: Retrieves the value of an environment variable.
*   `env.Must(key string) string`: Retrieves the value of an environment variable. If the variable is not set, the program will panic.

```go
import "github.com/dracory/base/env"

func main() {
    apiKey := env.Value("API_KEY")
    dbPassword := env.Must("DB_PASSWORD")
    fmt.Println("API Key:", apiKey)
    fmt.Println("DB Password:", dbPassword)
}
```

### Special Prefixes

The `env.Value()` function supports special prefixes for processing environment variable values:

*   `base64:`: Decodes the value as a base64 encoded string.
*   `obfuscated:`: Deobfuscates the value using the `envenc` package.

```go
import "github.com/dracory/base/env"

func main() {
    decodedValue := env.Value("BASE64_VALUE") // If BASE64_VALUE is "base64:SGVsbG8gV29ybGQh", it will return "Hello World!"
    obfuscatedValue := env.Value("OBFUSCATED_VALUE") // If OBFUSCATED_VALUE is "obfuscated:some_obfuscated_string", it will deobfuscate the string
    fmt.Println("Decoded Value:", decodedValue)
    fmt.Println("Obfuscated Value:", obfuscatedValue)
}
