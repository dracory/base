package str

import "bytes"

// AddSlashes returns a string with backslashes added before characters that need to be escaped.
//
// These characters are:
// single quote (')
// double quote (")
// backslash (\)
//
func AddSlashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// AddSlashesCustom returns a string with backslashes added before characters specified in the escapeChars string.
func AddSlashesCustom(str string, escapeChars string) string {
	var buf bytes.Buffer
	for _, char := range str {
		if containsRune(escapeChars, char) {
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// containsRune checks if a string contains a specific rune.
func containsRune(str string, r rune) bool {
	for _, char := range str {
		if char == r {
			return true
		}
	}
	return false
}
