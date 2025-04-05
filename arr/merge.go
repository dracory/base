package arr

// Merge merges multiple slices into a single slice.
func Merge[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}

	var result []T

	for _, s := range slices {
		if s == nil {
			continue
		}
		result = append(result, s...)
	}

	return result
}
