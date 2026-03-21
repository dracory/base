package http

import "net/http"

// Redirect performs an HTTP redirect to the specified URL using temporary redirect (307).
// It wraps the standard http.Redirect function for consistency across the codebase.
//
// Parameters:
//   - w: http.ResponseWriter to write the redirect response
//   - r: *http.Request representing the current request
//   - url: string target URL to redirect to
//
// Returns:
//   - string: empty string (for compatibility with controller return patterns)
func Redirect(w http.ResponseWriter, r *http.Request, url string) string {
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return ""
}
