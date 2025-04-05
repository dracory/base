package str

import (
	"encoding/json"
	"strings"
)

// IsMap returns true if the string is a valid Map.
func IsMap(str string) bool {
	if str == "" {
		return false
	}

	if str == "null" {
		return false
	}

	if !strings.HasPrefix(str, "{") {
		return false
	}

	if !strings.HasSuffix(str, "}") {
		return false
	}

	var obj map[string]interface{}
	return json.Unmarshal([]byte(str), &obj) == nil
}
