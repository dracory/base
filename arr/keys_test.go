package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []string{"a", "b", "c"}, // Order is not guaranteed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Keys(tt.input)
			if !areSlicesEquivalentString(actual, tt.expected) {
				t.Errorf("arr.Keys(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func areSlicesEquivalentString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count the occurrences of each element
	countA := make(map[string]int)
	countB := make(map[string]int)

	for _, val := range a {
		countA[val]++
	}
	for _, val := range b {
		countB[val]++
	}

	// Compare the counts
	return reflect.DeepEqual(countA, countB)
}
