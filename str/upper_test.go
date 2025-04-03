package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestUpper tests the Upper function with various scenarios
func TestUpper(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// Basic cases
		{"", ""},
		{"hello", "HELLO"},
		{"Hello", "HELLO"},
		{"HELLO", "HELLO"},
		{"hElLo", "HELLO"},

		// Multiple words
		{"hello world", "HELLO WORLD"},
		{"Hello World", "HELLO WORLD"},
		{"HELLO WORLD", "HELLO WORLD"},

		// With punctuation
		{"hello, world!", "HELLO, WORLD!"},
		{"Hello-World", "HELLO-WORLD"},
		{"Hello.World", "HELLO.WORLD"},

		// Numbers
		{"hello123", "HELLO123"},
		{"123hello", "123HELLO"},
		{"123hello456", "123HELLO456"},

		// Cyrillic
		{"привет", "ПРИВЕТ"},
		{"Привет", "ПРИВЕТ"},
		{"ПРИВЕТ", "ПРИВЕТ"},
		{"привет мир", "ПРИВЕТ МИР"},

		// Special characters
		{"hello!@#$%^&*()", "HELLO!@#$%^&*()"},
		{"hello\tworld", "HELLO\tWORLD"},
		{"hello\nworld", "HELLO\nWORLD"},
	}

	for _, tc := range testCases {
		result := str.Upper(tc.input)
		if result != tc.expected {
			t.Errorf("Upper(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
