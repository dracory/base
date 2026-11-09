package widgets

import (
	"net/http"
	"slices"

	basesession "github.com/dracory/base/session"

	"github.com/samber/lo"
)

var _ Widget = (*visibleWidget)(nil) // verify it extends the interface

// AllowedEnvironments is the list of environment names accepted by the
// "environment" parameter of the visible widget.
var AllowedEnvironments = []string{
	"development",
	"local",
	"production",
	"staging",
	"testing",
}

// == CONSTUCTOR ==============================================================

// NewVisibleWidget creates a new instance of the visible widget.
//
// Parameters:
//   - isEnvironment: reports whether the given environment name is the
//     currently active one (e.g. func(env string) bool { return appEnv == env }).
//     If nil, the "environment" parameter never matches.
//
// Returns:
//   - *visibleWidget - A pointer to the visible widget
func NewVisibleWidget(isEnvironment func(environment string) bool) *visibleWidget {
	return &visibleWidget{
		isEnvironment: isEnvironment,
	}
}

// == WIDGET ================================================================

// visibleWidget is the struct that will be used to render the visible shortcode.
//
// This shortcode is used to show the result of the provided content
// if a condition is met.
//
// Examples:
// <x-visible environment="production">content</x-visible>
// <x-visible environment="development" auth="yes">content</x-visible>
// <x-visible environment="staging" auth="no">content</x-visible>
// <x-visible environment="local" auth="yes">content</x-visible>
type visibleWidget struct {
	isEnvironment func(environment string) bool
}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (w *visibleWidget) Alias() string {
	return "x-visible"
}

// Description a user-friendly description of the shortcode.
func (w *visibleWidget) Description() string {
	return "Renders the content if the condition is met"
}

// Render implements the shortcode interface.
func (w *visibleWidget) Render(r *http.Request, content string, params map[string]string) string {
	environment := lo.ValueOr(params, "environment", "")
	auth := lo.ValueOr(params, "auth", "")
	showContent := []bool{}

	if w.isEnvAllowedValue(environment) {
		showContent = append(showContent, w.isEnvironmentMatch(environment))
	}

	if auth != "" {
		showContent = append(showContent, w.isAuthMatch(r, auth))
	}

	if w.allTrue(showContent) {
		return content
	}

	return "" // No content is shown by default
}

// == PRIVATE METHODS ========================================================

// allTrue returns true if the provided array is not empty and all values
// in the array are true
func (w *visibleWidget) allTrue(arr []bool) bool {
	if len(arr) == 0 {
		return false
	}

	for _, val := range arr {
		if !val {
			return false
		}
	}
	return true
}

// isAuthAllowedValue returns true if the provided value is a valid
// auth value: yes, no
func (w *visibleWidget) isAuthAllowedValue(auth string) bool {
	if auth == "" {
		return false
	}

	return slices.Contains([]string{"yes", "no"}, auth)
}

// isAuthMatch returns true if the provided value matches the authentication
// status of the user
func (w *visibleWidget) isAuthMatch(req *http.Request, authenticated string) bool {
	if authenticated == "" {
		return false
	}

	if !w.isAuthAllowedValue(authenticated) {
		return false
	}

	authUser := basesession.GetAuthUser(req)

	isAuth := lo.Ternary(authUser != nil, true, false)

	if authenticated == "yes" && isAuth {
		return true
	}

	if authenticated == "no" && !isAuth {
		return true
	}

	return false
}

// isEnvAllowedValue returns true if the provided value is a valid environment
// value: development, local, production, staging, testing
func (w *visibleWidget) isEnvAllowedValue(environment string) bool {
	if environment == "" {
		return false
	}

	return slices.Contains(AllowedEnvironments, environment)
}

// isEnvironmentMatch returns true if the provided environment name matches
// the currently active environment via the injected checker.
func (w *visibleWidget) isEnvironmentMatch(environment string) bool {
	if environment == "" || w.isEnvironment == nil {
		return false
	}

	return w.isEnvironment(environment)
}
