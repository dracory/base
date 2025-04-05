package arr

// Count counts the number of elements in the collection.
func Count[T comparable](collection []T) (count int) {
	return len(collection)
}
