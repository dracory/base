# Filesystem Utilities

The filesystem package provides utilities for working with embedded file systems in Go applications.

## Features

- **EmbeddedFileToBytes**: Read files from embedded filesystem as byte slices
- **EmbeddedFileToString**: Read files from embedded filesystem as strings

## Installation

```go
import "github.com/dracory/base/filesystem"
```

## Usage

### Reading Embedded Files

```go
//go:embed templates/*.html
var templateFS embed.FS

// Read as bytes
bytes, err := filesystem.EmbeddedFileToBytes(templateFS, "templates/index.html")
if err != nil {
    log.Fatal(err)
}

// Read as string
content, err := filesystem.EmbeddedFileToString(templateFS, "templates/index.html")
if err != nil {
    log.Fatal(err)
}
fmt.Println(content)
```

## Functions

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

See the test files for comprehensive examples of how to use these utilities with embedded filesystems.

## Testing

The package includes comprehensive tests using embedded test data:

```bash
go test ./filesystem
```

## Dependencies

- Go 1.16+ (for embed support)
- No external dependencies
