# Blogtheme Package

The `blogtheme` package provides a lightweight theme system specifically designed for blog content rendering. It converts JSON-structured block content into HTML with blog-appropriate styling.

## Features

- **JSON-based Input**: Accepts blocks as JSON strings for easy storage and transport
- **Blog-focused Styling**: Pre-configured CSS styles optimized for blog content
- **Lightweight**: Minimal block types focused on common blog content needs
- **Error Handling**: Graceful handling of unsupported block types
- **Extensible**: Easy to add new block renderers

## Usage

### Basic Theme Usage

```go
import (
    "github.com/dracory/base/blogtheme"
    "github.com/dracory/ui"
)

// Create blocks
blocks := []ui.BlockInterface{
    func() ui.BlockInterface {
        block := ui.NewBlock()
        block.SetType("paragraph")
        block.SetParameter("content", "Welcome to my blog!")
        return block
    }(),
    func() ui.BlockInterface {
        block := ui.NewBlock()
        block.SetType("heading")
        block.SetParameter("level", "1")
        block.SetParameter("content", "My First Post")
        return block
    }(),
}

// Convert to JSON
blocksJSON, err := ui.MarshalBlocksToJson(blocks)
if err != nil {
    log.Fatal(err)
}

// Create theme and render HTML
theme, err := blogtheme.New(blocksJSON)
if err != nil {
    log.Fatal(err)
}

html := theme.ToHtml()
style := theme.Style()

fmt.Println(html)
// Output: <p>Welcome to my blog!</p><h1 style="margin-bottom:20px;margin-top:20px;">My First Post</h1>
```

### Integration with Blog Systems

```go
// Example: Rendering blog post content
func renderBlogPost(contentJSON string) (string, string, error) {
    theme, err := blogtheme.New(contentJSON)
    if err != nil {
        return "", "", err
    }
    
    html := theme.ToHtml()
    style := theme.Style()
    
    return html, style, nil
}
```

## Supported Block Types

The blogtheme supports a focused set of block types commonly needed for blog content:

### Text Content
- **paragraph**: Standard text paragraphs
- **heading**: H1-H6 headings with configurable levels
- **text**: Inline text spans

### Media
- **image**: Images with URL support
- **hyperlink**: Links with URL and text content

### Lists
- **unordered_list**: Bulleted lists (UL elements)
- **ordered_list**: Numbered lists (OL elements)  
- **list_item**: List items (LI elements)

### Special
- **raw**: Raw HTML content for custom markup

## Block Parameters

### Common Parameters
- **content**: The main content/text of the block

### Heading-specific
- **level**: Heading level (1-6)

### Image-specific  
- **image_url**: URL of the image

### Hyperlink-specific
- **url**: Link URL
- **content**: Link text/HTML

## Styling

The blogtheme provides built-in CSS styling optimized for blog content:

```css
.BlogTitle {
    font-family: Roboto, sans-serif;
}
.BlogContent {
    font-family: Roboto, sans-serif;
}
h1 { 
    margin-bottom: 20px;
    font-size: 48px;
}
h2 { 
    margin-bottom: 20px;
    font-size: 36px;
}
h3 {
    margin-bottom: 20px;
    font-size: 24px;
}
h4 {
    margin-bottom: 20px;
    font-size: 18px;
}
h5 {
    margin-bottom: 20px;
    font-size: 16px;
}
h6 {
    margin-bottom: 20px;
    font-size: 14px;
}
```

## JSON Input Format

The theme expects blocks in JSON format. Here's an example:

```json
[
  {
    "id": "block-1",
    "type": "heading",
    "sequence": 1,
    "parentId": "",
    "content": "Blog Post Title",
    "attributes": {
      "level": "1"
    }
  },
  {
    "id": "block-2", 
    "type": "paragraph",
    "sequence": 2,
    "parentId": "",
    "content": "This is the blog post content.",
    "attributes": {}
  }
]
```

## Error Handling

### Invalid JSON
```go
theme, err := blogtheme.New("{ invalid json }")
if err != nil {
    // Handle JSON parsing error
    return "", err
}
```

### Unsupported Block Types
Unsupported block types are rendered with warning messages:

```html
<div class="alert alert-warning">
  Block custom_type renderer does not exist
</div>
```

## Comparison with Webtheme

The `blogtheme` is a simplified version of `webtheme` focused on blog content:

| Feature | Blogtheme | Webtheme |
|---------|-----------|----------|
| Input Format | JSON only | UI blocks or JSON |
| Block Types | 8 basic types | 17+ types |
| Styling | Blog-focused | General purpose |
| Dependencies | Minimal | Extended (forms, blockeditor) |
| Use Case | Blog content | CMS/Block editor |

## Dependencies

- `github.com/dracory/ui`: Block interface and JSON utilities
- `github.com/dracory/hb`: HTML building library  
- `github.com/dracory/blockeditor`: Block tree utilities
- `github.com/samber/lo`: Functional utilities
- `github.com/spf13/cast`: Type casting utilities

## License

This package is part of the Dracory base package and follows the same licensing terms.
