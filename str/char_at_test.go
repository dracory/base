package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

// TestCharAt tests the CharAt function with various scenarios
func TestCharAt(t *testing.T) {
	testCases := []struct {
		input    string
		index    int
		expected string
	}{
		{"Hello World", 0, "H"},
		{"Hello World", 1, "e"},
		{"Hello World", 2, "l"},
		{"Hello World", 3, "l"},
		{"Hello World", 4, "o"},
		{"Hello World", 5, " "},
		{"Hello World", 6, "W"},
		{"Hello World", 7, "o"},
		{"Hello World", 8, "r"},
		{"Hello World", 9, "l"},
		{"Hello World", 10, "d"},
		{"Hello World", 11, ""},
		{"Hello World", -1, "d"},
		{"Hello World", -2, "l"},
		{"Hello World", -3, "r"},
		{"Hello World", -4, "o"},
		{"Hello World", -5, "W"},
		{"Hello World", -6, " "},
		{"Hello World", -7, "o"},
		{"Hello World", -8, "l"},
		{"Hello World", -9, "l"},
		{"Hello World", -10, "e"},
		{"Hello World", -11, "H"},
		{"Hello World", -12, ""},
		{"Hello World", 20, ""},
		{"", 0, ""},
		{"", -1, ""},
		{"Привет Мир", 0, "П"},
		{"Привет Мир", 1, "р"},
		{"Привет Мир", 2, "и"},
		{"Привет Мир", 3, "в"},
		{"Привет Мир", 4, "е"},
		{"Привет Мир", 5, "т"},
		{"Привет Мир", 6, " "},
		{"Привет Мир", 7, "М"},
		{"Привет Мир", 8, "и"},
		{"Привет Мир", 9, "р"},
		{"Привет Мир", 10, ""},
		{"Привет Мир", 11, ""},
		{"Привет Мир", -1, "р"},
		{"Привет Мир", -2, "и"},
		{"Привет Мир", -3, "М"},
		{"Привет Мир", -4, " "},
		{"Привет Мир", -5, "т"},
		{"Привет Мир", -6, "е"},
		{"Привет Мир", -7, "в"},
		{"Привет Мир", -8, "и"},
		{"Привет Мир", -9, "р"},
		{"Привет Мир", -10, "П"},
		{"Привет Мир", -11, ""},
		{"Привет Мир", -12, ""},
		{"Привет Мир", 20, ""},
		{"", 0, ""},
		{"", -1, ""},
		{"", 1, ""},
	}

	for _, tc := range testCases {
		result := str.CharAt(tc.input, tc.index)
		if result != tc.expected {
			t.Errorf("CharAt(%q, %d) = %q, want %q", tc.input, tc.index, result, tc.expected)
		}
	}
}
