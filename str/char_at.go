package str

// CharAt returns the character at the specified index.
//
// If the specified index is negative, it is counted from the end of the string.
//
// Business logic:
// 1. Convert the string to a rune slice for proper handling of UTF-8 encoding.
// 2. Get the length of the rune slice.
// 3. Handle negative indices by converting them to positive indices (e.g. -1 -> length - 1).
// 4. Check if the index is out of bounds.
// 5. If the index is out of bounds, return an empty string.
// 6. Return the character at the specified index.
//
// Example:
//
//	str.CharAt("Hello World", 0) // Returns "H"
//	str.CharAt("Hello World", -1) // Returns "d"
//	str.CharAt("Hello World", 20) // Returns ""
//
// Parameters:
// - str: The string to get the character from.
// - index: The index of the character to get.
//
// Returns:
// - The character at the specified index.
func CharAt(str string, index int) string {
	runes := []rune(str)
	length := len(runes)

	// Handle negative indices
	if index < 0 {
		index += length
		if index < 0 {
			return ""
		}
	}

	// Check bounds
	if index >= length {
		return ""
	}

	// Convert back to string
	return string(runes[index])
}
