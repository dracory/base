package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected bool
	}{
		{
			name:     "empty b",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: true,
		},
		{
			name:     "empty a",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "all elements present",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{3, 1, 5},
			expected: true,
		},
		{
			name:     "not all elements present",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{3, 1, 6},
			expected: false,
		},
		{
			name:     "duplicates in b",
			a:        []int{1, 2, 3, 3, 4, 5},
			b:        []int{3, 1, 3},
			expected: true,
		},
		{
			name:     "duplicates in b not enough",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{3, 1, 3},
			expected: false,
		},
		{
			name:     "duplicates in a",
			a:        []int{1, 1, 2, 3, 4, 5},
			b:        []int{3, 1},
			expected: true,
		},
		{
			name:     "a shorter than b",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "same elements different order",
			a:        []int{1, 2, 3},
			b:        []int{3, 2, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.ContainsAll(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("ContainsAll(%v, %v) = %v, want %v", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

func TestContainsAllString(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected bool
	}{
		{
			name:     "empty b",
			a:        []string{"a", "b", "c"},
			b:        []string{},
			expected: true,
		},
		{
			name:     "empty a",
			a:        []string{},
			b:        []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "all elements present",
			a:        []string{"a", "b", "c", "d", "e"},
			b:        []string{"c", "a", "e"},
			expected: true,
		},
		{
			name:     "not all elements present",
			a:        []string{"a", "b", "c", "d", "e"},
			b:        []string{"c", "a", "f"},
			expected: false,
		},
		{
			name:     "duplicates in b",
			a:        []string{"a", "b", "c", "c", "d", "e"},
			b:        []string{"c", "a", "c"},
			expected: true,
		},
		{
			name:     "duplicates in b not enough",
			a:        []string{"a", "b", "c", "d", "e"},
			b:        []string{"c", "a", "c"},
			expected: false,
		},
		{
			name:     "duplicates in a",
			a:        []string{"a", "a", "b", "c", "d", "e"},
			b:        []string{"c", "a"},
			expected: true,
		},
		{
			name:     "a shorter than b",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "c", "d"},
			expected: false,
		},
		{
			name:     "same elements different order",
			a:        []string{"a", "b", "c"},
			b:        []string{"c", "b", "a"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.ContainsAll(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("ContainsAll(%v, %v) = %v, want %v", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}
