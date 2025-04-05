package arr

import (
	"github.com/samber/lo/mutable"
)

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](collection []T) []T {
	mutable.Reverse(collection)
	return collection
}
