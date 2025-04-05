package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Map - Empty",
			input:    "{}",
			expected: true,
		},
		{
			name:     "Valid Map - Simple",
			input:    `{"key": "value"}`,
			expected: true,
		},
		{
			name:     "Valid Map - Numbers",
			input:    `{"key": 123}`,
			expected: true,
		},
		{
			name:     "Valid Map - Boolean",
			input:    `{"key": true}`,
			expected: true,
		},
		{
			name:     "Valid Map - Null",
			input:    `{"key": null}`,
			expected: true,
		},
		{
			name:     "Valid Map - Nested Map",
			input:    `{"key": {"nestedKey": "nestedValue"}}`,
			expected: true,
		},
		{
			name:     "Valid Map - Nested Slice",
			input:    `{"key": [1, 2, 3]}`,
			expected: true,
		},
		{
			name:     "Valid Map - Multiple Keys",
			input:    `{"key1": "value1", "key2": "value2"}`,
			expected: true,
		},
		{
			name:     "Invalid Map - Missing Bracket",
			input:    `{"key": "value"`,
			expected: false,
		},
		{
			name:     "Invalid Map - Extra Bracket",
			input:    `{"key": "value"}}`,
			expected: false,
		},
		{
			name:     "Invalid Map - Invalid JSON",
			input:    `{"key":}`,
			expected: false,
		},
		{
			name:     "Invalid Map - Not a Map",
			input:    `[1, 2, 3]`,
			expected: false,
		},
		{
			name:     "Invalid Map - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid Map - Null",
			input:    "null",
			expected: false,
		},
		{
			name:     "Invalid Map - Number",
			input:    "123",
			expected: false,
		},
		{
			name:     "Invalid Map - String",
			input:    `"hello"`,
			expected: false,
		},
		{
			name:     "Invalid Map - Boolean",
			input:    "true",
			expected: false,
		},
		{
			name:     "Valid Map - Float",
			input:    `{"key": 1.1}`,
			expected: true,
		},
		{
			name:     "Valid Map - Unicode",
			input:    `{"привет": "мир"}`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsMap(tc.input)
			if result != tc.expected {
				t.Errorf("IsMap(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
