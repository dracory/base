package markdown

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

// ToHtml converts a markdown text to html
//
// 1. the text is trimmed of any white spaces
// 2. if the text is empty, it returns an empty string
// 3. the text is converted to html using the goldmark library
func MarkdownToHtml(text string) (string, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return "", nil
	}

	var buf bytes.Buffer

	md := goldmark.New(
		// goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
		// html.WithHardWraps(),			html.WithXHTML(),
		),
	)
	if err := md.Convert([]byte(text), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
