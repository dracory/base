package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

func TestAfter_EmptyNeedle(t *testing.T) {
	input := "hello world"
	needle := ""
	want := input
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_NeedleNotFound(t *testing.T) {
	input := "hello world"
	needle := "xyz"
	want := input
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_NeedleAtStart(t *testing.T) {
	input := "hello world"
	needle := "hello"
	want := " world"
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_NeedleInMiddle(t *testing.T) {
	input := "hello world"
	needle := "o"
	want := " world"
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_NeedleAtEnd(t *testing.T) {
	input := "hello world"
	needle := "world"
	want := ""
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_EmptyInput(t *testing.T) {
	input := ""
	needle := "x"
	want := ""
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}

func TestAfter_EmptyInputAndNeedle(t *testing.T) {
	input := ""
	needle := ""
	want := ""
	got := str.After(input, needle)
	if got != want {
		t.Errorf("After(%q, %q) = %q, want %q", input, needle, got, want)
	}
}
