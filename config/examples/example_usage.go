package main

import (
	"fmt"
	"os"

	"github.com/dracory/base/config"
)

// Example 1: Basic usage with local vault file
func basicExample() {
	fmt.Println("=== Basic Example ===")

	err := config.InitializeEnvEncVariablesFromFile(
		"development",
		"demo-public-key",
		"demo-private-key",
	)

	if err != nil {
		fmt.Printf("Expected error (no vault file): %v\n", err)
	} else {
		fmt.Println("Environment variables loaded successfully")
	}
}

// Example 2: With embedded resources
func embeddedResourcesExample() {
	fmt.Println("\n=== Embedded Resources Example ===")

	// Simulate a resource loader (in real apps, this would load from embed.FS)
	resourceLoader := func(name string) (string, error) {
		if name == ".env.production.vault" {
			return "encrypted-content-here", nil
		}
		return "", fmt.Errorf("resource not found: %s", name)
	}

	err := config.InitializeEnvEncVariablesFromResources(
		"production",
		"demo-public-key",
		"demo-private-key",
		resourceLoader,
	)

	if err != nil {
		fmt.Printf("Expected error (invalid keys): %v\n", err)
	} else {
		fmt.Println("Environment variables loaded from embedded resources")
	}
}

// Example 3: Error handling examples
func errorHandlingExample() {
	fmt.Println("\n=== Error Handling Examples ===")

	// Missing app environment
	err := config.InitializeEnvEncVariablesFromFile("", "public", "private")
	if err != nil {
		fmt.Printf("Missing app environment: %v\n", err)
	}

	// Missing public key
	err = config.InitializeEnvEncVariablesFromFile("development", "", "private")
	if err != nil {
		fmt.Printf("Missing public key: %v\n", err)
	}

	// Missing private key
	err = config.InitializeEnvEncVariablesFromFile("development", "public", "")
	if err != nil {
		fmt.Printf("Missing private key: %v\n", err)
	}

	// Testing environment (should be skipped)
	err = config.InitializeEnvEncVariablesFromFile("testing", "public", "private")
	if err != nil {
		fmt.Printf("Unexpected error in testing: %v\n", err)
	} else {
		fmt.Println("Testing environment correctly skipped")
	}
}

// Example 4: Integration with environment variables
func integrationExample() {
	fmt.Println("\n=== Integration Example ===")

	// Set up environment variables (normally these would be set externally)
	os.Setenv("APP_ENVIRONMENT", "development")
	os.Setenv("ENV_ENCRYPTION_KEY_PUBLIC", "demo-public-key")
	os.Setenv("ENV_ENCRYPTION_KEY_PRIVATE", "demo-private-key")

	// Load configuration
	appEnv := os.Getenv("APP_ENVIRONMENT")
	publicKey := os.Getenv("ENV_ENCRYPTION_KEY_PUBLIC")
	privateKey := os.Getenv("ENV_ENCRYPTION_KEY_PRIVATE")

	err := config.InitializeEnvEncVariablesFromFile(appEnv, publicKey, privateKey)
	if err != nil {
		fmt.Printf("Configuration loading failed: %v\n", err)
		return
	}

	fmt.Println("Configuration loaded successfully")
	fmt.Printf("App Environment: %s\n", appEnv)

	// Now you can use other environment variables that were loaded from the vault
	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		fmt.Printf("Database Host: %s\n", dbHost)
	} else {
		fmt.Println("DB_HOST not set (vault file not found)")
	}
}

// Example 5: Custom error handling
func customErrorHandlingExample() {
	fmt.Println("\n=== Custom Error Handling Example ===")

	err := config.InitializeEnvEncVariablesFromFile("production", "invalid", "keys")

	if err != nil {
		switch e := err.(type) {
		case *config.MissingEnvError:
			fmt.Printf("Missing environment variable: %s (%s)\n", e.Key, e.Context)
		case *config.EnvEncError:
			fmt.Printf("Vault operation failed in %s: %s\n", e.Operation, e.Message)
		default:
			fmt.Printf("Unexpected error: %v\n", err)
		}
	}
}

func main() {
	fmt.Println("Config Encryption Loader Examples")
	fmt.Println("================================")

	basicExample()
	embeddedResourcesExample()
	errorHandlingExample()
	integrationExample()
	customErrorHandlingExample()

	fmt.Println("\n=== Summary ===")
	fmt.Println("The config encryption loader provides:")
	fmt.Println("- Automatic environment variable loading from encrypted vaults")
	fmt.Println("- Support for both local files and embedded resources")
	fmt.Println("- Environment-specific vault loading")
	fmt.Println("- Comprehensive error handling")
	fmt.Println("- Testing environment support")
	fmt.Println("- Easy integration with existing applications")
}
