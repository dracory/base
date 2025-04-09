# BBCode Package

The bbcode package provides BBCode to HTML conversion functionality for the Dracory framework. It enables converting BBCode formatted text into clean, valid HTML output.

## Features

- Convert BBCode text to HTML
- Support for common BBCode tags:
  - Text formatting (`[b]`, `[i]`, `[u]`, `[s]`)
  - Links (`[url]`)
  - Images (`[img]`)
  - Lists (`[list]`, `[ul]`, `[ol]`)
  - Code blocks (`[code]`)
  - Quotes (`[quote]`)
  - Colors and sizes (`[color]`, `[size]`)
- Proper error handling
- Empty input handling
- XSS protection
- Configurable tag support

## Installation

```bash
go get github.com/dracory/base/bbcode
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/dracory/base/bbcode"
)

func main() {
    // Example BBCode text
    bbcodeText := `[b]Hello World[/b]
    
This is a [i]formatted[/i] text.
    
[url=https://example.com]Click here[/url]
    
[list]
[*]List item 1
[*]List item 2
[/list]`

    // Convert BBCode to HTML
    html, err := bbcode.BBCodeToHtml(bbcodeText)
    if err != nil {
        fmt.Printf("Error converting BBCode: %v\n", err)
        return
    }

    fmt.Println(html)
}
```

## Configuration

The package will support the following configuration options:

- Enable/disable specific BBCode tags
- Custom tag definitions
- HTML sanitization options
- Output format (HTML5, XHTML)
- Custom CSS class names

## Error Handling

The `BBCodeToHtml` function will return an error if:
- The BBCode conversion fails
- The input text cannot be processed
- Invalid BBCode syntax is encountered

Empty or whitespace-only input will be handled gracefully and return an empty string without error.

## Security

The package will implement security measures:
- XSS protection through HTML sanitization
- URL validation for links and images
- Configurable security policies
- Safe default settings

## License

This package is part of the dracory/base project and is licensed under the same terms. See the main [README.md](../README.md) for license information. 