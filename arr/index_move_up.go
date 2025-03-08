package arr

// indexMoveUp moves the element at the given index up
//
// Business logic:
// - if the index is first index, will be ignored
// - if the index out of bounds, will be ignored
//
// Parameters:
// - slice: the slice to move the element from
// - index: the index of the element to move
//
// Returns:
// - []T: the new slice
//
// Example:
// 	indices := []int{1, 2, 3, 4, 5}
// 	IndexMoveUp(indices, 2)
// 	fmt.Println(indices) // [1, 3, 2, 4, 5]
func IndexMoveUp[T any](slice []T, index int) []T {
	if index <= 0 || index >= len(slice) {
		return slice // Nothing to move or invalid index
	}

	// Swap the elements at the given index and the element above it
	current := slice[index]
	upper := slice[index-1]

	slice[index] = upper
	slice[index-1] = current

	return slice
}
