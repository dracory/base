package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsEmpty(t *testing.T) {
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
			name:     "Non-Empty String",
			input:    "hello",
			expected: false,
		},
		{
			name:     "String with Spaces",
			input:    " ",
			expected: false,
		},
		{
			name:     "String with Special Characters",
			input:    "!@#$",
			expected: false,
		},
		{
			name:     "Unicode String",
			input:    "привет",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsEmpty(tc.input)
			if result != tc.expected {
				t.Errorf("IsEmpty(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
