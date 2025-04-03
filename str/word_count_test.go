package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestWordCount tests the WordCount function with various scenarios
func TestWordCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		// Basic cases
		{"", 0},
		{"Hello", 1},
		{"Hello World", 2},
		{"Hello World Again", 3},

		// Multiple spaces
		{"Hello   World", 2},
		{"  Hello  World  ", 2},
		{"Hello    World", 2},

		// Different whitespace characters
		{"Hello\tWorld", 2},
		{"Hello\nWorld", 2},
		{"Hello\rWorld", 2},

		// Cyrillic
		{"Привет Мир", 2},
		{"Здравствуйте", 1},

		// Numbers
		{"123 456 789", 3},
		{"123456789", 1},

		// Punctuation
		{"Hello, World!", 2},
		{"Hello-World", 1},
		{"Hello.World", 1},
	}

	for _, tc := range testCases {
		result := str.WordCount(tc.input)
		if result != tc.expected {
			t.Errorf("WordCount(%q) = %d, want %d", tc.input, result, tc.expected)
		}
	}
}
