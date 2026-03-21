package htmx

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestIsHtmx(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected bool
	}{
		{
			name:     "no HX-Request header",
			headers:  map[string]string{},
			expected: false,
		},
		{
			name:     "HX-Request header present",
			headers:  map[string]string{"HX-Request": "true"},
			expected: true,
		},
		{
			name:     "HX-Request header with empty value",
			headers:  map[string]string{"HX-Request": ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := IsHtmx(req)
			if result != tt.expected {
				t.Errorf("IsHtmx() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsHxBoosted(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected bool
	}{
		{
			name:     "no HX-Boosted header",
			headers:  map[string]string{},
			expected: false,
		},
		{
			name:     "HX-Boosted header with true",
			headers:  map[string]string{"HX-Boosted": "true"},
			expected: true,
		},
		{
			name:     "HX-Boosted header with false",
			headers:  map[string]string{"HX-Boosted": "false"},
			expected: false,
		},
		{
			name:     "HX-Boosted header with empty value",
			headers:  map[string]string{"HX-Boosted": ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := IsHxBoosted(req)
			if result != tt.expected {
				t.Errorf("IsHxBoosted() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsHxHistoryRestoreRequest(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected bool
	}{
		{
			name:     "no HX-History-Restore-Request header",
			headers:  map[string]string{},
			expected: false,
		},
		{
			name:     "HX-History-Restore-Request header with true",
			headers:  map[string]string{"HX-History-Restore-Request": "true"},
			expected: true,
		},
		{
			name:     "HX-History-Restore-Request header with false",
			headers:  map[string]string{"HX-History-Restore-Request": "false"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := IsHxHistoryRestoreRequest(req)
			if result != tt.expected {
				t.Errorf("IsHxHistoryRestoreRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsHxRequest(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected bool
	}{
		{
			name:     "no HX-Request header",
			headers:  map[string]string{},
			expected: false,
		},
		{
			name:     "HX-Request header with true",
			headers:  map[string]string{"HX-Request": "true"},
			expected: true,
		},
		{
			name:     "HX-Request header with false",
			headers:  map[string]string{"HX-Request": "false"},
			expected: false,
		},
		{
			name:     "HX-Request header with empty value",
			headers:  map[string]string{"HX-Request": ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := IsHxRequest(req)
			if result != tt.expected {
				t.Errorf("IsHxRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsHxTrigger(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected bool
	}{
		{
			name:     "no HX-Trigger header",
			headers:  map[string]string{},
			expected: false,
		},
		{
			name:     "HX-Trigger header with true",
			headers:  map[string]string{"HX-Trigger": "true"},
			expected: true,
		},
		{
			name:     "HX-Trigger header with false",
			headers:  map[string]string{"HX-Trigger": "false"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := IsHxTrigger(req)
			if result != tt.expected {
				t.Errorf("IsHxTrigger() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHxPrompt(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "no HX-Prompt header",
			headers:  map[string]string{},
			expected: "",
		},
		{
			name:     "HX-Prompt header with value",
			headers:  map[string]string{"HX-Prompt": "Enter your name"},
			expected: "Enter your name",
		},
		{
			name:     "HX-Prompt header with empty value",
			headers:  map[string]string{"HX-Prompt": ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := HxPrompt(req)
			if result != tt.expected {
				t.Errorf("HxPrompt() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHxTarget(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "no HX-Target header",
			headers:  map[string]string{},
			expected: "",
		},
		{
			name:     "HX-Target header with value",
			headers:  map[string]string{"HX-Target": "#my-element"},
			expected: "#my-element",
		},
		{
			name:     "HX-Target header with empty value",
			headers:  map[string]string{"HX-Target": ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := HxTarget(req)
			if result != tt.expected {
				t.Errorf("HxTarget() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHxTriggerName(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "no HX-Trigger-Name header",
			headers:  map[string]string{},
			expected: "",
		},
		{
			name:     "HX-Trigger-Name header with value",
			headers:  map[string]string{"HX-Trigger-Name": "click"},
			expected: "click",
		},
		{
			name:     "HX-Trigger-Name header with empty value",
			headers:  map[string]string{"HX-Trigger-Name": ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createRequest(tt.headers)
			result := HxTriggerName(req)
			if result != tt.expected {
				t.Errorf("HxTriggerName() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHxHideIndicatorCSS(t *testing.T) {
	css := HxHideIndicatorCSS()
	if !contains(css, ".htmx-indicator") {
		t.Errorf("CSS does not contain .htmx-indicator")
	}
	if !contains(css, ".htmx-request .htmx-indicator") {
		t.Errorf("CSS does not contain .htmx-request .htmx-indicator")
	}
	if !contains(css, "display: none") {
		t.Errorf("CSS does not contain display: none")
	}
	if !contains(css, "display: inline-block") {
		t.Errorf("CSS does not contain display: inline-block")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Helper function to create an HTTP request with headers
func createRequest(headers map[string]string) *http.Request {
	req := &http.Request{
		Header: make(http.Header),
	}

	// Set a proper URL to avoid nil pointer issues
	req.URL, _ = url.Parse("http://example.com")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}
