package str

import "strings"

// AfterLast returns the substring after the last occurrence of the specified needle.
func AfterLast(str, needle string) string {
	index := strings.LastIndex(str, needle)
	if index != -1 {
		str = str[index+len(needle):]
	}

	return str
}
