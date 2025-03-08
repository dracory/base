package arr

import (
	"testing"
)

func TestIndexMoveDown(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{"move down", []int{1, 2, 3, 4}, 1, []int{1, 3, 2, 4}},
		{"last element", []int{1, 2, 3, 4}, 3, []int{1, 2, 3, 4}},
		{"out of bounds", []int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{"empty slice", []int{}, 0, []int{}},
		{"one element", []int{1}, 0, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndexMoveDown(tt.slice, tt.index)
			if !equal(result, tt.expected) {
				t.Errorf("IndexMoveDown(%v, %d) = %v, want %v", tt.slice, tt.index, result, tt.expected)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
