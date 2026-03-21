# Webtheme Package

The `webtheme` package provides a flexible theme system for rendering UI blocks to HTML. It's designed to work with block editors and content management systems.

## Features

- **Block Rendering**: Convert UI blocks to HTML with customizable renderers
- **Block Editor Definitions**: Pre-defined block types for use in block editors
- **Bootstrap Icons**: Comprehensive Bootstrap icon list for icon selection
- **Extensible**: Easy to add new block types and renderers

## Usage

### Basic Theme Usage

```go
import (
    "github.com/dracory/base/webtheme"
    "github.com/dracory/ui"
)

// Create blocks
paragraphBlock := ui.NewBlock()
paragraphBlock.SetType("paragraph")
paragraphBlock.SetParameter("content", "Hello, World!")
paragraphBlock.SetParameter("status", "published")

headingBlock := ui.NewBlock()
headingBlock.SetType("heading")
headingBlock.SetParameter("level", "1")
headingBlock.SetParameter("content", "Welcome")
headingBlock.SetParameter("status", "published")

// Create theme and render HTML
blocks := []ui.BlockInterface{paragraphBlock, headingBlock}
theme := webtheme.New(blocks)
html := theme.ToHtml()

fmt.Println(html)
// Output: <p>Hello, World!</p><h1 style="margin-bottom:20px;margin-top:20px;">Welcome</h1>
```

### Block Editor Integration

```go
// Get block editor definitions for your block editor
definitions := webtheme.BlockEditorDefinitions()

// Each definition contains:
// - Type: Block type identifier
// - Icon: Icon for the block editor UI
// - Fields: Form fields for block configuration
// - ToTag: Optional custom rendering function
// - Wrapper: Optional wrapper function
```

## Supported Block Types

The theme comes with pre-defined renderers for:

- **paragraph**: Text paragraphs
- **heading**: H1-H6 headings
- **text**: Inline text spans
- **image**: Images with URL support
- **icon**: Bootstrap icons
- **hyperlink**: Links with target options
- **container**: Bootstrap containers
- **row**: Bootstrap rows
- **column**: Responsive columns (xs, sm, md, lg, xl, xxl)
- **section**: HTML sections
- **div**: Generic div elements
- **unordered_list**: UL elements
- **ordered_list**: OL elements
- **list_item**: LI elements
- **breadcrumbs**: Breadcrumb navigation
- **raw_html**: Raw HTML content

## Block Parameters

All blocks support the following standard parameters:

- **status**: Set to "published" for the block to render
- Plus block-specific parameters (see individual block types)

### Common Parameters by Block Type

#### Paragraph
- `content`: HTML content of the paragraph

#### Heading
- `level`: Heading level (1-6)
- `content`: HTML content of the heading

#### Image
- `image_url`: URL of the image (supports base64)

#### Icon
- `icon`: Bootstrap icon class (e.g., "bi bi-house")

#### Hyperlink
- `url`: Link URL
- `content`: Link text/HTML
- `target`: Link target (_blank, _self, _parent, _top)

#### Column
- `width_xs`: Column width on extra small screens
- `width_sm`: Column width on small screens
- `width_md`: Column width on medium screens
- `width_lg`: Column width on large screens
- `width_xl`: Column width on extra large screens
- `width_xxl`: Column width on extra extra large screens

#### Breadcrumbs
- `breadcrumb1_url`, `breadcrumb1_text`: First breadcrumb
- `breadcrumb2_url`, `breadcrumb2_text`: Second breadcrumb
- `breadcrumb3_url`, `breadcrumb3_text`: Third breadcrumb

## Bootstrap Icons

The package includes a comprehensive list of Bootstrap icons:

```go
icons := webtheme.bootstrapIconList()
for _, icon := range icons {
    fmt.Printf("Name: %s, Icon: %s\n", icon.Name, icon.Icon)
}
```

## Custom Block Renderers

You can extend the theme by adding custom renderers:

```go
// Create a custom theme
theme := webtheme.New(blocks)

// The theme automatically registers all built-in renderers
// You can extend by modifying the theme struct or creating custom block types
```

## Integration with CMS

This package is designed to work seamlessly with content management systems:

```go
// Example CMS integration
frontend := cmsFrontend.New(cmsFrontend.Config{
    BlockEditorRenderer: func(blocks []ui.BlockInterface) string {
        return webtheme.New(blocks).ToHtml()
    },
    Store: registry.GetCmsStore(),
    // ... other configuration
})
```

## Dependencies

- `github.com/dracory/ui`: Block interface and utilities
- `github.com/dracory/hb`: HTML building library
- `github.com/dracory/blockeditor`: Block editor definitions
- `github.com/dracory/form`: Form field definitions
- `github.com/dracory/base/img`: Image utilities (for placeholders)

## License

This package is part of the Dracory base package and follows the same licensing terms.
