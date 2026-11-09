package fileserve

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func stubNotFound(w http.ResponseWriter, r *http.Request) string {
	w.WriteHeader(http.StatusNotFound)
	return "404 page not found"
}

func stubReader(uri string) (string, error) {
	if uri == "js/app.js" {
		return "console.log('hi');", nil
	}
	return "", errors.New("resource not found")
}

// TestResourceHandlerServesResource ensures an existing resource is served
// with the correct Content-Type.
func TestResourceHandlerServesResource(t *testing.T) {
	c := newResourceController(stubReader, stubNotFound)

	req := httptest.NewRequest(http.MethodGet, "/resources/js/app.js", nil)
	rec := httptest.NewRecorder()

	res := c.Handler(rec, req)

	if res != "console.log('hi');" {
		t.Fatalf("expected resource content, got: %q", res)
	}

	ct := rec.Header().Get("Content-Type")
	if ct != "application/javascript" {
		t.Fatalf("expected Content-Type 'application/javascript', got: %q", ct)
	}
}

// TestResourceHandlerPrivateResource ensures paths starting with '.' return
// 404 and the not-found handler output.
func TestResourceHandlerPrivateResource(t *testing.T) {
	c := newResourceController(stubReader, stubNotFound)

	req := httptest.NewRequest(http.MethodGet, "/resources/.secrets.txt", nil)
	rec := httptest.NewRecorder()

	res := c.Handler(rec, req)

	if res != "404 page not found" {
		t.Fatalf("expected not-found message, got: %q", res)
	}
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

// TestResourceHandlerMissingResource ensures missing resources go through
// the not-found handler.
func TestResourceHandlerMissingResource(t *testing.T) {
	c := newResourceController(stubReader, stubNotFound)

	req := httptest.NewRequest(http.MethodGet, "/resources/js/does_not_exist.js", nil)
	rec := httptest.NewRecorder()

	res := c.Handler(rec, req)

	if res != "404 page not found" {
		t.Fatalf("expected not-found message, got: %q", res)
	}
}

// TestResourceHandlerDefaultNotFound ensures a nil not-found handler falls
// back to a default message.
func TestResourceHandlerDefaultNotFound(t *testing.T) {
	c := newResourceController(stubReader, nil)

	req := httptest.NewRequest(http.MethodGet, "/resources/js/does_not_exist.js", nil)
	rec := httptest.NewRecorder()

	res := c.Handler(rec, req)

	if res != "404 page not found" {
		t.Fatalf("expected default not-found message, got: %q", res)
	}
}

// TestResourceHandlerReaderError ensures reader errors go through the
// not-found handler.
func TestResourceHandlerReaderError(t *testing.T) {
	c := newResourceController(func(uri string) (string, error) {
		return "", errors.New("boom")
	}, stubNotFound)

	req := httptest.NewRequest(http.MethodGet, "/resources/js/app.js", nil)
	rec := httptest.NewRecorder()

	if res := c.Handler(rec, req); res != "404 page not found" {
		t.Fatalf("expected not-found message, got: %q", res)
	}
}
