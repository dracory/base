package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsNotEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Non-Empty String",
			input:    "hello",
			expected: true,
		},
		{
			name:     "String with Spaces",
			input:    " ",
			expected: true,
		},
		{
			name:     "String with Special Characters",
			input:    "!@#$",
			expected: true,
		},
		{
			name:     "Unicode String",
			input:    "привет",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsNotEmpty(tc.input)
			if result != tc.expected {
				t.Errorf("IsNotEmpty(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
