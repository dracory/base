package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIntToBase36_PositiveNumber(t *testing.T) {
	input := 123456
	want := "2n9c"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_Zero(t *testing.T) {
	input := 0
	want := "0"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_NegativeNumber(t *testing.T) {
	input := -123456
	want := "-2n9c"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_MaxInt(t *testing.T) {
	input := 2147483647 // Max int32
	want := "zik0zj"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_MinInt(t *testing.T) {
	input := -2147483648 // Min int32
	want := "-zik0zk"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_PowerOf36(t *testing.T) {
	input := 36
	want := "10"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_SingleDigit(t *testing.T) {
	input := 10
	want := "a"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase36_MultipleDigits(t *testing.T) {
	input := 35
	want := "z"
	got := str.IntToBase36(input)
	if got != want {
		t.Errorf("IntToBase36(%d) = %q, want %q", input, got, want)
	}
}
