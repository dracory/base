package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestAddSlashes(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"hello'world", "hello\\'world"},
		{"hello\"world", "hello\\\"world"},
		{"hello\\world", "hello\\\\world"},
		{"hello'\"\\world", "hello\\'\\\"\\\\world"},
		{"", ""},
		{"'\"\\", "\\'\\\"\\\\"},
		{"no special chars", "no special chars"},
		{"'single' \"double\" \\backslash\\", "\\'single\\' \\\"double\\\" \\\\backslash\\\\"},
	}

	for _, tc := range testCases {
		result := str.AddSlashes(tc.input)
		if result != tc.expected {
			t.Errorf("AddSlashes(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
