package arr

// FilterEmpty takes a slice of strings and returns a new slice with all empty strings removed.
//
// Example:
//  arr.FilterEmpty([]string{"", "hello", "", "world"}) // returns []string{"hello", "world"}
//
// Parameters:
// - slice: the slice of strings to filter.
//
// Returns:
// - []string: a new slice with all empty strings removed.
func FilterEmpty(slice []string) []string {
	if slice == nil {
		return nil // Return nil if the slice is nil.
	}

	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
