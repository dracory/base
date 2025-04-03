package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

func TestAfterLast_EmptyString(t *testing.T) {
	got := str.AfterLast("", ",")
	want := ""
	if got != want {
		t.Errorf("AfterLast(\"\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_NoMatch(t *testing.T) {
	got := str.AfterLast("hello", ",")
	want := "hello"
	if got != want {
		t.Errorf("AfterLast(\"hello\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_SingleMatch(t *testing.T) {
	got := str.AfterLast("hello,world", ",")
	want := "world"
	if got != want {
		t.Errorf("AfterLast(\"hello,world\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_MultipleMatches(t *testing.T) {
	got := str.AfterLast("hello,world,again", ",")
	want := "again"
	if got != want {
		t.Errorf("AfterLast(\"hello,world,again\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_WithSpaces(t *testing.T) {
	got := str.AfterLast("hello world", " ")
	want := "world"
	if got != want {
		t.Errorf("AfterLast(\"hello world\", \" \") = %q, want %q", got, want)
	}
}

func TestAfterLast_WithUnicode(t *testing.T) {
	got := str.AfterLast("привет,мир", ",")
	want := "мир"
	if got != want {
		t.Errorf("AfterLast(\"привет,мир\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_NeedleAtEnd(t *testing.T) {
	got := str.AfterLast("hello,", ",")
	want := ""
	if got != want {
		t.Errorf("AfterLast(\"hello,\", \",\") = %q, want %q", got, want)
	}
}

func TestAfterLast_NeedleAtStart(t *testing.T) {
	got := str.AfterLast(",hello", ",")
	want := "hello"
	if got != want {
		t.Errorf("AfterLast(\",hello\", \",\") = %q, want %q", got, want)
	}
}
