package widgets

import "net/http"

// Widget is the shortcode contract used by CMS-rendered content.
// It is structurally identical to cmsstore.ShortcodeInterface, so widgets
// implementing it can be registered with a CMS store directly.
type Widget interface {
	// Alias returns a unique identifier for the shortcode.
	Alias() string

	// Description provides a brief explanation of the shortcode's purpose.
	Description() string

	// Render generates the final output of the shortcode.
	Render(r *http.Request, content string, params map[string]string) string
}
