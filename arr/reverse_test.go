package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestReverse(t *testing.T) {
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
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "even number of elements",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3},
			expected: []int{-3, -2, -1},
		},
		{
			name:     "mixed numbers",
			input:    []int{-1, 0, 1},
			expected: []int{1, 0, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Reverse(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Reverse(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
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
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "even number of elements",
			input:    []string{"a", "b", "c", "d"},
			expected: []string{"d", "c", "b", "a"},
		},
		{
			name:     "mixed strings",
			input:    []string{"hello", "world", "foo"},
			expected: []string{"foo", "world", "hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Reverse(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Reverse(%v) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
