package str_test

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
	"github.com/dracory/base/str"
)

func TestBcryptHashCompare_ValidMatch(t *testing.T) {
	// Generate a bcrypt hash for testing
	input := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to generate bcrypt hash: %v", err)
	}

	want := true
	got := str.BcryptHashCompare(input, string(hash))
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", input, hash, got, want)
	}
}

func TestBcryptHashCompare_InvalidMatch(t *testing.T) {
	// Generate a bcrypt hash for testing
	input := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to generate bcrypt hash: %v", err)
	}

	// Try with a different password
	want := false
	got := str.BcryptHashCompare("wrongpassword", string(hash))
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", "wrongpassword", hash, got, want)
	}
}

func TestBcryptHashCompare_EmptyString(t *testing.T) {
	// Generate a bcrypt hash for testing
	input := ""
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to generate bcrypt hash: %v", err)
	}

	want := true
	got := str.BcryptHashCompare(input, string(hash))
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", input, hash, got, want)
	}
}

func TestBcryptHashCompare_EmptyHash(t *testing.T) {
	input := "password123"
	want := false
	got := str.BcryptHashCompare(input, "")
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", input, "", got, want)
	}
}

func TestBcryptHashCompare_InvalidHashFormat(t *testing.T) {
	input := "password123"
	invalidHash := "invalidhashformat"
	want := false
	got := str.BcryptHashCompare(input, invalidHash)
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", input, invalidHash, got, want)
	}
}

func TestBcryptHashCompare_DifferentCost(t *testing.T) {
	// Generate a bcrypt hash with different cost
	input := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("Failed to generate bcrypt hash: %v", err)
	}

	want := true
	got := str.BcryptHashCompare(input, string(hash))
	if got != want {
		t.Errorf("BcryptHashCompare(%q, %q) = %v, want %v", input, hash, got, want)
	}
}
