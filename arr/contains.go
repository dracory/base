package arr

import "slices"

func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}
