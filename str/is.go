package str

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Is returns true if the string matches any of the given patterns.
func Is(str string, patterns ...string) bool {
	for _, pattern := range patterns {
		if pattern == str {
			return true
		}

		// Escape special characters in the pattern
		pattern = regexp.QuoteMeta(pattern)

		// Replace asterisks with regular expression wildcards
		pattern = strings.ReplaceAll(pattern, `\*`, ".*")

		// Create a regular expression pattern for matching
		regexPattern := "^" + pattern + "$"

		// Compile the regular expression
		regex := regexp.MustCompile(regexPattern)

		// Check if the value matches the pattern
		if regex.MatchString(str) {
			return true
		}
	}

	return false
}

// IsEmpty returns true if the string is empty.
func IsEmpty(str string) bool {
	return str == ""
}

// IsNotEmpty returns true if the string is not empty.
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsAscii returns true if the string contains only ASCII characters.
func IsAscii(str string) bool {
	return IsMatch(str, `^[\\x00-\\x7F]+$`)
}

// IsMap returns true if the string is a valid Map.
func IsMap(str string) bool {
	var obj map[string]interface{}
	return json.Unmarshal([]byte(str), &obj) == nil
}

// IsSlice returns true if the string is a valid Slice.
func IsSlice(str string) bool {
	var arr []interface{}
	return json.Unmarshal([]byte(str), &arr) == nil
}

// IsUlid returns true if the string is a valid ULID.
func IsUlid(str string) bool {
	return IsMatch(str, `^[0-9A-Z]{26}$`)
}

// IsUuid returns true if the string is a valid UUID.
func IsUuid(str string) bool {
	return IsMatch(str, `(?i)^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
}
