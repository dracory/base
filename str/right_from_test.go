package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestRightFrom_Found tests the case where the needle is found in the string
func TestRightFrom_Found(t *testing.T) {
	input := "hello world"
	needle := "hello "
	expected := "world"
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestRightFrom_NotFound tests the case where the needle is not found in the string
func TestRightFrom_NotFound(t *testing.T) {
	input := "hello world"
	needle := "universe"
	expected := ""
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestRightFrom_EmptyInput tests the case where the input string is empty
func TestRightFrom_EmptyInput(t *testing.T) {
	input := ""
	needle := "world"
	expected := ""
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestRightFrom_EmptyNeedle tests the case where the needle is empty
func TestRightFrom_EmptyNeedle(t *testing.T) {
	input := "hello world"
	needle := ""
	expected := input
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestRightFrom_NeedleAtStart tests the case where the needle is at the start of the string
func TestRightFrom_NeedleAtStart(t *testing.T) {
	input := "world hello"
	needle := "world "
	expected := "hello"
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}

// TestRightFrom_NeedleAtEnd tests the case where the needle is at the end of the string
func TestRightFrom_NeedleAtEnd(t *testing.T) {
	input := "hello world"
	needle := "world"
	expected := ""
	
	result := str.RightFrom(input, needle)
	if result != expected {
		t.Errorf("RightFrom(%q, %q) = %q, want %q", input, needle, result, expected)
	}
}
