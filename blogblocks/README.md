# Blogblocks Package

The `blogblocks` package provides block editor definitions specifically designed for blog content creation. It offers a focused set of block types optimized for blog writing and content management.

## Features

- **Blog-focused Blocks**: Essential block types for blog content
- **Block Editor Integration**: Ready-to-use definitions for block editor systems
- **Bootstrap Icons**: Consistent iconography using Bootstrap icons
- **Form Fields**: Pre-configured form fields for each block type
- **HTML Rendering**: Built-in HTML rendering functions
- **Child Support**: Proper parent-child relationships for nested content

## Usage

### Basic Usage

```go
import (
    "github.com/dracory/base/blogblocks"
    "github.com/dracory/blockeditor"
)

// Get block editor definitions
definitions := blogblocks.BlockEditorDefinitions()

// Use in your block editor system
blockEditor := blockeditor.New(blockeditor.Config{
    Definitions: definitions,
    // ... other configuration
})
```

### Integration with Form Systems

```go
// Example: Creating a form field for block editor content
field := form.NewField(form.FieldOptions{
    Name:  "content",
    Label: "Content",
    Type:  form.FORM_FIELD_TYPE_BLOCKEDITOR,
    BlockEditorOptions: blockeditor.BlockEditorOptions{
        Definitions: blogblocks.BlockEditorDefinitions(),
    },
})
```

## Supported Block Types

The blogblocks package provides 7 focused block types:

### Text Content
- **heading**: H1-H6 headings with level selection
- **paragraph**: Standard text paragraphs

### Media
- **image**: Images with URL support and placeholder generation
- **hyperlink**: Links with URL, text, and target options

### Lists
- **unordered_list**: Bulleted lists that support list item children
- **ordered_list**: Numbered lists that support list item children
- **list_item**: Individual list items with HTML content

## Block Type Details

### Heading Block
- **Icon**: `bi bi-type-h1`
- **Fields**:
  - `level`: Select dropdown (Heading Level 1-6)
  - `content`: Textarea for HTML content
- **Rendering**: `<h{level}>content</h{level}>`

### Paragraph Block
- **Icon**: `bi bi-paragraph`
- **Fields**:
  - `content`: Textarea for HTML content
- **Rendering**: `<p>content</p>`

### Image Block
- **Icon**: `bi bi-image`
- **Fields**:
  - `image_url`: Textarea for image URL (supports base64)
- **Rendering**: `<img src="image_url">`
- **Placeholder**: Auto-generates placeholder image using Picsum if URL is empty

### Hyperlink Block
- **Icon**: `bi bi-link-45deg`
- **Fields**:
  - `url`: Textarea for link URL
  - `content`: Textarea for link text/HTML
  - `target`: Select dropdown (_blank, _self, _parent, _top)
- **Rendering**: `<a href="url" target="target">content</a>`

### Unordered List Block
- **Icon**: `bi bi-list-ul`
- **Fields**: None (container for list items)
- **Children**: Allowed - `list_item` only
- **Rendering**: `<ul>children</ul>`

### Ordered List Block
- **Icon**: `bi bi-list-ol`
- **Fields**: None (container for list items)
- **Children**: Allowed - `list_item` only
- **Rendering**: `<ol>children</ol>`

### List Item Block
- **Icon**: `bi bi-list`
- **Fields**:
  - `content`: Textarea for HTML content
- **Rendering**: `<li>content</li>`

## Block Parameters

### Common Parameters
- **content**: HTML content for text-based blocks

### Heading-specific
- **level`: Heading level (1-6)

### Image-specific
- **image_url**: URL of the image (supports base64 encoded images)

### Hyperlink-specific
- **url**: Link URL
- **content**: Link text/HTML
- **target**: Link target (_blank, _self, _parent, _top)

## Form Field Types

Each block definition uses appropriate form field types:

- **FORM_FIELD_TYPE_SELECT**: For dropdown selections (heading levels, link targets)
- **FORM_FIELD_TYPE_TEXTAREA**: For multi-line content (text, URLs, HTML)
- **FORM_FIELD_TYPE_HIDDEN**: For system fields (post IDs, etc.)

## HTML Rendering

Each block definition includes a `ToTag` function that converts block parameters to HTML:

```go
// Example: Heading block rendering
func (block *headingDefinition) ToTag(block ui.BlockInterface) *hb.Tag {
    level := block.Parameter("level")
    content := block.Parameter("content")
    
    if level == "" {
        level = "1"
    }
    
    return hb.NewTag("h"+level).
        HTMLIf(content != "", `Add heading text`).
        HTML(content)
}
```

## Parent-Child Relationships

The package properly handles nested content:

- **Lists**: Both unordered and ordered lists can contain list items
- **List Items**: Cannot contain other blocks (leaf nodes)
- **Text Blocks**: Standalone blocks without children

## Comparison with Webtheme

| Feature | Blogblocks | Webtheme |
|---------|------------|----------|
| Focus | Blog content | General CMS |
| Block Types | 7 focused types | 17+ comprehensive types |
| Complexity | Simple | Advanced |
| Dependencies | Minimal | Extended |
| Use Case | Blog writing | Full CMS |

## Dependencies

- `github.com/dracory/base/img`: Image placeholder generation
- `github.com/dracory/blockeditor`: Block editor framework
- `github.com/dracory/form`: Form field definitions
- `github.com/dracory/hb`: HTML building library
- `github.com/dracory/ui`: Block interface utilities

## Integration Examples

### Blog Editor Integration

```go
// Create blog content editor
editor := blockeditor.New(blockeditor.Config{
    Definitions: blogblocks.BlockEditorDefinitions(),
    Renderer: func(blocks []ui.BlockInterface) string {
        // Use blogtheme for rendering
        theme, err := blogtheme.New(blocksJSON)
        if err != nil {
            return "Error rendering content"
        }
        return theme.ToHtml()
    },
})
```

### Form Field Integration

```go
// Add block editor to blog post form
contentField := form.NewField(form.FieldOptions{
    Name:  "content",
    Label: "Content",
    Type:  form.FORM_FIELD_TYPE_BLOCKEDITOR,
    BlockEditorOptions: blockeditor.BlockEditorOptions{
        Definitions: blogblocks.BlockEditorDefinitions(),
        MaxHeight:    "400px",
    },
})
```

## Testing

The package includes comprehensive tests covering:

- Block definition structure validation
- HTML rendering functionality
- Form field configuration
- Icon presence
- Parent-child relationships
- Error handling

Run tests with:
```bash
go test ./blogblocks/...
```

## License

This package is part of the Dracory base package and follows the same licensing terms.
