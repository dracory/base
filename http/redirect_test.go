package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRedirect(t *testing.T) {
	tests := []struct {
		name           string
		targetURL      string
		expectedStatus int
		expectedLoc    string
	}{
		{
			name:           "Basic redirect",
			targetURL:      "https://example.com",
			expectedStatus: http.StatusTemporaryRedirect,
			expectedLoc:    "https://example.com",
		},
		{
			name:           "Relative URL redirect",
			targetURL:      "/login",
			expectedStatus: http.StatusTemporaryRedirect,
			expectedLoc:    "/login",
		},
		{
			name:           "URL with query parameters",
			targetURL:      "/search?q=test&page=1",
			expectedStatus: http.StatusTemporaryRedirect,
			expectedLoc:    "/search?q=test&page=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test request
			req := httptest.NewRequest("GET", "/current", nil)
			w := httptest.NewRecorder()

			// Call the redirect function
			result := Redirect(w, req, tt.targetURL)

			// Check response status
			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			// Check Location header
			location := resp.Header.Get("Location")
			if location != tt.expectedLoc {
				t.Errorf("Expected location %q, got %q", tt.expectedLoc, location)
			}

			// Check return value (should be empty string for controller compatibility)
			if result != "" {
				t.Errorf("Expected empty string return, got %q", result)
			}
		})
	}
}

func TestRedirectWithFullURL(t *testing.T) {
	// Test with a full URL including scheme and host
	targetURL := "https://example.com/path?query=value#fragment"

	req := httptest.NewRequest("GET", "/current", nil)
	w := httptest.NewRecorder()

	Redirect(w, req, targetURL)

	resp := w.Result()

	// Verify status code
	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("Expected status %d, got %d", http.StatusTemporaryRedirect, resp.StatusCode)
	}

	// Verify location header
	location := resp.Header.Get("Location")
	if location != targetURL {
		t.Errorf("Expected location %q, got %q", targetURL, location)
	}

	// Verify the URL is properly formatted
	parsedURL, err := url.Parse(location)
	if err != nil {
		t.Errorf("Failed to parse redirect URL: %v", err)
	}

	if parsedURL.Scheme != "https" {
		t.Errorf("Expected scheme 'https', got %q", parsedURL.Scheme)
	}

	if parsedURL.Host != "example.com" {
		t.Errorf("Expected host 'example.com', got %q", parsedURL.Host)
	}
}

func TestRedirectIntegration(t *testing.T) {
	// Test that the redirect works like the standard http.Redirect
	req := httptest.NewRequest("POST", "/current", nil)
	req.Form = url.Values{}
	req.Form.Add("foo", "bar")

	w := httptest.NewRecorder()

	Redirect(w, req, "/target")

	resp := w.Result()

	// Should be temporary redirect (307) to preserve method
	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("Expected status %d, got %d", http.StatusTemporaryRedirect, resp.StatusCode)
	}

	// Location should be set correctly
	if location := resp.Header.Get("Location"); location != "/target" {
		t.Errorf("Expected location '/target', got %q", location)
	}
}
