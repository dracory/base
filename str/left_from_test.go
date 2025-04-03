package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestLeftFrom_Found tests the case where the needle is found in the string
func TestLeftFrom_Found(t *testing.T) {
	input := "hello world"
	needle := "world"
	expected := "hello "
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestLeftFrom_NotFound tests the case where the needle is not found in the string
func TestLeftFrom_NotFound(t *testing.T) {
	input := "hello world"
	needle := "universe"
	expected := ""
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestLeftFrom_EmptyInput tests the case where the input string is empty
func TestLeftFrom_EmptyInput(t *testing.T) {
	input := ""
	needle := "world"
	expected := ""
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestLeftFrom_EmptyNeedle tests the case where the needle is empty
func TestLeftFrom_EmptyNeedle(t *testing.T) {
	input := "hello world"
	needle := ""
	expected := ""
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestLeftFrom_NeedleAtStart tests the case where the needle is at the start of the string
func TestLeftFrom_NeedleAtStart(t *testing.T) {
	input := "world hello"
	needle := "world"
	expected := ""
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestLeftFrom_NeedleAtEnd tests the case where the needle is at the end of the string
func TestLeftFrom_NeedleAtEnd(t *testing.T) {
	input := "hello world"
	needle := "world"
	expected := "hello "
	
	result := str.LeftFrom(input, needle)
	if result != expected {
		t.Errorf("LeftFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}
