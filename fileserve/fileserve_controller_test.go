package fileserve

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandlerWithNilStorage verifies that Handler returns the proper error message
// when the controller is constructed without a storage implementation.
func TestHandlerWithNilStorage(t *testing.T) {
	c := &fileServeController{storage: nil}

	req := httptest.NewRequest(http.MethodGet, "/media/example.txt", nil)
	rec := httptest.NewRecorder()

	res := c.Handler(rec, req)

	if res != "File storage not configured" {
		t.Fatalf("expected 'File storage not configured', got: %q", res)
	}

	// Ensure that nothing was written to the response body in this early-return case
	if rec.Body.Len() != 0 {
		t.Fatalf("expected empty response body when storage not configured, got: %q", rec.Body.String())
	}
}

// TestFindFileNameRoot ensures last segment extraction works from root-level path
func TestFindFileNameRoot(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/hello.txt")
	if got != "hello.txt" {
		t.Fatalf("expected 'hello.txt', got: %q", got)
	}
}

// TestFindFileNameNested ensures last segment extraction works from nested paths
func TestFindFileNameNested(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/a/b/c/image.png")
	if got != "image.png" {
		t.Fatalf("expected 'image.png', got: %q", got)
	}
}

// TestFindFileNameEmpty ensures empty or slash-only path returns empty string
func TestFindFileNameEmpty(t *testing.T) {
	c := fileServeController{}
	if got := c.findFileName(""); got != "" {
		t.Fatalf("expected '', got: %q", got)
	}
	if got := c.findFileName("/"); got != "" {
		t.Fatalf("expected '', got: %q", got)
	}
}

// TestFindExtensionHappyPath ensures extension is parsed after last dot
func TestFindExtensionHappyPath(t *testing.T) {
	c := fileServeController{}
	if got := c.findExtension("/dir/file.html"); got != "html" {
		t.Fatalf("expected 'html', got: %q", got)
	}
	if got := c.findExtension("/dir/file.css"); got != "css" {
		t.Fatalf("expected 'css', got: %q", got)
	}
	if got := c.findExtension("/dir/file.jpg"); got != "jpg" {
		t.Fatalf("expected 'jpg', got: %q", got)
	}
}

// TestFindExtensionNoName ensures paths without filename produce empty extension
func TestFindExtensionNoName(t *testing.T) {
	c := fileServeController{}
	if got := c.findExtension("/"); got != "" {
		t.Fatalf("expected '', got: %q", got)
	}
	if got := c.findExtension(""); got != "" {
		t.Fatalf("expected '', got: %q", got)
	}
}

// TestFindExtensionNoDot ensures filenames without dot produce empty extension
func TestFindExtensionNoDot(t *testing.T) {
	c := fileServeController{}
	if got := c.findExtension("/dir/readme"); got != "" {
		t.Fatalf("expected '', got: %q", got)
	}
}

// TestFindMIMETypeAllFormats ensures all supported MIME types are correctly detected
func TestFindMIMETypeAllFormats_Html(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("html")
	if got != "text/html" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "html", got, "text/html")
	}
}

func TestFindMIMETypeAllFormats_Css(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("css")
	if got != "text/css" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "css", got, "text/css")
	}
}

func TestFindMIMETypeAllFormats_Js(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("js")
	if got != "application/javascript" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "js", got, "application/javascript")
	}
}

func TestFindMIMETypeAllFormats_Json(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("json")
	if got != "application/json" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "json", got, "application/json")
	}
}

func TestFindMIMETypeAllFormats_Png(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("png")
	if got != "image/png" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "png", got, "image/png")
	}
}

func TestFindMIMETypeAllFormats_Jpg(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("jpg")
	if got != "image/jpeg" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "jpg", got, "image/jpeg")
	}
}

func TestFindMIMETypeAllFormats_Jpeg(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("jpeg")
	if got != "image/jpeg" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "jpeg", got, "image/jpeg")
	}
}

func TestFindMIMETypeAllFormats_Gif(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("gif")
	if got != "image/gif" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "gif", got, "image/gif")
	}
}

func TestFindMIMETypeAllFormats_Svg(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("svg")
	if got != "image/svg+xml" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "svg", got, "image/svg+xml")
	}
}

func TestFindMIMETypeAllFormats_Ico(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("ico")
	if got != "image/x-icon" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "ico", got, "image/x-icon")
	}
}

func TestFindMIMETypeAllFormats_Pdf(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("pdf")
	if got != "application/pdf" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "pdf", got, "application/pdf")
	}
}

func TestFindMIMETypeAllFormats_Zip(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("zip")
	if got != "application/zip" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "zip", got, "application/zip")
	}
}

func TestFindMIMETypeAllFormats_Mp3(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("mp3")
	if got != "audio/mpeg" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "mp3", got, "audio/mpeg")
	}
}

func TestFindMIMETypeAllFormats_Webm(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("webm")
	if got != "video/webm" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "webm", got, "video/webm")
	}
}

func TestFindMIMETypeAllFormats_Unknown(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("unknown")
	if got != "application/octet-stream" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "unknown", got, "application/octet-stream")
	}
}

func TestFindMIMETypeAllFormats_Empty(t *testing.T) {
	c := fileServeController{}
	got := c.findMIMEType("")
	if got != "application/octet-stream" {
		t.Errorf("findMIMEType(%q) = %q, want %q", "", got, "application/octet-stream")
	}
}

// TestNewFileServeController verifies the constructor properly initializes the controller
func TestNewFileServeController(t *testing.T) {
	// Test with nil storage
	c := NewFileServeController(nil)
	if c == nil {
		t.Fatal("NewFileServeController(nil) should not return nil")
	}
	if c.storage != nil {
		t.Error("expected storage to be nil")
	}
}

// TestFindFileNameWithSpecialCharacters tests edge cases in filename extraction
func TestFindFileNameWithSpecialCharacters_FileWithDashes(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/path/to/file-with-dashes.txt")
	if got != "file-with-dashes.txt" {
		t.Errorf("findFileName(%q) = %q, want %q", "/path/to/file-with-dashes.txt", got, "file-with-dashes.txt")
	}
}

func TestFindFileNameWithSpecialCharacters_FileWithUnderscores(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/path/to/file_with_underscores.txt")
	if got != "file_with_underscores.txt" {
		t.Errorf("findFileName(%q) = %q, want %q", "/path/to/file_with_underscores.txt", got, "file_with_underscores.txt")
	}
}

func TestFindFileNameWithSpecialCharacters_FileWithMultipleDots(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/path/to/file.multiple.dots.txt")
	if got != "file.multiple.dots.txt" {
		t.Errorf("findFileName(%q) = %q, want %q", "/path/to/file.multiple.dots.txt", got, "file.multiple.dots.txt")
	}
}

func TestFindFileNameWithSpecialCharacters_FilenameWithoutPath(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("filename.txt")
	if got != "filename.txt" {
		t.Errorf("findFileName(%q) = %q, want %q", "filename.txt", got, "filename.txt")
	}
}

func TestFindFileNameWithSpecialCharacters_TrailingSlash(t *testing.T) {
	c := fileServeController{}
	got := c.findFileName("/trailing/slash/")
	if got != "slash" {
		t.Errorf("findFileName(%q) = %q, want %q", "/trailing/slash/", got, "slash")
	}
}

// TestFindExtensionWithMultipleDots tests extension extraction with multiple dots
// Note: Implementation returns nameParts[1] (first extension after dot)
func TestFindExtensionWithMultipleDots_TarGz(t *testing.T) {
	c := fileServeController{}
	got := c.findExtension("/path/to/file.tar.gz")
	if got != "tar" {
		t.Errorf("findExtension(%q) = %q, want %q", "/path/to/file.tar.gz", got, "tar")
	}
}

func TestFindExtensionWithMultipleDots_TarBz2(t *testing.T) {
	c := fileServeController{}
	got := c.findExtension("/path/to/archive.tar.bz2")
	if got != "tar" {
		t.Errorf("findExtension(%q) = %q, want %q", "/path/to/archive.tar.bz2", got, "tar")
	}
}

func TestFindExtensionWithMultipleDots_HiddenFile(t *testing.T) {
	c := fileServeController{}
	got := c.findExtension("/path/.hiddenfile")
	if got != "hiddenfile" {
		t.Errorf("findExtension(%q) = %q, want %q", "/path/.hiddenfile", got, "hiddenfile")
	}
}

// fakeStorage is an in-memory StorageInterface for handler tests.
type fakeStorage struct {
	files    map[string][]byte
	exists   bool
	existsOK bool
	err      error
}

func (s *fakeStorage) Exists(path string) (bool, error) {
	if s.err != nil {
		return false, s.err
	}
	if s.existsOK {
		return s.exists, nil
	}
	_, ok := s.files[path]
	return ok, nil
}

func (s *fakeStorage) ReadFile(path string) ([]byte, error) {
	if s.err != nil {
		return nil, s.err
	}
	content, ok := s.files[path]
	if !ok {
		return nil, errFakeNotFound
	}
	return content, nil
}

var errFakeNotFound = errors.New("not found")

func TestHandler_ServesFileWithMimeType(t *testing.T) {
	storage := &fakeStorage{files: map[string][]byte{
		"/hello.txt": []byte("hi"),
	}}
	c := NewFileServeController(storage)

	req := httptest.NewRequest(http.MethodGet, "/files/hello.txt", nil)
	rec := httptest.NewRecorder()
	if res := c.Handler(rec, req); res != "" {
		t.Fatalf("Handler returned %q, want empty", res)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/octet-stream" {
		t.Errorf("Content-Type = %q, want application/octet-stream", ct)
	}
	if rec.Body.String() != "hi" {
		t.Errorf("body = %q, want %q", rec.Body.String(), "hi")
	}
}

func TestHandler_ServesPng(t *testing.T) {
	storage := &fakeStorage{files: map[string][]byte{
		"/img.png": {0x89, 0x50},
	}}
	c := NewFileServeController(storage)

	req := httptest.NewRequest(http.MethodGet, "/media/img.png", nil)
	rec := httptest.NewRecorder()
	if res := c.Handler(rec, req); res != "" {
		t.Fatalf("Handler returned %q, want empty", res)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "image/png" {
		t.Errorf("Content-Type = %q, want image/png", ct)
	}
}

func TestHandler_NotFound(t *testing.T) {
	c := NewFileServeController(&fakeStorage{files: map[string][]byte{}})

	req := httptest.NewRequest(http.MethodGet, "/files/missing.txt", nil)
	rec := httptest.NewRecorder()
	if res := c.Handler(rec, req); res != "File not found" {
		t.Errorf("Handler = %q, want 'File not found'", res)
	}
}

func TestHandler_StorageError(t *testing.T) {
	c := NewFileServeController(&fakeStorage{err: errors.New("boom")})

	req := httptest.NewRequest(http.MethodGet, "/files/x.txt", nil)
	rec := httptest.NewRecorder()
	if res := c.Handler(rec, req); res != "boom" {
		t.Errorf("Handler = %q, want 'boom'", res)
	}
}
