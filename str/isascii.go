package str

// IsAscii returns true if the string contains only ASCII characters.
func IsAscii(str string) bool {
	for _, r := range str {
		if r > 127 {
			return false
		}
	}
	return true
}
