package fileserve

import (
	"net/http"
	"strings"

	"github.com/dracory/str"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// StorageInterface is the minimal file storage contract required by the
// file-serving controller. It is satisfied by filesystem.StorageInterface
// and compatible implementations.
type StorageInterface interface {
	// Exists checks if a file exists at the given path
	Exists(filePath string) (bool, error)

	// ReadFile reads the file at the given path
	ReadFile(filePath string) ([]byte, error)
}

// == CONTROLLER ==============================================================

// fileServeController serves files from a storage backend with correct
// content-type detection. It can be mounted at any path prefix (e.g.
// "/files/*" or "/media/*") — the prefix is stripped from the request path.
type fileServeController struct {
	storage StorageInterface
}

// == CONSTRUCTOR =============================================================

// NewFileServeController creates a controller that serves files from the
// given storage. A nil storage produces a "File storage not configured"
// response instead of panicking.
func NewFileServeController(storage StorageInterface) *fileServeController {
	return &fileServeController{storage: storage}
}

// == PUBLIC METHODS ==========================================================

// Handler serves the file identified by the request path from the storage.
// The leading path prefix ("/files", "/file", or "/media") is stripped before
// the storage lookup.
//
// Returns an error message string on failure, or an empty string on success.
func (c *fileServeController) Handler(w http.ResponseWriter, r *http.Request) string {
	if c.storage == nil {
		return "File storage not configured"
	}

	filePath := lo.IfF(strings.HasPrefix(r.URL.Path, "/files"), func() string { return str.RightFrom(r.URL.Path, "/files") }).
		ElseIfF(strings.HasPrefix(r.URL.Path, "/file"), func() string { return str.RightFrom(r.URL.Path, "/file") }).
		ElseIfF(strings.HasPrefix(r.URL.Path, "/media"), func() string { return str.RightFrom(r.URL.Path, "/media") }).
		Else(r.URL.Path)

	exists, err := c.storage.Exists(filePath)

	if err != nil {
		return err.Error()
	}

	if !exists {
		return "File not found"
	}

	content, err := c.storage.ReadFile(filePath)

	if err != nil {
		return err.Error()
	}

	extension := c.findExtension(filePath)
	mimeType := c.findMIMEType(extension)

	if extension == "" {
		return "File not found"
	}

	w.Header().Set("Content-Type", mimeType)

	if mimeType == "application/octet-stream" {
		w.Header().Set("Content-Disposition", "attachment; filename="+r.URL.Path)
		w.Header().Set("Content-Length", cast.ToString(len(content)))
	}

	if _, err := w.Write(content); err != nil { // #nosec G705 -- content is file data from validated path
		return "Failed to write file content: " + err.Error()
	}

	return ""
}

// findExtension finds the file extension from a path.
//
// Parameter(s):
//   - path string - the path
//
// Return type(s):
//   - string - the file extension
func (controller fileServeController) findExtension(path string) string {
	fileName := controller.findFileName(path)

	if fileName == "" {
		return ""
	}

	nameParts := strings.Split(fileName, ".")

	if len(nameParts) < 2 {
		return ""
	}

	return nameParts[1]
}

// findMIMEType maps a file extension to its MIME type.
func (controller fileServeController) findMIMEType(extension string) string {
	switch extension {
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	case "js":
		return "application/javascript"
	case "json":
		return "application/json"
	case "png":
		return "image/png"
	case "jpg", "jpeg":
		return "image/jpeg"
	case "gif":
		return "image/gif"
	case "svg":
		return "image/svg+xml"
	case "ico":
		return "image/x-icon"
	case "pdf":
		return "application/pdf"
	case "zip":
		return "application/zip"
	case "mp3":
		return "audio/mpeg"
	case "webm":
		return "video/webm"
	default:
		return "application/octet-stream"
	}
}

// findFileName finds the file name from a path.
//
// Parameter(s):
//   - path string - the path
//
// Return type(s):
//   - string - the file name
func (controller fileServeController) findFileName(path string) string {
	uriParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(uriParts) < 1 {
		return ""
	}

	return uriParts[len(uriParts)-1]
}
