package str

// IsUlid returns true if the string is a valid ULID.
func IsUlid(str string) bool {
	return IsMatch(str, `^[0-9A-Z]{26}$`)
}
