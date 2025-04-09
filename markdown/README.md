# Markdown Package

The markdown package provides Markdown to HTML conversion functionality for the Dracory framework. It uses the Goldmark library to convert Markdown text into clean, valid HTML with support for GitHub Flavored Markdown (GFM).

## Features

- Convert Markdown text to HTML
- Support for GitHub Flavored Markdown (GFM)
- Automatic heading IDs
- XHTML output
- Hard line breaks
- Proper error handling
- Empty input handling

## Installation

```bash
go get github.com/dracory/base/markdown
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/dracory/base/markdown"
)

func main() {
    // Example Markdown text
    markdownText := `# Hello World
    
This is a **bold** text and this is *italic*.
    
- List item 1
- List item 2
    
[Link](https://example.com)`

    // Convert Markdown to HTML
    html, err := markdown.MarkdownToHtml(markdownText)
    if err != nil {
        fmt.Printf("Error converting markdown: %v\n", err)
        return
    }

    fmt.Println(html)
}
```

## Configuration

The package uses Goldmark with the following default configuration:

- GitHub Flavored Markdown (GFM) extension enabled
- Automatic heading IDs
- XHTML output
- Unsafe HTML (allows raw HTML in markdown)
- Hard line breaks

## Error Handling

The `MarkdownToHtml` function returns an error if:
- The markdown conversion fails
- The input text cannot be processed

Empty or whitespace-only input is handled gracefully and returns an empty string without error.

## Dependencies

- [Goldmark](https://github.com/yuin/goldmark) - A markdown parser and renderer
- [Goldmark GFM Extension](https://github.com/yuin/goldmark#github-flavored-markdown) - GitHub Flavored Markdown support

## License

This package is part of the dracory/base project and is licensed under the same terms. See the main [README.md](../README.md) for license information. 