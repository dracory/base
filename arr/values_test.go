package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []int{1, 2, 3}, // Order is not guaranteed
		},
		{
			name:     "duplicate values",
			input:    map[string]int{"a": 1, "b": 1, "c": 1},
			expected: []int{1, 1, 1}, // Order is not guaranteed
		},
		{
			name:     "mixed values",
			input:    map[string]int{"a": 1, "b": -2, "c": 0},
			expected: []int{1, -2, 0}, // Order is not guaranteed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Values(tt.input)

			// Since the order of map values is not guaranteed, we need to sort both slices before comparing
			if !areSlicesEquivalent(actual, tt.expected) {
				t.Errorf("arr.Values(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

// areSlicesEquivalent checks if two slices contain the same elements, regardless of order.
func areSlicesEquivalent(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count the occurrences of each element
	countA := make(map[int]int)
	countB := make(map[int]int)

	for _, val := range a {
		countA[val]++
	}
	for _, val := range b {
		countB[val]++
	}

	// Compare the counts
	return reflect.DeepEqual(countA, countB)
}
