# System Patterns

The project uses Go for backend development. The `bbcode` package is responsible for converting BBCode to HTML. The tests are written in Go and use the `testing` package.

The project includes functionalities for:
*   BBCode to HTML conversion
*   Slice manipulation
*   Database interaction
*   Error handling and validation
*   Image manipulation
*   URL downloading
*   QR code generation
*   HTTP request handling
*   Timezone conversion
*   Date and datetime validation
*   Web server functionality

## Environment Variables

The project uses environment variables for configuration. The `env` package provides utilities for loading and accessing these variables.

*   The `env/initialize.go` file uses the `github.com/joho/godotenv` library to load environment variables from `.env` files.
*   The `env/value.go` file provides a `Value` function to retrieve the value of an environment variable using `os.Getenv()`. It also includes processing for `base64:` and `obfuscated:` prefixes.
*   The `env/must.go` file provides a `Must` function, which retrieves an environment variable and panics if the variable is not set.
