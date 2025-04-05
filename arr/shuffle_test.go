package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{}, // We can't predict the exact order, so we'll check for length and content later
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3},
			expected: []int{}, // We can't predict the exact order, so we'll check for length and content later
		},
		{
			name:     "mixed numbers",
			input:    []int{-1, 0, 1},
			expected: []int{}, // We can't predict the exact order, so we'll check for length and content later
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Shuffle(tt.input)

			// Check if the length is the same
			if len(actual) != len(tt.input) {
				t.Errorf("arr.Shuffle(%v) length = %d, want %d", tt.input, len(actual), len(tt.input))
			}

			// Check if the content is the same (ignoring order)
			if len(tt.input) > 1 {
				if !areSlicesEquivalentInt(actual, tt.input) {
					t.Errorf("arr.Shuffle(%v) = %v, want same elements", tt.input, actual)
				}
			} else {
				if !reflect.DeepEqual(actual, tt.input) {
					t.Errorf("arr.Shuffle(%v) = %v, want %v", tt.input, actual, tt.input)
				}
			}
		})
	}
}

func TestShuffleString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{}, // We can't predict the exact order, so we'll check for length and content later
		},
		{
			name:     "mixed strings",
			input:    []string{"hello", "world", "foo"},
			expected: []string{}, // We can't predict the exact order, so we'll check for length and content later
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Shuffle(tt.input)

			// Check if the length is the same
			if len(actual) != len(tt.input) {
				t.Errorf("arr.Shuffle(%v) length = %d, want %d", tt.input, len(actual), len(tt.input))
			}

			// Check if the content is the same (ignoring order)
			if len(tt.input) > 1 {
				if !areSlicesEquivalentString(actual, tt.input) {
					t.Errorf("arr.Shuffle(%v) = %v, want same elements", tt.input, actual)
				}
			} else {
				if !reflect.DeepEqual(actual, tt.input) {
					t.Errorf("arr.Shuffle(%v) = %v, want %v", tt.input, actual, tt.input)
				}
			}
		})
	}
}

func areSlicesEquivalentInt(a, b []int) bool {
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
