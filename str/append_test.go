package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestAppend tests the Append function with various scenarios
func TestAppend(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"Hello"}, "Hello"},
		{[]string{"Hello", "World"}, "HelloWorld"},
		{[]string{"Hello", " ", "World"}, "Hello World"},
		{[]string{"Hello", "World", "!"}, "HelloWorld!"},
		{[]string{"Hello", "", "World"}, "HelloWorld"},
		{[]string{"Hello", "  ", "World"}, "Hello  World"},
		{[]string{"Hello", "\n", "World"}, "Hello\nWorld"},
		{[]string{"Hello", "\t", "World"}, "Hello\tWorld"},
		{[]string{"Hello", "\r", "World"}, "Hello\rWorld"},
		{[]string{"Hello", "\r\n", "World"}, "Hello\r\nWorld"},
	}

	for _, tc := range testCases {
		result := str.Append(tc.input...)
		if result != tc.expected {
			t.Errorf("Append(%v) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
