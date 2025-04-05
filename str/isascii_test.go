package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsAscii(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: true,
		},
		{
			name:     "ASCII String",
			input:    "hello world",
			expected: true,
		},
		{
			name:     "ASCII String with Numbers",
			input:    "hello123world",
			expected: true,
		},
		{
			name:     "ASCII String with Special Characters",
			input:    "hello!@#$%^&*()_+",
			expected: true,
		},
		{
			name:     "Non-ASCII String",
			input:    "привет мир",
			expected: false,
		},
		{
			name:     "Mixed String",
			input:    "hello мир",
			expected: false,
		},
		{
			name:     "Mixed String with Special Characters",
			input:    "hello!@#$мир",
			expected: false,
		},
		{
			name:     "Control Character",
			input:    string([]byte{0x07}), // Bell character
			expected: true,
		},
		{
			name:     "Extended ASCII",
			input:    string([]byte{0x80}), // Extended ASCII character
			expected: false,
		},
		{
			name:     "Null Character",
			input:    string([]byte{0x00}),
			expected: true,
		},
		{
			name:     "Delete Character",
			input:    string([]byte{0x7F}),
			expected: true,
		},
		{
			name:     "Tab Character",
			input:    string([]byte{0x09}),
			expected: true,
		},
		{
			name:     "Newline Character",
			input:    string([]byte{0x0A}),
			expected: true,
		},
		{
			name:     "Carriage Return Character",
			input:    string([]byte{0x0D}),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsAscii(tc.input)
			if result != tc.expected {
				t.Errorf("IsAscii(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
