package markdown_test

import (
	"testing"

	"github.com/dracory/base/markdown"
)

func TestMarkdownToHtmlEmpty(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "empty string",
			markdown: "",
			html:     "",
		},
		{
			name:     "whitespace only",
			markdown: "   ",
			html:     "",
		},
		{
			name:     "whitespace with newlines",
			markdown: "   \n\t   ",
			html:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlBasic(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "plain text",
			markdown: "Hello World",
			html:     "<p>Hello World</p>\n",
		},
		{
			name:     "simple paragraph",
			markdown: "Hello world",
			html:     "<p>Hello world</p>\n",
		},
		{
			name:     "multiple paragraphs",
			markdown: "First paragraph\n\nSecond paragraph",
			html:     "<p>First paragraph</p>\n<p>Second paragraph</p>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlFormatting(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "bold text",
			markdown: "**bold**",
			html:     "<p><strong>bold</strong></p>\n",
		},
		{
			name:     "italic text",
			markdown: "*italic*",
			html:     "<p><em>italic</em></p>\n",
		},
		{
			name:     "bold text with spaces",
			markdown: "**bold text**",
			html:     "<p><strong>bold text</strong></p>\n",
		},
		{
			name:     "italic text with spaces",
			markdown: "*italic text*",
			html:     "<p><em>italic text</em></p>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlHeadings(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "heading 1",
			markdown: "# Heading",
			html:     "<h1 id=\"heading\">Heading</h1>\n",
		},
		{
			name:     "heading 1 with number",
			markdown: "# Heading 1",
			html:     "<h1 id=\"heading-1\">Heading 1</h1>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlLists(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "unordered list",
			markdown: "- item1\n- item2",
			html:     "<ul>\n<li>item1</li>\n<li>item2</li>\n</ul>\n",
		},
		{
			name:     "unordered list with spaces",
			markdown: "- Item 1\n- Item 2",
			html:     "<ul>\n<li>Item 1</li>\n<li>Item 2</li>\n</ul>\n",
		},
		{
			name:     "ordered list",
			markdown: "1. First\n2. Second",
			html:     "<ol>\n<li>First</li>\n<li>Second</li>\n</ol>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlCode(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "code block",
			markdown: "```\ncode\n```",
			html:     "<pre><code>code\n</code></pre>\n",
		},
		{
			name:     "code block with language",
			markdown: "```go\nfunc main() {}\n```",
			html:     "<pre><code class=\"language-go\">func main() {}\n</code></pre>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlLinksAndImages(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		html     string
	}{
		{
			name:     "link",
			markdown: "[Google](https://www.google.com)",
			html:     "<p><a href=\"https://www.google.com\">Google</a></p>\n",
		},
		{
			name:     "image",
			markdown: "![alt text](image.jpg)",
			html:     "<p><img src=\"image.jpg\" alt=\"alt text\" /></p>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdown.MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}

func TestMarkdownToHtmlBlockquote(t *testing.T) {
	input := `> This is a blockquote
> With multiple lines
> And some **bold** text`

	expected := `<blockquote>
<p>This is a blockquote<br />
With multiple lines<br />
And some <strong>bold</strong> text</p>
</blockquote>
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

func TestMarkdownToHtmlTable(t *testing.T) {
	input := `| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |`

	expected := `<table>
<thead>
<tr>
<th>Header 1</th>
<th>Header 2</th>
</tr>
</thead>
<tbody>
<tr>
<td>Cell 1</td>
<td>Cell 2</td>
</tr>
<tr>
<td>Cell 3</td>
<td>Cell 4</td>
</tr>
</tbody>
</table>
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

func TestMarkdownToHtmlHorizontalRule(t *testing.T) {
	input := `---`

	expected := `<hr />
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

func TestMarkdownToHtmlReferenceLink(t *testing.T) {
	input := `Final paragraph with a [reference link][1].

[1]: https://example.com/reference`

	expected := `<p>Final paragraph with a <a href="https://example.com/reference">reference link</a>.</p>
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

func TestMarkdownToHtmlNestedLists(t *testing.T) {
	input := `1. First ordered item
2. Second ordered item
   - Nested unordered item
   - Another nested item
3. Third ordered item`

	expected := `<ol>
<li>First ordered item</li>
<li>Second ordered item
<ul>
<li>Nested unordered item</li>
<li>Another nested item</li>
</ul>
</li>
<li>Third ordered item</li>
</ol>
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

func TestMarkdownToHtmlInlineCode(t *testing.T) {
	input := `This is a paragraph with **bold text** and *italic text*. You can also use ` + "`inline code`" + ` within text.`

	expected := `<p>This is a paragraph with <strong>bold text</strong> and <em>italic text</em>. You can also use <code>inline code</code> within text.</p>
`

	result, err := markdown.MarkdownToHtml(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}
