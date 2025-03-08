package arr

// IndexRemove removes the element at the given index from the slice.
//
// Parameters:
// - slice: the slice to remove the element from.
// - index: the index of the element to remove.
//
// Returns:
// - []T: a new slice with the element at the given index removed.
//
// Business Logic:
// - If the index is out of bounds, the original slice is returned unchanged.
// - This function does not panic on an out-of-bounds index.
//
// Example:
//  arr := []int{1, 2, 3, 4}
//  result := IndexRemove(arr, 2)
//  // result is now [1, 2, 4]
func IndexRemove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice // Return original slice if index is invalid
	}

	// Remove the element at the given index by creating a new slice
	return append(slice[:index], slice[index+1:]...)
}
