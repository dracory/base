package arr

// Index finds the index of the first occurrence of an item in a slice.
//
// Returns -1 if the item is not found.
func Index[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}
