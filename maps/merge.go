package maps

import "github.com/samber/lo"

// Merge merges multiple maps from left to right.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	return lo.Assign(maps...)
}
