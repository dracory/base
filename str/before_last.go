package str

import "strings"

// BeforeLast returns the substring before the last occurrence of the specified search string.
func BeforeLast(str, needle string) string {
	if needle == "" {
		return str
	}

	index := strings.LastIndex(str, needle)
	if index != -1 {
		return str[:index]
	}
	return str
}
