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

The package provides style constants for consistent email styling:

- `StyleHeading`: Style for headings
- `StyleParagraph`: Style for paragraphs
- `StyleButton`: Style for buttons
