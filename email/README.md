# Email Package

This package provides email functionality for the Dracory framework. It includes:

- SMTP email sending
- Responsive HTML email templates
- Plain text conversion from HTML

## Usage

### Sending Emails

```go
import (
    "github.com/yourusername/dracory/base/email"
    "log/slog"
)

// Create a logger
logger := slog.Default()

// Create an email sender
sender := email.NewSMTPSender(email.Config{
    Host:     "smtp.example.com",
    Port:     "587",
    Username: "username",
    Password: "password",
    Logger:   logger,
})

// Send an email
err := sender.Send(email.SendOptions{
    From:     "sender@example.com",
    To:       []string{"recipient@example.com"},
    Subject:  "Test Email",
    HtmlBody: "<h1>Hello World</h1><p>This is a test email.</p>",
})

if err != nil {
    // Handle error
}
```

### Using Email Templates

```go
import "github.com/yourusername/dracory/base/email"

// Create email content
content := "<h1 style='" + email.StyleHeading + "'>Welcome!</h1>" +
           "<p style='" + email.StyleParagraph + "'>Thank you for registering.</p>" +
           "<a href='https://example.com/confirm' style='" + email.StyleButton + "'>Confirm Email</a>"

// Generate email template
template := email.DefaultTemplate(email.TemplateOptions{
    Title:   "Welcome to Our App",
    Content: content,
    AppName: "My Application",
    HeaderLinks: map[string]string{
        "Login": "https://example.com/login",
    },
})

// Use the template in SendOptions
sender.Send(email.SendOptions{
    From:     "sender@example.com",
    To:       []string{"recipient@example.com"},
    Subject:  "Welcome to Our App",
    HtmlBody: template,
})
```

## Customization

The email template can be customized by providing different options to the `DefaultTemplate` function:

- `Title`: The email title (appears in the browser tab)
- `Content`: The HTML content of the email
- `AppName`: The application name (appears in header and footer)
- `HeaderBackgroundColor`: The background color of the header (default: #17A2B8)
- `Year`: The copyright year (default: current year)
- `HeaderLinks`: A map of link text to URLs for the header

## Style Constants

The package provides comprehensive style constants for consistent email styling across all projects. These constants define inline CSS styles required for proper email client rendering.

### Typography Styles

- `StyleHeading1` - Primary heading (24px, bold)
- `StyleHeading2` - Secondary heading (20px, bold)
- `StyleHeading3` - Tertiary heading (18px, bold)
- `StyleParagraph` - Standard paragraph (16px, 1.6 line height)
- `StyleSmall` - Small text for disclaimers (14px)

### Button Styles

- `StyleButtonPrimary` - Main call-to-action (blue background)
- `StyleButtonSecondary` - Secondary action (outlined, no background)
- `StyleButtonSuccess` - Positive actions (green background)
- `StyleButtonDanger` - Destructive actions (red background)
- `StyleButtonSmall` - Smaller button variant

### Layout Styles

- `StyleContainer` - Centered container (max-width 600px)
- `StyleSection` - Content section with light background
- `StyleDivider` - Horizontal rule/separator
- `StyleCard` - Bordered card container

### Alert Styles

- `StyleAlertInfo` - Informational messages (blue)
- `StyleAlertSuccess` - Success messages (green)
- `StyleAlertWarning` - Warning messages (yellow)
- `StyleAlertDanger` - Error/danger messages (red)

### List Styles

- `StyleListUnordered` - Unordered (bulleted) lists
- `StyleListOrdered` - Ordered (numbered) lists
- `StyleListItem` - Individual list items

### Table Styles

- `StyleTable` - Base table style
- `StyleTableHead` - Table header cells
- `StyleTableCell` - Table data cells

### Utility Styles

- `StyleTextCenter` - Center text horizontally
- `StyleTextRight` - Align text to the right
- `StyleTextMuted` - Muted/secondary text (gray)
- `StyleTextPrimary` - Primary brand color text (blue)
- `StyleTextSuccess` - Success state text (green)
- `StyleTextDanger` - Danger/error state text (red)
- `StyleTextWarning` - Warning state text (yellow)
- `StyleBgLight` - Light background color
- `StyleBgDark` - Dark background color

### Usage Example

```go
import (
    "github.com/dracory/base/email"
    "github.com/dracory/hb"
)

// Create styled email content
h1 := hb.Heading1().
    HTML("Welcome to Our Service").
    Style(email.StyleHeading1)

p := hb.Paragraph().
    HTML("Thank you for registering with us.").
    Style(email.StyleParagraph)

button := hb.Hyperlink().
    Text("Get Started").
    Href("https://example.com/start").
    Style(email.StyleButtonPrimary)

alert := hb.Div().
    HTML("Important: Please verify your email address.").
    Style(email.StyleAlertInfo)

content := hb.Div().Children([]hb.TagInterface{
    h1,
    p,
    button,
    alert,
}).ToHTML()

// Generate full email template
template := email.DefaultTemplate(email.TemplateOptions{
    Title:   "Welcome",
    Content: content,
    AppName: "My Application",
})
```

### Using Email Components (Laravel-style)

This package provides a set of pre-styled components that make building emails as easy as in Laravel. They return `*hb.Tag` objects that you can further customize if needed.

```go
import "github.com/dracory/base/email"

// Build a complex email body using components
content := email.Container(
    email.Heading("Welcome to Dracory!", 1),
    email.Paragraph("We're glad to have you on board."),
    email.Section(
        email.Heading("Your Profile Status", 2),
        email.Alert("Your email is not verified yet.", "warning"),
        email.Button("Verify Email", "https://example.com/verify"),
    ),
    email.Divider(),
    email.Table(
        []string{"Feature", "Status"},
        [][]string{
            {"API Access", "Enabled"},
            {"Storage", "10GB"},
        },
    ),
).ToHTML()

// Use the generated content in the template
htmlBody := email.DefaultTemplate(email.TemplateOptions{
    Title:   "Welcome",
    Content: content,
    AppName: "MyApp",
})
```


## Brand Color Customization

The package provides brand color variables that can be customized to match your project's brand identity.

### Available Color Variables

- `ColorPrimary` - Main brand color (default: `#007BFF` blue)
- `ColorSecondary` - Secondary elements (default: `#6C757D` gray)
- `ColorSuccess` - Success states (default: `#28A745` green)
- `ColorDanger` - Error/danger states (default: `#DC3545` red)
- `ColorWarning` - Warnings (default: `#FFC107` yellow)
- `ColorInfo` - Informational messages (default: `#17A2B8` teal)
- `ColorLight` - Light backgrounds (default: `#F8F9FA`)
- `ColorDark` - Dark elements (default: `#343A40`)

### Method 1: Override Individual Colors

Customize specific colors in your project's `init()` function:

```go
package main

import (
    "github.com/dracory/base/email"
)

func init() {
    // Set your brand's primary color (e.g., purple)
    email.ColorPrimary = "#6f42c1"
    
    // Set your brand's header color for templates
    email.ColorInfo = "#17A2B8" // Teal for Dracory brand
}
```

### Method 2: Set All Brand Colors at Once

Use `SetBrandColors()` to customize all colors in one call:

```go
func init() {
    email.SetBrandColors(
        "#6f42c1", // Primary (purple)
        "#6C757D", // Secondary (gray)
        "#28A745", // Success (green)
        "#DC3545", // Danger (red)
        "#FFC107", // Warning (yellow)
        "#17A2B8", // Info (teal)
        "#F8F9FA", // Light
        "#343A40", // Dark
    )
}
```

### Method 3: Custom Style Builders

For one-off custom colors, use style builder functions:

```go
// Create a custom purple button
purpleButton := email.ButtonStyle("#6f42c1", "#5a32a3", "white")

// Create an outlined button with custom color
outlinedPurple := email.ButtonStyleSecondary("#6f42c1")

// Create a custom alert with brand colors
customAlert := email.AlertStyle("#e9d5ff", "#d8b4fe", "#3b0764")
```

### Usage Example with Custom Brand

```go
package main

import (
    "github.com/dracory/base/email"
    "github.com/dracory/hb"
)

func init() {
    // Customize for a purple brand theme
    email.ColorPrimary = "#6f42c1"
}

func sendWelcomeEmail() {
    // This button will now be purple instead of blue
    button := hb.Hyperlink().
        Text("Get Started").
        Href("https://example.com/start").
        Style(email.StyleButtonPrimary)
    
    // This text will also use the purple brand color
    brandText := hb.Span().
        Text("Premium Features").
        Style(email.StyleTextPrimary)
    
    // Generate email with custom header color
    template := email.DefaultTemplate(email.TemplateOptions{
        Title:                 "Welcome",
        Content:               content,
        AppName:               "MyApp",
        HeaderBackgroundColor: "#6f42c1", // Match brand color
    })
}
```

### Important Notes

- Color variables must be set before using any styles that depend on them
- The best place to set colors is in your project's `init()` function
- Color changes affect all emails using the base package styles
- For template headers, use `HeaderBackgroundColor` in `TemplateOptions`
```
