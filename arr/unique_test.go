package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestUnique(t *testing.T) {
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
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "duplicates",
			input:    []int{1, 2, 2, 3, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "mixed duplicates",
			input:    []int{1, 2, 1, 3, 2, 4, 3, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -2, -3, -1},
			expected: []int{-1, -2, -3},
		},
		{
			name:     "zero and negative numbers",
			input:    []int{0, -1, -2, 0, -1},
			expected: []int{0, -1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Unique(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Unique(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
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
			name:     "no duplicates",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "duplicates",
			input:    []string{"a", "b", "b", "c", "d", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "all duplicates",
			input:    []string{"a", "a", "a", "a", "a"},
			expected: []string{"a"},
		},
		{
			name:     "mixed duplicates",
			input:    []string{"a", "b", "a", "c", "b", "d", "c", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Unique(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Unique(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
