package str

import "strings"

// Words returns the string truncated to the given number of words.
// If limit is less than 1, it returns an empty string.
// If limit is greater than or equal to the number of words, it returns the full string.
// An optional end string can be provided to customize the truncation suffix.
func Words(str string, limit int, end ...string) string {
	if limit < 1 {
		return ""
	}

	defaultEnd := "..."
	if len(end) > 0 {
		defaultEnd = end[0]
	}

	words := strings.Fields(str)
	if len(words) <= limit {
		return str
	}

	return strings.Join(words[:limit], " ") + defaultEnd
}
