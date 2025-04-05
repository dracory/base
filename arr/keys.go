package arr

import (
	"github.com/samber/lo"
)

// Keys creates an array of the map keys.
func Keys[K comparable, V any](in map[K]V) []K {
	return lo.Keys(in)
}
