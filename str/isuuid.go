package str

import "strings"

// IsUuid returns true if the string is a valid UUID.
func IsUuid(str string) bool {
	if str == "" {
		return false
	}

	str = strings.ToLower(str)

	if len(str) != 36 && len(str) != 32 {
		return false
	}

	if len(str) == 36 && !ContainsOnly(str, "0123456789abcdef-") {
		return false
	}

	if len(str) == 32 && !ContainsOnly(str, "0123456789abcdef") {
		return false
	}

	if len(str) == 36 {
		return IsMatch(str, `(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	}

	return IsMatch(str, `(?i)^[0-9a-f]{32}$`)
}
