package tz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	basesession "github.com/dracory/base/session"

	"github.com/dracory/userstore"
)

// Test FromUser with nil
func TestFromUser_NilUser(t *testing.T) {
	result := FromUser(nil)
	if result != "UTC" {
		t.Errorf("FromUser(nil) = %v, want UTC", result)
	}
}

// Test FromRequest with nil request
func TestFromRequest_NilRequest(t *testing.T) {
	result := FromRequest(nil)
	if result != "UTC" {
		t.Errorf("FromRequest(nil) = %v, want UTC", result)
	}
}

// Test FromRequest without user in context
func TestFromRequest_NoUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	result := FromRequest(req)
	if result != "UTC" {
		t.Errorf("FromRequest(req) without user = %v, want UTC", result)
	}
}

// TestFromUser_WithUserTimezone tests retrieving timezone from user
func TestFromUser_WithUserTimezone(t *testing.T) {
	user := userstore.NewUser().
		SetEmail("timezone@example.com").
		SetFirstName("TimeZone").
		SetLastName("User").
		SetTimezone("America/New_York")

	result := FromUser(user)
	if result != "America/New_York" {
		t.Errorf("FromUser = %v, want America/New_York", result)
	}
}

// TestFromRequest_WithUserTimezone tests retrieving timezone from request
func TestFromRequest_WithUserTimezone(t *testing.T) {
	user := userstore.NewUser().
		SetEmail("timezone@example.com").
		SetFirstName("TimeZone").
		SetLastName("User").
		SetTimezone("America/New_York")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = basesession.SetAuthUser(req, user)

	result := FromRequest(req)
	if result != "America/New_York" {
		t.Errorf("FromRequest = %v, want America/New_York", result)
	}
}

// TestFromUser_UserWithEmptyTimezone tests that empty timezone returns UTC
func TestFromUser_UserWithEmptyTimezone(t *testing.T) {
	user := userstore.NewUser().
		SetEmail("notimezone@example.com").
		SetFirstName("NoTimeZone").
		SetLastName("User")

	result := FromUser(user)
	if result != "UTC" {
		t.Errorf("FromUser = %v, want UTC", result)
	}
}

// TestFromRequest_UserWithEmptyTimezone tests that empty timezone returns UTC
func TestFromRequest_UserWithEmptyTimezone(t *testing.T) {
	user := userstore.NewUser().
		SetEmail("notimezone@example.com").
		SetFirstName("NoTimeZone").
		SetLastName("User")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = basesession.SetAuthUser(req, user)

	result := FromRequest(req)
	if result != "UTC" {
		t.Errorf("FromRequest = %v, want UTC", result)
	}
}
