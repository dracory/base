package user

import (
	"testing"

	"github.com/dracory/userstore"
)

func TestDisplayNameFull(t *testing.T) {
	// Test with nil user
	name := DisplayNameFull(nil)
	if name != "n/a" {
		t.Errorf("Expected 'n/a', got '%s'", name)
	}

	// Test with user created from userstore (minimal implementation)
	user := userstore.NewUser()
	user.SetFirstName("John")
	user.SetLastName("Doe")
	user.SetEmail("john@example.com")

	name = DisplayNameFull(user)
	expected := "John Doe"
	if name != expected {
		t.Errorf("Expected '%s', got '%s'", expected, name)
	}

	// Test with user with only first name
	user = userstore.NewUser()
	user.SetFirstName("John")
	user.SetEmail("john@example.com")

	name = DisplayNameFull(user)
	expected = "John " // First name + space + empty last name
	if name != expected {
		t.Errorf("Expected '%s', got '%s'", expected, name)
	}

	// Test with user with only last name
	user = userstore.NewUser()
	user.SetLastName("Doe")
	user.SetEmail("john@example.com")

	name = DisplayNameFull(user)
	expected = " Doe" // Empty first name + space + last name
	if name != expected {
		t.Errorf("Expected '%s', got '%s'", expected, name)
	}

	// Test with user with empty names
	user = userstore.NewUser()
	user.SetFirstName("")
	user.SetLastName("")
	user.SetEmail("john@example.com")

	name = DisplayNameFull(user)
	expected = "john@example.com"
	if name != expected {
		t.Errorf("Expected '%s', got '%s'", expected, name)
	}

	// Test with user with whitespace names
	user = userstore.NewUser()
	user.SetFirstName("   ")
	user.SetLastName("   ")
	user.SetEmail("john@example.com")

	name = DisplayNameFull(user)
	expected = "john@example.com"
	if name != expected {
		t.Errorf("Expected '%s', got '%s'", expected, name)
	}
}

func TestIsClient(t *testing.T) {
	// Test with nil user
	isClient := IsClient(nil)
	if isClient {
		t.Error("Expected false for nil user")
	}

	// Test with user marked as client
	user := userstore.NewUser()
	err := user.SetMeta("is_client", "yes")
	if err != nil {
		t.Fatalf("Failed to set meta: %v", err)
	}

	isClient = IsClient(user)
	if !isClient {
		t.Error("Expected true for user with is_client=yes")
	}

	// Test with user not marked as client
	user = userstore.NewUser()
	err = user.SetMeta("is_client", "no")
	if err != nil {
		t.Fatalf("Failed to set meta: %v", err)
	}

	isClient = IsClient(user)
	if isClient {
		t.Error("Expected false for user with is_client=no")
	}

	// Test with user without is_client meta
	user = userstore.NewUser()
	err = user.SetMeta("other_key", "other_value")
	if err != nil {
		t.Fatalf("Failed to set meta: %v", err)
	}

	isClient = IsClient(user)
	if isClient {
		t.Error("Expected false for user without is_client meta")
	}
}

func TestSetIsClient(t *testing.T) {
	// Test with nil user
	result := SetIsClient(nil, true)
	if result != nil {
		t.Error("Expected nil for nil user")
	}

	// Test setting client to true
	user := userstore.NewUser()
	result = SetIsClient(user, true)
	if result == nil {
		t.Error("Expected user to be returned")
	}

	isClient := IsClient(user)
	if !isClient {
		t.Error("Expected user to be marked as client")
	}

	// Test setting client to false
	result = SetIsClient(user, false)
	if result == nil {
		t.Error("Expected user to be returned")
	}

	isClient = IsClient(user)
	if isClient {
		t.Error("Expected user to not be marked as client")
	}

	// Test setting client to true again
	result = SetIsClient(user, true)
	if result == nil {
		t.Error("Expected user to be returned")
	}

	isClient = IsClient(user)
	if !isClient {
		t.Error("Expected user to be marked as client")
	}
}
