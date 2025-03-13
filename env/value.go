package env

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gouniverse/envenc"
)

// Value returns the value for an environment key
//
// Parameters:
//   - key: The environment key
//
// Returns:
//   - The value for the environment key
func Value(key string) string {
	value := os.Getenv(key)

	valueProcessed := envProcess(value)

	return valueProcessed
}

// ValueOr returns the value for an environment key with a default value
//
// Parameters:
//   - key: The environment key
//   - defaultValue: The default value
//
// Returns:
//   - The value for the environment key
func ValueOr(key string, defaultValue string) string {
	value := os.Getenv(key)

	valueProcessed := envProcess(value)

	if valueProcessed == "" {
		return defaultValue
	}

	return valueProcessed
}

// EnvProcess processes the value for an environment key
//
// This function handles base64 and obfuscated prefixes.
//
// Args:
//   - value: The value to process
//
// Returns:
//   - The processed value
func envProcess(value string) string {
	valueTrimmed := strings.TrimSpace(value)

	if strings.HasPrefix(valueTrimmed, "base64:") {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `base64:`)

		valueDecoded, err := base64.URLEncoding.DecodeString(valueNoPrefix)

		if err != nil {
			return err.Error()
		}

		return string(valueDecoded)
	}

	if strings.HasPrefix(valueTrimmed, "obfuscated:") {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `obfuscated:`)

		valueDecoded, err := envenc.Deobfuscate(valueNoPrefix)

		if err != nil {
			return err.Error()
		}

		return string(valueDecoded)
	}

	return valueTrimmed
}
