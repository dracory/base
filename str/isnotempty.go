package str

// IsNotEmpty returns true if the string is not empty.
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
