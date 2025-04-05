package str

import "encoding/json"

// IsSlice returns true if the string is a valid Slice.
func IsSlice(str string) bool {
	if str == "" {
		return false
	}

	if str == "null" {
		return false
	}

	var arr []interface{}
	return json.Unmarshal([]byte(str), &arr) == nil
}
