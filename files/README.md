# File Utilities

The files package provides utilities for working with files in Go applications, including both local file operations and embedded filesystem access.

## Features

- **DownloadURL**: Download files from URLs to local filesystem
- **SaveToTempDir**: Save multipart files to temporary directory
- **EmbeddedFileToBytes**: Read files from embedded filesystem as byte slices
- **EmbeddedFileToString**: Read files from embedded filesystem as strings

## Installation

```go
import "github.com/dracory/base/files"
```

## Usage

### Downloading Files from URLs

```go
err := files.DownloadURL("https://example.com/file.pdf", "local/path/file.pdf")
if err != nil {
    log.Fatal(err)
}
```

### Saving Uploaded Files to Temp Directory

```go
// Assuming you have a multipart.File from an HTTP upload
tempPath, err := files.SaveToTempDir("uploaded.jpg", uploadedFile)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("File saved to: %s", tempPath)
```

### Reading Embedded Files

```go
//go:embed templates/*.html
var templateFS embed.FS

// Read as bytes
bytes, err := files.EmbeddedFileToBytes(templateFS, "templates/index.html")
if err != nil {
    log.Fatal(err)
}

// Read as string
content, err := files.EmbeddedFileToString(templateFS, "templates/index.html")
if err != nil {
    log.Fatal(err)
}
fmt.Println(content)
```

## Functions

### DownloadURL(url string, localFilepath string) error

Downloads a file from a URL to a local file path. Efficient because it writes as it downloads rather than loading the entire file into memory.

**Parameters:**
- `url`: URL to download from
- `localFilepath`: Local file path to save the file

**Returns:**
- `error`: Error if download fails

### SaveToTempDir(fileName string, file multipart.File) (string, error)

Saves a multipart file to a temporary directory with the appropriate file extension.

**Parameters:**
- `fileName`: Original filename (used for extension)
- `file`: Multipart file to save

**Returns:**
- `string`: Path to the saved temporary file
- `error`: Error if save fails

### EmbeddedFileToBytes(embeddedFileSystem embed.FS, path string) ([]byte, error)

Reads a file from the embedded filesystem and returns its content as a byte slice.

**Parameters:**
- `embeddedFileSystem`: The embedded filesystem (embed.FS)
- `path`: Path to the file within the embedded filesystem

**Returns:**
- `[]byte`: File content as bytes
- `error`: Error if file cannot be read

### EmbeddedFileToString(embeddedFileSystem embed.FS, path string) (string, error)

Reads a file from the embedded filesystem and returns its content as a string.

**Parameters:**
- `embeddedFileSystem`: The embedded filesystem (embed.FS)
- `path`: Path to the file within the embedded filesystem

**Returns:**
- `string`: File content as string
- `error`: Error if file cannot be read

## Examples

### Web Application File Upload

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("upload")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    tempPath, err := files.SaveToTempDir(header.Filename, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Process the file at tempPath...
    fmt.Fprintf(w, "File uploaded to: %s", tempPath)
}
```

### Embedded Template System

```go
//go:embed templates/*.html
var templateFS embed.FS

func renderTemplate(templateName string, data interface{}) (string, error) {
    content, err := files.EmbeddedFileToString(templateFS, "templates/"+templateName)
    if err != nil {
        return "", err
    }
    
    // Process template with data...
    return content, nil
}
```

## Testing

The package includes comprehensive tests for all functionality:

```bash
go test ./files
```

## Dependencies

- Go 1.16+ (for embed support)
- No external dependencies

## Migration Notes

This package was created by merging the previous `filesystem` package functionality into the existing `files` package to eliminate redundancy. If you were previously importing `github.com/dracory/base/filesystem`, update your imports to `github.com/dracory/base/files`.

**Previous:**
```go
import "github.com/dracory/base/filesystem"
```

**Current:**
```go
import "github.com/dracory/base/files"
```
