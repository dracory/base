package env

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gouniverse/envenc"
)

// Value returns the value for an environment key
//
// If the value is not found, or if the value cannot be processed,
// returns an empty string.
//
// If you want a default value, use ValueOrDefault.
// If you want an error, use ValueOrError.
// If you want a panic, use ValueOrPanic.
//
// Parameters:
//   - key: The environment key
//
// Returns:
//   - The value for the environment key
func Value(key string) string {
	value := os.Getenv(key)

	valueProcessed, err := envProcess(value)

	if valueProcessed == "" || err != nil {
		return ""
	}

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
func ValueOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	valueProcessed, err := envProcess(value)

	if valueProcessed == "" || err != nil {
		return defaultValue
	}

	return valueProcessed
}

func ValueOrError(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", errors.New("value not found")
	}

	valueProcessed, err := envProcess(value)

	return valueProcessed, err
}

func ValueOrPanic(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Panicf("Environment variable %s is required but not set", key)
	}

	valueProcessed, err := envProcess(value)

	if valueProcessed == "" || err != nil {
		log.Panicf("Environment variable %s is required but not set", key)
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
func envProcess(value string) (string, error) {
	valueTrimmed := strings.TrimSpace(value)
	isBase64 := strings.HasPrefix(valueTrimmed, "base64:")
	isObfuscated := strings.HasPrefix(valueTrimmed, "obfuscated:")

	if isBase64 {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `base64:`)

		valueDecoded, err := base64.URLEncoding.DecodeString(valueNoPrefix)

		if err != nil {
			return "", err
		}

		return string(valueDecoded), nil
	}

	if isObfuscated {
		valueNoPrefix := strings.TrimPrefix(valueTrimmed, `obfuscated:`)

		valueDecoded, err := envenc.Deobfuscate(valueNoPrefix)

		if err != nil {
			return "", err
		}

		return string(valueDecoded), nil
	}

	return valueTrimmed, nil
}
