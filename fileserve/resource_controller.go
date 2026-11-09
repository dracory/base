package fileserve

import (
	"net/http"
	"strings"

	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

// resourceController serves embedded/static resources (CSS, JS, images,
// fonts) from an injected reader. Paths starting with "." are treated as
// private and answered with the configured not-found handler.
type resourceController struct {
	readResource func(uri string) (string, error)
	notFound     func(w http.ResponseWriter, r *http.Request) string
}

// == CONSTRUCTOR =============================================================

// NewResourceController creates a controller that serves resources through
// the given reader.
//
// Parameters:
//   - readResource: resolves a resource URI to its content (required)
//   - notFound: fallback handler for private or missing resources. If nil,
//     a plain "404 page not found" message is returned.
func NewResourceController(
	readResource func(uri string) (string, error),
	notFound func(w http.ResponseWriter, r *http.Request) string,
) *resourceController {
	if notFound == nil {
		notFound = func(w http.ResponseWriter, r *http.Request) string {
			return "404 page not found"
		}
	}
	return &resourceController{readResource: readResource, notFound: notFound}
}

// PUBLIC METHODS =============================================================

// Handler serves the resource identified by the request URI. The leading
// "/resources/" prefix is stripped before the lookup.
//
// Returns the resource content on success, or the not-found handler output.
func (controller resourceController) Handler(w http.ResponseWriter, r *http.Request) string {
	uri := r.RequestURI

	// Remove leading /resources path
	if strings.HasPrefix(uri, "/resources/") {
		uri = strings.Replace(uri, "/resources/", "", 1)
	}

	// Is resource private?
	if strings.HasPrefix(uri, ".") {
		w.WriteHeader(http.StatusNotFound)
		return controller.notFound(w, r)
	}

	contentType := lo.If(strings.HasSuffix(uri, ".css"), "text/css").
		ElseIf(strings.HasSuffix(uri, ".gif"), "image/gif").
		ElseIf(strings.HasSuffix(uri, ".html"), "text/html").
		ElseIf(strings.HasSuffix(uri, ".ico"), "image/x-icon").
		ElseIf(strings.HasSuffix(uri, ".jpg"), "image/jpeg").
		ElseIf(strings.HasSuffix(uri, ".jpeg"), "image/jpeg").
		ElseIf(strings.HasSuffix(uri, ".js"), "application/javascript").
		ElseIf(strings.HasSuffix(uri, ".png"), "image/png").
		ElseIf(strings.HasSuffix(uri, ".ttf"), "font/ttf").
		ElseIf(strings.HasSuffix(uri, ".svg"), "image/svg+xml").
		ElseIf(strings.HasSuffix(uri, ".webp"), "image/webp").
		ElseIf(strings.HasSuffix(uri, ".woff"), "font/woff").
		ElseIf(strings.HasSuffix(uri, ".woff2"), "font/woff2").
		Else("text/plain")

	resourceContent, err := controller.readResource(uri)

	if err != nil {
		return controller.notFound(w, r)
	}

	if resourceContent == "" {
		return controller.notFound(w, r)
	}

	w.Header().Set("Content-Type", contentType)
	return resourceContent
}
