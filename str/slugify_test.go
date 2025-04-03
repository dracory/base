package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

// TestSlugify_Basic tests basic slugification
func TestSlugify_Basic(t *testing.T) {
	input := "Hello World"
	expected := "hello-world"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_SpecialChars tests slugification with special characters
func TestSlugify_SpecialChars(t *testing.T) {
	input := "Hello!@#$%World"
	expected := "hello-world"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_Unicode tests slugification with Unicode characters
func TestSlugify_Unicode(t *testing.T) {
	input := "Héllo Wörld"
	expected := "hello-world"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_Cyrillic tests slugification with Cyrillic characters
func TestSlugify_Cyrillic(t *testing.T) {
	input := "Привет Мир"
	expected := "privet-mir"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_MultipleReplacements tests slugification with multiple replacements
func TestSlugify_MultipleReplacements(t *testing.T) {
	input := "Hello---World---Again"
	expected := "hello-world-again"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_LeadingTrailing tests slugification with leading and trailing replacements
func TestSlugify_LeadingTrailing(t *testing.T) {
	input := "---Hello World---"
	expected := "hello-world"

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_DifferentReplacement tests slugification with different replacement character
func TestSlugify_DifferentReplacement(t *testing.T) {
	input := "Hello World"
	expected := "hello_world"

	result := str.Slugify(input, '_')
	if result != expected {
		t.Errorf("Slugify(%q, '_') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_EmptyString tests slugification with empty string
func TestSlugify_EmptyString(t *testing.T) {
	input := ""
	expected := ""

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestSlugify_OnlySpecialChars tests slugification with only special characters
func TestSlugify_OnlySpecialChars(t *testing.T) {
	input := "!@#$%"
	expected := ""

	result := str.Slugify(input, '-')
	if result != expected {
		t.Errorf("Slugify(%q, '-') = %q, want %q", input, result, expected)
	}
}

// TestCyrillicToLatin tests the CyrillicToLatin function
func Test_Slugify_CyrillicToLatin(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Привет", "privet"},
		{"Мир", "mir"},
		{"АБВГД", "abvgd"},
		{"ЕЖЗИЙ", "ezhziy"},
		{"КЛМНО", "klmno"},
		{"ПРСТУ", "prstu"},
		{"ФХЦЧШ", "fhtschsh"},
		{"ЩЪЬЮЯ", "shtayyuya"},
		{"абвгд", "abvgd"},
		{"ежзий", "ezhziy"},
		{"клмно", "klmno"},
		{"прсту", "prstu"},
		{"фхцчш", "fhtschsh"},
		{"щъьюя", "shtayyuya"},
	}

	for _, tc := range testCases {
		result := str.Slugify(tc.input, '-')
		if result != tc.expected {
			t.Errorf("Slugify(%q, '-') = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
