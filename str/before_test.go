package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestBefore tests the Before function with various scenarios
func TestBefore(t *testing.T) {
	testCases := []struct {
		input    string
		needle   string
		expected string
	}{
		{"Hello World", "World", "Hello "},
		{"Hello World", "Hello", ""},
		{"Hello World", "W", "Hello "},
		{"Hello World", "X", "Hello World"},
		{"Hello World", "", "Hello World"},
		{"", "World", ""},
		{"", "", ""},
		{"Hello World", "Hello World", ""},
		{"Hello World", "World Hello", "Hello World"},
		{"Hello World World", "World", "Hello "},
		{"Hello World World", "World World", "Hello "},
	}

	for _, tc := range testCases {
		result := str.Before(tc.input, tc.needle)
		if result != tc.expected {
			t.Errorf("Before(%q, %q) = %q, want %q", tc.input, tc.needle, result, tc.expected)
		}
	}
}
