package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Slice - Empty",
			input:    "[]",
			expected: true,
		},
		{
			name:     "Valid Slice - Numbers",
			input:    "[1, 2, 3]",
			expected: true,
		},
		{
			name:     "Valid Slice - Strings",
			input:    `["hello", "world"]`,
			expected: true,
		},
		{
			name:     "Valid Slice - Mixed Types",
			input:    `[1, "hello", true]`,
			expected: true,
		},
		{
			name:     "Valid Slice - Nested Slice",
			input:    `[1, [2, 3]]`,
			expected: true,
		},
		{
			name:     "Valid Slice - Nested Map",
			input:    `[1, {"key": "value"}]`,
			expected: true,
		},
		{
			name:     "Invalid Slice - Missing Bracket",
			input:    "[1, 2, 3",
			expected: false,
		},
		{
			name:     "Invalid Slice - Extra Bracket",
			input:    "[1, 2, 3]]",
			expected: false,
		},
		{
			name:     "Invalid Slice - Invalid JSON",
			input:    "[1, 2,]",
			expected: false,
		},
		{
			name:     "Invalid Slice - Not a Slice",
			input:    `{"key": "value"}`,
			expected: false,
		},
		{
			name:     "Invalid Slice - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid Slice - Null",
			input:    "null",
			expected: false,
		},
		{
			name:     "Invalid Slice - Number",
			input:    "123",
			expected: false,
		},
		{
			name:     "Invalid Slice - String",
			input:    `"hello"`,
			expected: false,
		},
		{
			name:     "Invalid Slice - Boolean",
			input:    "true",
			expected: false,
		},
		{
			name:     "Valid Slice - Boolean",
			input:    "[true, false]",
			expected: true,
		},
		{
			name:     "Valid Slice - Null",
			input:    "[null, null]",
			expected: true,
		},
		{
			name:     "Valid Slice - Float",
			input:    "[1.1, 2.2, 3.3]",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsSlice(tc.input)
			if result != tc.expected {
				t.Errorf("IsSlice(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
