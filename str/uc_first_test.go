package str

import (
	"testing"
)

func TestUcFirst_EmptyString(t *testing.T) {
	got := UcFirst("")
	want := ""
	if got != want {
		t.Errorf("UcFirst(\"\") = %q, want %q", got, want)
	}
}

func TestUcFirst_SingleLowercaseLetter(t *testing.T) {
	got := UcFirst("a")
	want := "A"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "a", got, want)
	}
}

func TestUcFirst_SingleUppercaseLetter(t *testing.T) {
	got := UcFirst("A")
	want := "A"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "A", got, want)
	}
}

func TestUcFirst_LowercaseWord(t *testing.T) {
	got := UcFirst("hello")
	want := "Hello"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "hello", got, want)
	}
}

func TestUcFirst_UppercaseWord(t *testing.T) {
	got := UcFirst("WORLD")
	want := "WORLD"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "WORLD", got, want)
	}
}

func TestUcFirst_MixedCaseWord(t *testing.T) {
	got := UcFirst("hElLo")
	want := "HElLo"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "hElLo", got, want)
	}
}

func TestUcFirst_UnicodeCharacter(t *testing.T) {
	got := UcFirst("äpple")
	want := "Äpple"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "äpple", got, want)
	}
}

func TestUcFirst_MultipleWords(t *testing.T) {
	got := UcFirst("hello world")
	want := "Hello world"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "hello world", got, want)
	}
}

func TestUcFirst_CyrillicCharacter(t *testing.T) {
	got := UcFirst("привет")
	want := "Привет"
	if got != want {
		t.Errorf("UcFirst(%q) = %q, want %q", "привет", got, want)
	}
}
