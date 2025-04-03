package str

import "strings"

// After returns the substring after the first occurrence of the specified needle.
func After(str, needle string) string {
	if needle == "" {
		return str
	}

	index := strings.Index(str, needle)
	if index != -1 {
		str = str[index+len(needle):]
	}
	return str
}
