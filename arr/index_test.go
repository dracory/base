package arr

import (
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected int
	}{
		{
			name:     "empty slice",
			slice:    []int{},
			item:     1,
			expected: -1,
		},
		{
			name:     "item present",
			slice:    []int{1, 2, 3, 4, 5},
			item:     3,
			expected: 2,
		},
		{
			name:     "item not present",
			slice:    []int{1, 2, 3, 4, 5},
			item:     6,
			expected: -1,
		},
		{
			name:     "item at start",
			slice:    []int{1, 2, 3, 4, 5},
			item:     1,
			expected: 0,
		},
		{
			name:     "item at end",
			slice:    []int{1, 2, 3, 4, 5},
			item:     5,
			expected: 4,
		},
		{
			name:     "negative numbers",
			slice:    []int{-1, -2, -3},
			item:     -2,
			expected: 1,
		},
		{
			name:     "mixed numbers",
			slice:    []int{-1, 0, 1},
			item:     0,
			expected: 1,
		},
		{
			name:     "mixed numbers not present",
			slice:    []int{-1, 0, 1},
			item:     2,
			expected: -1,
		},
		{
			name:     "duplicate item",
			slice:    []int{1, 2, 3, 3, 4, 5},
			item:     3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Index(tt.slice, tt.item)
			if actual != tt.expected {
				t.Errorf("Index(%v, %v) = %v, want %v", tt.slice, tt.item, actual, tt.expected)
			}
		})
	}
}

func TestIndexString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected int
	}{
		{
			name:     "empty slice",
			slice:    []string{},
			item:     "a",
			expected: -1,
		},
		{
			name:     "item present",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "c",
			expected: 2,
		},
		{
			name:     "item not present",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "f",
			expected: -1,
		},
		{
			name:     "item at start",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "a",
			expected: 0,
		},
		{
			name:     "item at end",
			slice:    []string{"a", "b", "c", "d", "e"},
			item:     "e",
			expected: 4,
		},
		{
			name:     "mixed case",
			slice:    []string{"a", "B", "c"},
			item:     "B",
			expected: 1,
		},
		{
			name:     "mixed case not present",
			slice:    []string{"a", "B", "c"},
			item:     "b",
			expected: -1,
		},
		{
			name:     "duplicate item",
			slice:    []string{"a", "b", "c", "c", "d", "e"},
			item:     "c",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Index(tt.slice, tt.item)
			if actual != tt.expected {
				t.Errorf("Index(%v, %v) = %v, want %v", tt.slice, tt.item, actual, tt.expected)
			}
		})
	}
}
