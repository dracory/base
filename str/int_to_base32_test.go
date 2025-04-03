package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

func TestIntToBase32_PositiveNumber(t *testing.T) {
	input := 123456
	want := "3oi0"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_Zero(t *testing.T) {
	input := 0
	want := "0"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_NegativeNumber(t *testing.T) {
	input := -123456
	want := "-3oi0"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_MaxInt(t *testing.T) {
	input := 2147483647 // Max int32
	want := "1vvvvvv"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_MinInt(t *testing.T) {
	input := -2147483648 // Min int32
	want := "-2000000"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_PowerOf32(t *testing.T) {
	input := 32
	want := "10"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_SingleDigit(t *testing.T) {
	input := 10
	want := "a"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}

func TestIntToBase32_MultipleDigits(t *testing.T) {
	input := 31
	want := "v"
	got := str.IntToBase32(input)
	if got != want {
		t.Errorf("IntToBase32(%d) = %q, want %q", input, got, want)
	}
}
