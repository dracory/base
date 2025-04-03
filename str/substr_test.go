package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestSubstr tests the Substr function with various scenarios
func TestSubstr(t *testing.T) {
	testCases := []struct {
		input    string
		start    int
		length   []int
		expected string
	}{
		{"Hello World", 0, []int{}, "Hello World"},
		{"Hello World", 0, []int{5}, "Hello"},
		{"Hello World", 6, []int{}, "World"},
		{"Hello World", 6, []int{5}, "World"},
		{"Hello World", -6, []int{}, " World"},
		{"Hello World", -6, []int{5}, " Worl"},
		{"Hello World", 0, []int{-5}, "Hello "},
		{"Hello World", 6, []int{-5}, ""},
		{"Hello World", -6, []int{-5}, " "},
		{"Hello World", 10, []int{}, "d"},
		{"Hello World", 10, []int{5}, "d"},
		{"Hello World", -11, []int{}, "Hello World"},
		{"Hello World", -11, []int{5}, "Hello"},
		{"Привет Мир", 0, []int{}, "Привет Мир"},
		{"Привет Мир", 0, []int{6}, "Привет"},
		{"Привет Мир", 7, []int{}, "Мир"},
		{"Привет Мир", 7, []int{3}, "Мир"},
		{"Привет Мир", -3, []int{}, "Мир"},
		{"Привет Мир", -3, []int{3}, "Мир"},
		{"Привет Мир", 0, []int{-3}, "Привет "},
		{"Привет Мир", 7, []int{-3}, ""},
		{"Привет Мир", -3, []int{-3}, ""},
		{"Привет Мир", 10, []int{}, ""},
		{"Привет Мир", 10, []int{5}, ""},
		{"Привет Мир", -11, []int{}, "Привет Мир"},
		{"Привет Мир", -11, []int{6}, "Привет"},
	}

	for _, tc := range testCases {
		result := str.Substr(tc.input, tc.start, tc.length...)
		if result != tc.expected {
			t.Errorf("Substr(%q, %d, %v) = %q, want %q", tc.input, tc.start, tc.length, result, tc.expected)
		}
	}
}
