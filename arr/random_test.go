package arr_test

import (
	"testing"

	"slices"

	"github.com/dracory/base/arr"
)

func TestRandom(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected bool // We can't predict the exact random value, so we check if a value was returned (not zero)
	}{
		{"empty slice", []int{}, false},
		{"single element", []int{1}, true},
		{"multiple elements", []int{1, 2, 3, 4, 5}, true},
		{"negative numbers", []int{-1, -2, -3}, true},
		{"mixed numbers", []int{-1, 0, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Random(tt.slice)
			// Check if the slice is empty
			if len(tt.slice) == 0 {
				// If the slice is empty, check if the returned value is the zero value of int (0)
				if actual != 0 {
					t.Errorf("arr.Random(%v) = %v, want zero value", tt.slice, actual)
				}
			} else {
				// If the slice is not empty, check if the returned value is one of the elements in the slice
				found := false
				for _, v := range tt.slice {
					if int(actual) == v {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("arr.Random(%v) = %v, want one of the elements in the slice", tt.slice, actual)
				}
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected bool // We can't predict the exact random value, so we check if a value was returned (not zero)
	}{
		{"empty slice", []string{}, false},
		{"single element", []string{"a"}, true},
		{"multiple elements", []string{"a", "b", "c", "d", "e"}, true},
		{"multiple elements", []string{"hello", "world", "foo"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Random(tt.slice)
			// Check if the slice is empty
			if len(tt.slice) == 0 {
				// If the slice is empty, check if the returned value is the zero value of string ("")
				if actual != "" {
					t.Errorf("arr.Random(%v) = %v, want zero value", tt.slice, actual)
				}
			} else {
				// If the slice is not empty, check if the returned value is one of the elements in the slice
				found := slices.Contains(tt.slice, actual)
				if !found {
					t.Errorf("arr.Random(%v) = %v, want one of the elements in the slice", tt.slice, actual)
				}
			}
		})
	}
}
