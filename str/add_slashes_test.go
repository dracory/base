package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestAddSlashes(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No special chars", "hello", "hello"},
		{"Single quote", "hello'world", "hello\\'world"},
		{"Double quote", "hello\"world", "hello\\\"world"},
		{"Backslash", "hello\\world", "hello\\\\world"},
		{"All special chars", "hello'\"\\world", "hello\\'\\\"\\\\world"},
		{"Special chars only", "'\"\\", "\\'\\\"\\\\"},
		{"Special chars at start", "'hello", "\\'hello"},
		{"Special chars at end", "hello\\", "hello\\\\"},
		{"Multiple consecutive special chars", "hello'''world", "hello\\'\\'\\'world"},
		{"Multiple consecutive special chars at start", "'''hello", "\\'\\'\\'hello"},
		{"Multiple consecutive special chars at end", "hello\\\\\\", "hello\\\\\\\\\\\\"},
		{"Mixed special chars", "hello'\"\\world'\"", "hello\\'\\\"\\\\world\\'\\\""},
		{"Long string", "This is a very long string with 'quotes' and \"double quotes\" and \\backslashes\\.", "This is a very long string with \\'quotes\\' and \\\"double quotes\\\" and \\\\backslashes\\\\."},
		{"Unicode string", "привет'мир", "привет\\'мир"},
		{"Unicode and special chars", "привет'\"\\мир", "привет\\'\\\"\\\\мир"},
		{"Control character", "hello\nworld", "hello\nworld"},
		{"Multiple control characters", "hello\n\r\tworld", "hello\n\r\tworld"},
		{"Nil input", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.AddSlashes(tc.input)
			if result != tc.expected {
				t.Errorf("AddSlashes(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAddSlashesCustom(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		escapeChars string
		expected    string
	}{
		{"Empty string, empty escapeChars", "", "", ""},
		{"Empty string, with escapeChars", "", "'\"", ""},
		{"No special chars, empty escapeChars", "hello", "", "hello"},
		{"No special chars, with escapeChars", "hello", "'\"", "hello"},
		{"Single char to escape", "hello'world", "'", "hello\\'world"},
		{"Multiple chars to escape", "hello'\"world", "'\"", "hello\\'\\\"world"},
		{"Escape char at start", "'hello", "'", "\\'hello"},
		{"Escape char at end", "hello'", "'", "hello\\'"},
		{"Multiple consecutive escape chars", "hello'''world", "'", "hello\\'\\'\\'world"},
		{"Mixed chars to escape and not", "hello'world\"!", "'\"", "hello\\'world\\\"!"},
		{"Unicode string", "привет'мир", "'", "привет\\'мир"},
		{"Unicode and escape chars", "привет'\"\\мир", "'\"\\", "привет\\'\\\"\\\\мир"},
		{"Control character", "hello\nworld", "\n", "hello\\\nworld"},
		{"Multiple control characters", "hello\n\r\tworld", "\n\r\t", "hello\\\n\\\r\\\tworld"},
		{"Long string", "This is a very long string with 'quotes' and \"double quotes\" and \\backslashes\\.", "'\"\\", "This is a very long string with \\'quotes\\' and \\\"double quotes\\\" and \\\\backslashes\\\\."},
		{"Escape only backslash", "hello\\world", "\\", "hello\\\\world"},
		{"Escape only double quote", "hello\"world", "\"", "hello\\\"world"},
		{"Escape only single quote", "hello'world", "'", "hello\\'world"},
		{"Escape multiple chars", "hello'\"\\world", "'\"\\", "hello\\'\\\"\\\\world"},
		{"Escape multiple chars at start", "'\"\\hello", "'\"\\", "\\'\\\"\\\\hello"},
		{"Escape multiple chars at end", "hello'\"\\", "'\"\\", "hello\\'\\\"\\\\"},
		{"Escape multiple chars consecutive", "hello'''\"\"\"\\\\\\world", "'\"\\", "hello\\'\\'\\'\\\"\\\"\\\"\\\\\\\\\\\\world"},
		{"Escape multiple chars mixed", "hello'\"\\world'\"\\", "'\"\\", "hello\\'\\\"\\\\world\\'\\\"\\\\"},
		{"Escape multiple chars mixed with unicode", "привет'\"\\мир'\"\\", "'\"\\", "привет\\'\\\"\\\\мир\\'\\\"\\\\"},
		{"Escape multiple chars mixed with control chars", "hello\n\r\t'\"\\world\n\r\t'\"\\", "'\"\\\n\r\t", "hello\\\n\\\r\\\t\\'\\\"\\\\world\\\n\\\r\\\t\\'\\\"\\\\"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.AddSlashesCustom(tc.input, tc.escapeChars)
			if result != tc.expected {
				t.Errorf("AddSlashesCustom(%q, %q) = %q, want %q", tc.input, tc.escapeChars, result, tc.expected)
			}
		})
	}
}
