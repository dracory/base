package config

import (
	"os"
	"strings"

	"github.com/dracory/envenc"
)

// EnvEncConfig represents the environment encryption configuration
type EnvEncConfig struct {
	Used       bool
	PrivateKey string
}

// validateInputs validates the common inputs for both functions
func validateInputs(appEnvironment, publicKey, privateKey string) error {
	if strings.TrimSpace(appEnvironment) == "" {
		return &MissingEnvError{
			Key:     "APP_ENVIRONMENT",
			Context: "required to initialize EnvEnc variables",
		}
	}

	if strings.TrimSpace(publicKey) == "" {
		return &MissingEnvError{
			Key:     "ENV_ENCRYPTION_KEY_PUBLIC",
			Context: "required to initialize EnvEnc variables",
		}
	}

	if strings.TrimSpace(privateKey) == "" {
		return &MissingEnvError{
			Key:     "ENV_ENCRYPTION_KEY_PRIVATE",
			Context: "required to initialize EnvEnc variables",
		}
	}

	return nil
}

// InitializeEnvEncVariablesFromFile initializes environment variables from encrypted vault files
// on the filesystem based on the application environment.
//
// Business logic:
//   - Requires ENV_ENCRYPTION_KEY_PUBLIC and ENV_ENCRYPTION_KEY_PRIVATE environment variables
//   - Looks for vault file named ".env.<app_environment>.vault" in the local filesystem
//   - Derives encryption key from public and private keys using envenc.DeriveKey
//   - Hydrates environment variables from the vault file using envenc.HydrateEnvFromFile
//
// Parameters:
// - appEnvironment: The application environment (e.g., "development", "production", "staging")
// - publicKey: The public encryption key (typically from ENV_ENCRYPTION_KEY_PUBLIC)
// - privateKey: The private encryption key (typically from ENV_ENCRYPTION_KEY_PRIVATE)
//
// Returns:
// - error: If any step fails (missing keys, vault file not found, decryption failed, etc.)
func InitializeEnvEncVariablesFromFile(appEnvironment, publicKey, privateKey string) error {
	// Validate inputs
	if err := validateInputs(appEnvironment, publicKey, privateKey); err != nil {
		return err
	}

	appEnvironment = strings.ToLower(appEnvironment)
	vaultFilePath := ".env." + appEnvironment + ".vault"

	// Check if vault file exists
	if !fileExists(vaultFilePath) {
		return &EnvEncError{
			Operation: "vault_not_found",
			Message:   "Vault file not found: '" + vaultFilePath + "'",
		}
	}

	// Derive encryption key
	derivedKey, err := envenc.DeriveKey(publicKey, privateKey)
	if err != nil {
		return &EnvEncError{
			Operation: "derive_key",
			Message:   "Failed to derive encryption key: " + err.Error(),
		}
	}

	// Load from filesystem
	if err := envenc.HydrateEnvFromFile(vaultFilePath, derivedKey); err != nil {
		return &EnvEncError{
			Operation: "hydrate_from_file",
			Message:   "Failed to hydrate environment from vault file: " + err.Error(),
		}
	}

	return nil
}

// InitializeEnvEncVariablesFromResources initializes environment variables from encrypted vault files
// from embedded resources based on the application environment.
//
// Business logic:
//   - Requires ENV_ENCRYPTION_KEY_PUBLIC and ENV_ENCRYPTION_KEY_PRIVATE environment variables
//   - Looks for vault resource named ".env.<app_environment>.vault" in embedded resources
//   - Derives encryption key from public and private keys using envenc.DeriveKey
//   - Hydrates environment variables from the vault content using envenc.HydrateEnvFromString
//
// Parameters:
// - appEnvironment: The application environment (e.g., "development", "production", "staging")
// - publicKey: The public encryption key (typically from ENV_ENCRYPTION_KEY_PUBLIC)
// - privateKey: The private encryption key (typically from ENV_ENCRYPTION_KEY_PRIVATE)
// - resourceLoader: Function to load embedded resources, returns vault content
//
// Returns:
// - error: If any step fails (missing keys, resource not found, decryption failed, etc.)
func InitializeEnvEncVariablesFromResources(appEnvironment, publicKey, privateKey string, resourceLoader func(string) (string, error)) error {
	// Validate inputs
	if err := validateInputs(appEnvironment, publicKey, privateKey); err != nil {
		return err
	}

	if resourceLoader == nil {
		return &MissingEnvError{
			Key:     "resourceLoader",
			Context: "required to load embedded resources",
		}
	}

	appEnvironment = strings.ToLower(appEnvironment)
	resourceName := ".env." + appEnvironment + ".vault"

	// Load from embedded resources
	vaultContent, err := resourceLoader(resourceName)
	if err != nil {
		return &EnvEncError{
			Operation: "resource_not_found",
			Message:   "Embedded resource not found: '" + resourceName + "' (" + err.Error() + ")",
		}
	}

	if vaultContent == "" {
		return &EnvEncError{
			Operation: "resource_empty",
			Message:   "Embedded resource is empty: '" + resourceName + "'",
		}
	}

	// Derive encryption key
	derivedKey, err := envenc.DeriveKey(publicKey, privateKey)
	if err != nil {
		return &EnvEncError{
			Operation: "derive_key",
			Message:   "Failed to derive encryption key: " + err.Error(),
		}
	}

	// Load from embedded resources
	if err := envenc.HydrateEnvFromString(vaultContent, derivedKey); err != nil {
		return &EnvEncError{
			Operation: "hydrate_from_string",
			Message:   "Failed to hydrate environment from embedded resources: " + err.Error(),
		}
	}

	return nil
}

// fileExists checks if a file exists at the given path
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return !os.IsNotExist(err)
	}
}

// EnvEncError represents an error during environment encryption operations
type EnvEncError struct {
	Operation string
	Message   string
}

func (e *EnvEncError) Error() string {
	return "EnvEnc error in " + e.Operation + ": " + e.Message
}
