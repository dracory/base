package env

import (
	"strconv"
)

func getEnvBool(key string, defaultValue bool) bool {
	valueStr := Value(key)

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

// Bool returns the value for an environment key
//
// Parameters:
//   - key: The environment key
//
// Returns:
//   - The value for the environment key, false if not set
func Bool(key string) bool {
	return getEnvBool(key, false)
}

// BoolDefault returns the value for an environment key with a default value
//
// Parameters:
//   - key: The environment key
//   - defaultValue: The default value
//
// Returns:
//   - The value for the environment key
func BoolDefault(key string, defaultValue bool) bool {
	return getEnvBool(key, defaultValue)
}
