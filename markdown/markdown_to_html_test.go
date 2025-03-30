package markdown

import (
	"testing"
)

func TestMarkdownToHtml(t *testing.T) {
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
			name:     "plain text",
			markdown: "Hello World",
			html:     "<p>Hello World</p>\n",
		},
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
			name:     "heading",
			markdown: "# Heading",
			html:     "<h1 id=\"heading\">Heading</h1>\n",
		},
		{
			name:     "unordered list",
			markdown: "- item1\n- item2",
			html:     "<ul>\n<li>item1</li>\n<li>item2</li>\n</ul>\n",
		},
		{
			name:     "code block",
			markdown: "```\ncode\n```",
			html:     "<pre><code>code\n</code></pre>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MarkdownToHtml(tt.markdown)
			if err != nil {
				t.Fatal(err)
			}
			if result != tt.html {
				t.Fatalf("got %q, want %q", result, tt.html)
			}
		})
	}
}
