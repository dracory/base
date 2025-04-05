package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected bool
	}{
		{
			name:     "empty slices",
			a:        []int{},
			b:        []int{},
			expected: true,
		},
		{
			name:     "equal slices",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "different lengths",
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
			expected: false,
		},
		{
			name:     "different elements",
			a:        []int{1, 2, 3},
			b:        []int{1, 3, 2},
			expected: false,
		},
		{
			name:     "different elements at start",
			a:        []int{1, 2, 3},
			b:        []int{4, 2, 3},
			expected: false,
		},
		{
			name:     "different elements at end",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 4},
			expected: false,
		},
		{
			name:     "negative numbers",
			a:        []int{-1, -2, -3},
			b:        []int{-1, -2, -3},
			expected: true,
		},
		{
			name:     "mixed numbers",
			a:        []int{-1, 0, 1},
			b:        []int{-1, 0, 1},
			expected: true,
		},
		{
			name:     "mixed numbers different",
			a:        []int{-1, 0, 1},
			b:        []int{-1, 0, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Equals(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("arr.Equals(%v, %v) = %v, want %v", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

func TestEqualsString(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected bool
	}{
		{
			name:     "empty slices",
			a:        []string{},
			b:        []string{},
			expected: true,
		},
		{
			name:     "equal slices",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "different lengths",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b"},
			expected: false,
		},
		{
			name:     "different elements",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "c", "b"},
			expected: false,
		},
		{
			name:     "different elements at start",
			a:        []string{"a", "b", "c"},
			b:        []string{"d", "b", "c"},
			expected: false,
		},
		{
			name:     "different elements at end",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "d"},
			expected: false,
		},
		{
			name:     "mixed case",
			a:        []string{"a", "B", "c"},
			b:        []string{"a", "B", "c"},
			expected: true,
		},
		{
			name:     "mixed case different",
			a:        []string{"a", "B", "c"},
			b:        []string{"a", "b", "c"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Equals(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("arr.Equals(%v, %v) = %v, want %v", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}
