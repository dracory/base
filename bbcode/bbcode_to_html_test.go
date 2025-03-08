package bbcode

import (
	"testing"
)

func TestBbcodeToHtml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty input", "", ""},

		// Paragraph
		{"paragraph", "[paragraph]Hello World[/paragraph]", "<p>Hello World</p>"},
		{"paragraph (short)", "[p]Hello World[/p]", "<p>Hello World</p>"},

		// Headings (full)
		{"heading1", "[heading1]Hello World[/heading1]", "<h1>Hello World</h1>"},
		{"heading2", "[heading2]Hello World[/heading2]", "<h2>Hello World</h2>"},
		{"heading3", "[heading3]Hello World[/heading3]", "<h3>Hello World</h3>"},
		{"heading4", "[heading4]Hello World[/heading4]", "<h4>Hello World</h4>"},
		{"heading5", "[heading5]Hello World[/heading5]", "<h5>Hello World</h5>"},
		{"heading6", "[heading6]Hello World[/heading6]", "<h6>Hello World</h6>"},

		// Headings (short)
		{"h1", "[h1]Hello World[/h1]", "<h1>Hello World</h1>"},
		{"h2", "[h2]Hello World[/h2]", "<h2>Hello World</h2>"},
		{"h3", "[h3]Hello World[/h3]", "<h3>Hello World</h3>"},
		{"h4", "[h4]Hello World[/h4]", "<h4>Hello World</h4>"},
		{"h5", "[h5]Hello World[/h5]", "<h5>Hello World</h5>"},
		{"h6", "[h6]Hello World[/h6]", "<h6>Hello World</h6>"},

		// Expanded formats
		{"bold formatting", "[bold]Hello World[/bold]", "<b>Hello World</b>"},
		{"italic formatting", "[italic]Hello World[/italic]", "<i>Hello World</i>"},
		{"underline formatting", "[underline]Hello World[/underline]", "<u>Hello World</u>"},
		{"strike formatting", "[strike]Hello World[/strike]", "<s>Hello World</s>"},

		// Short formats
		{"bold formatting (short)", "[b]Hello World[/b]", "<b>Hello World</b>"},
		{"italic formatting (short)", "[i]Hello World[/i]", "<i>Hello World</i>"},
		{"underline formatting (short)", "[u]Hello World[/u]", "<u>Hello World</u>"},
		{"strike formatting (short)", "[s]Hello World[/s]", "<s>Hello World</s>"},

		// Formatting
		{"color", "[color=red]Hello World[/color]", "<span style=\"color:red\">Hello World</span>"},
		{"size", "[size=12]Hello World[/size]", "<span style=\"font-size:12\">Hello World</span>"},

		// Preformatted, quote and code blocks
		{"preformatted text", "[pre]Hello World[/pre]", "<pre>Hello World</pre>"},
		{"quote", "[quote]Hello World[/quote]", "<blockquote>Hello World</blockquote>"},
		{"code block", "[code]Hello World[/code]", "<code>Hello World</code>"},
		{"code block with language", "[code=go]Hello World[/code]", "<code lang=\"go\">Hello World</code>"},

		// Lists
		{"unordered list", "[list]*Item 1\n*Item 2\n[/list]", "<ul><li>Item 1</li><li>Item 2</li></ul>"},
		{"ordered list", "[list=1]*Item 1\n*Item 2\n[/list]", "<ol><li>Item 1</li><li>Item 2</li></ol>"},
		{"list item", "[item]Item 1[/item]", "<li>Item 1</li>"},
		{"list", "[list][item]Item 1[/item][item]Item 2[/item][/list]", "<ul><li>Item 1</li><li>Item 2</li></ul>"},

		// Email
		{"email link", "[email]example@example.com[/email]", "<a href=\"mailto:example@example.com\">example@example.com</a>"},

		// Links
		{"url", "[url]https://example.com[/url]", "<a href=\"https://example.com\">https://example.com</a>"},
		{"url with text", "[url=https://example.com]Example[/url]", "<a href=\"https://example.com\">Example</a>"},

		// Images
		{"image", "[img]https://example.com/image.jpg[/img]", "<img src=\"https://example.com/image.jpg\" />"},

		// Other
		{"link", "[url]https://example.com[/url]", "<a href=\"https://example.com\">https://example.com</a>"},
		{"section", "[section]Hello World[/section]", "<section>Hello World</section>"},

		{"div + content", "[div]Hello World[/div]", "<div>Hello World</div>"},
		{"divider", "[divider]Hello World[/divider]", "<hr />Hello World<hr />"},
		{"divider [hr]", "[hr]", "<hr />"},
		{"divider [rule]", "[rule]", "<hr />"},
		{"divider [rule + content]", "[rule]Hello World[/rule]", "<hr />Hello World<hr />"},
		{"divider [hr + content]", "[hr]Hello World[/hr]", "<hr />Hello World<hr />"},

		// Breaks
		{"break [br + content]", "[br]Hello World[/br]", "<br />Hello World<br />"},
		{"break [break+content]", "[break]Hello World[/break]", "<br />Hello World<br />"},
		{"break [br]", "[br]", "<br />"},
		{"break [break]", "[break]", "<br />"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := BbcodeToHtml(test.input)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
