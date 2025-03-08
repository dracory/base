package bbcode

import (
	"regexp"
)

// Replacement defines a BBCode to HTML replacement pattern
type replacement struct {
	BBCode string
	HTML   string
}

// BbcodeToHtml takes a string written in BBCode and returns the HTML
// representation of it.
//
// The function processes replacements in the order they are defined,
// ensuring that more specific patterns are processed before more general ones.
//
// The function returns the HTML representation of the input string.
func BbcodeToHtml(input string) string {
	// Process replacements in the defined order
	for _, replacement := range getReplacements() {
		re := regexp.MustCompile(replacement.BBCode)
		input = re.ReplaceAllString(input, replacement.HTML)
	}

	return input
}

// getReplacements returns the slice of BBCode to HTML replacements
// This is kept as a private function to avoid polluting the global namespace
func getReplacements() []replacement {

	// Define content pattern for reuse
	var content = `([\S\s]+?.*?|\s?)`

	return []replacement{
		// Tags with content should be processed first to avoid partial matches

		// Divider with content
		{`\[divider\]` + content + `\[/divider\]`, `<hr />$1<hr />`},
		{`\[rule\]` + content + `\[/rule\]`, `<hr />$1<hr />`},
		{`\[hr\]` + content + `\[/hr\]`, `<hr />$1<hr />`},

		// Break with content
		{`\[break\]` + content + `\[/break\]`, `<br />$1<br />`},
		{`\[br\]` + content + `\[/br\]`, `<br />$1<br />`},

		// Paragraph
		{`\[paragraph\]` + content + `\[/paragraph\]`, `<p>$1</p>`},
		{`\[p\]` + content + `\[/p\]`, `<p>$1</p>`},

		// Headings (full)
		{`\[heading1\]` + content + `\[/heading1\]`, `<h1>$1</h1>`},
		{`\[heading2\]` + content + `\[/heading2\]`, `<h2>$1</h2>`},
		{`\[heading3\]` + content + `\[/heading3\]`, `<h3>$1</h3>`},
		{`\[heading4\]` + content + `\[/heading4\]`, `<h4>$1</h4>`},
		{`\[heading5\]` + content + `\[/heading5\]`, `<h5>$1</h5>`},
		{`\[heading6\]` + content + `\[/heading6\]`, `<h6>$1</h6>`},

		// Headings (short)
		{`\[h1\]` + content + `\[/h1\]`, `<h1>$1</h1>`},
		{`\[h2\]` + content + `\[/h2\]`, `<h2>$1</h2>`},
		{`\[h3\]` + content + `\[/h3\]`, `<h3>$1</h3>`},
		{`\[h4\]` + content + `\[/h4\]`, `<h4>$1</h4>`},
		{`\[h5\]` + content + `\[/h5\]`, `<h5>$1</h5>`},
		{`\[h6\]` + content + `\[/h6\]`, `<h6>$1</h6>`},

		// Preformatted, quote and code blocks
		{`\[pre\]` + content + `\[/pre\]`, `<pre>$1</pre>`},
		{`\[quote\]` + content + `\[/quote\]`, `<blockquote>$1</blockquote>`},
		{`\[code\]` + content + `\[/code\]`, `<code>$1</code>`},
		{`\[code=(.*?)\]` + content + `\[/code\]`, `<code lang="$1">$2</code>`},

		// Lists
		{`\[list\]` + content + `\[/list\]`, `<ul>$1</ul>`},
		{`\[list=1\]` + content + `\[/list\]`, `<ol>$1</ol>`},
		{`\[item\]` + content + `\[/item\]`, `<li>$1</li>`},
		{"\\*" + content + "\n", `<li>$1</li>`},

		// Formatting
		{`\[bold\]` + content + `\[/bold\]`, `<b>$1</b>`},
		{`\[italic\]` + content + `\[/italic\]`, `<i>$1</i>`},
		{`\[underline\]` + content + `\[/underline\]`, `<u>$1</u>`},
		{`\[strike\]` + content + `\[/strike\]`, `<s>$1</s>`},
		{`\[color=(.*?)\]` + content + `\[/color\]`, `<span style="color:$1">$2</span>`},
		{`\[size=(.*?)\]` + content + `\[/size\]`, `<span style="font-size:$1">$2</span>`},

		// Formatted text (short)
		{`\[b\]` + content + `\[/b\]`, `<b>$1</b>`},
		{`\[i\]` + content + `\[/i\]`, `<i>$1</i>`},
		{`\[u\]` + content + `\[/u\]`, `<u>$1</u>`},
		{`\[s\]` + content + `\[/s\]`, `<s>$1</s>`},

		// Email
		{`\[email\]` + content + `\[/email\]`, `<a href="mailto:$1">$1</a>`},

		// Links
		{`\[url\]` + content + `\[/url\]`, `<a href="$1">$1</a>`},
		{`\[url=(.*?)\]` + content + `\[/url\]`, `<a href="$1">$2</a>`},

		// Images
		{`\[img\]` + content + `\[/img\]`, `<img src="$1" />`},
		{`\[img=(.*?)\]` + content + `\[/img\]`, `<img src="$1" alt="$2" />`},

		// Section
		{`\[section\]` + content + `\[/section\]`, `<section>$1</section>`},
		{`\[div\]` + content + `\[/div\]`, `<div>$1</div>`},

		// Simple tags should be processed last to avoid conflicts

		// Divider
		{`\[divider\]`, `<hr />`},
		{`\[rule\]`, `<hr />`},
		{`\[hr\]`, `<hr />`},
		{`\[hr=(.*?)\]`, `<hr style="border-color:$1" />`},

		// Line break
		{`\[break\]`, `<br />`},
		{`\[br\]`, `<br />`},
	}
}
