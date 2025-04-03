package str

import "strings"

// Before returns the substring before the first occurrence of the specified search string.
func Before(str, needle string) string {
	if needle == "" {
		return str
	}
	
	index := strings.Index(str, needle)
	if index != -1 {
		return str[:index]
	}
	return str
}
