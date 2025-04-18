package arr

// indexMoveDown moves the element at the given index down
//
// Business logic:
// - if the index is last index, will be ignored
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
//  arr := []int{1, 2, 3, 4}
//  result := IndexMoveDown(arr, 1)
//  // result is now [1, 4, 2, 3]
func IndexMoveDown[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice)-1 {
		return slice // Nothing to move or invalid index
	}

	// Get the current element and the element below it
	current := slice[index]
	lower := slice[index+1]

	// Swap the two elements
	slice[index] = lower
	slice[index+1] = current

	return slice
}
