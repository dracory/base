package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected bool
	}{
		{
			name:     "empty slice",
			slice:    []int{},
			item:     1,
			expected: false,
		},
		{
			name:     "item present",
			slice:    []int{1, 2, 3, 4, 5},
			item:     3,
			expected: true,
		},
		{
			name:     "item not present",
			slice:    []int{1, 2, 3, 4, 5},
			item:     6,
			expected: false,
		},
		{
			name:     "item at start",
			slice:    []int{1, 2, 3, 4, 5},
			item:     1,
			expected: true,
		},
		{
			name:     "item at end",
			slice:    []int{1, 2, 3, 4, 5},
			item:     5,
			expected: true,
		},
		{
			name:     "negative numbers",
			slice:    []int{-1, -2, -3},
			item:     -2,
			expected: true,
		},
		{
			name:     "mixed numbers",
			slice:    []int{-1, 0, 1},
			item:     0,
			expected: true,
		},
		{
			name:     "mixed numbers not present",
			slice:    []int{-1, 0, 1},
			item:     2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Contains(tt.slice, tt.item)
			if actual != tt.expected {
				t.Errorf("arr.Contains(%v, %v) = %v, want %v", tt.slice, tt.item, actual, tt.expected)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{
			name:     "empty slice",
			slice:    []string{},
			item:     "a",
			expected: false,
		},
		{
			name:     "item present",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "c",
			expected: true,
		},
		{
			name:     "item not present",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "f",
			expected: false,
		},
		{
			name:     "item at start",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "a",
			expected: true,
		},
		{
			name:     "item at end",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "e",
			expected: true,
		},
		{
			name:     "mixed case",
			slice:    []string{"a", "B", "c"},
			item:     "B",
			expected: true,
		},
		{
			name:     "mixed case not present",
			slice:    []string{"a", "B", "c"},
			item:     "b",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Contains(tt.slice, tt.item)
			if actual != tt.expected {
				t.Errorf("arr.Contains(%v, %v) = %v, want %v", tt.slice, tt.item, actual, tt.expected)
			}
		})
	}
}
