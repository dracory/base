package widgets

import (
	"net/http/httptest"
	"testing"
)

func TestAuthenticatedWidget_Render(t *testing.T) {
	w := NewAuthenticatedWidget()
	if w.Alias() != "x-authenticated" {
		t.Errorf("Alias() = %q", w.Alias())
	}
	// Not authenticated -> content hidden
	if got := w.Render(httptest.NewRequest("GET", "/", nil), "secret", nil); got != "" {
		t.Errorf("expected empty output for unauthenticated request, got %q", got)
	}
}

func TestUnauthenticatedWidget_Render(t *testing.T) {
	w := NewUnauthenticatedWidget()
	if w.Alias() != "x-unauthenticated" {
		t.Errorf("Alias() = %q", w.Alias())
	}
	// Not authenticated -> content shown
	if got := w.Render(httptest.NewRequest("GET", "/", nil), "public", nil); got != "public" {
		t.Errorf("expected content for unauthenticated request, got %q", got)
	}
}

func TestVisibleWidget_EnvironmentMatch(t *testing.T) {
	w := NewVisibleWidget(func(environment string) bool { return environment == "testing" })
	req := httptest.NewRequest("GET", "/", nil)

	if got := w.Render(req, "env-content", map[string]string{"environment": "testing"}); got != "env-content" {
		t.Errorf("expected content for matching environment, got %q", got)
	}
	if got := w.Render(req, "env-content", map[string]string{"environment": "production"}); got != "" {
		t.Errorf("expected empty output for non-matching environment, got %q", got)
	}
	if got := w.Render(req, "env-content", map[string]string{"environment": "bogus"}); got != "" {
		t.Errorf("expected empty output for invalid environment, got %q", got)
	}
}

func TestVisibleWidget_AuthMatch(t *testing.T) {
	w := NewVisibleWidget(nil)
	req := httptest.NewRequest("GET", "/", nil)

	if got := w.Render(req, "auth-content", map[string]string{"auth": "no"}); got != "auth-content" {
		t.Errorf("expected content for auth=no without session, got %q", got)
	}
	if got := w.Render(req, "auth-content", map[string]string{"auth": "yes"}); got != "" {
		t.Errorf("expected empty output for auth=yes without session, got %q", got)
	}
}

func TestVisibleWidget_NilEnvironment(t *testing.T) {
	w := NewVisibleWidget(nil)
	req := httptest.NewRequest("GET", "/", nil)

	if got := w.Render(req, "x", map[string]string{"environment": "testing"}); got != "" {
		t.Errorf("expected empty output with nil environment checker, got %q", got)
	}
}

func TestTermsOfUseWidget_Render(t *testing.T) {
	w := NewTermsOfUseWidget()
	if w.Alias() != "x-terms-of-use" {
		t.Errorf("Alias() = %q", w.Alias())
	}
	got := w.Render(httptest.NewRequest("GET", "/", nil), "", nil)
	if got == "" {
		t.Error("expected terms-of-use HTML output")
	}
}

func TestContactFormWidget_Render(t *testing.T) {
	w := NewContactFormWidget()
	if w.Alias() != "x-contact-form" {
		t.Errorf("Alias() = %q", w.Alias())
	}
	got := w.Render(httptest.NewRequest("GET", "/contact", nil), "", nil)
	if got == "" {
		t.Error("expected contact form link HTML")
	}
}

func TestPrintWidget_Render(t *testing.T) {
	w := NewPrintWidget()
	if w.Alias() != "x-print" {
		t.Errorf("Alias() = %q", w.Alias())
	}
	if got := w.Render(httptest.NewRequest("GET", "/", nil), "1 + 2", nil); got == "" {
		t.Error("expected evaluated output")
	}
}
