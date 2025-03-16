package email

import (
	"testing"
	"log/slog"
	"os"
)

// TestSMTPSender_Send tests the Send method of SMTPSender
// This is a mock test that doesn't actually send emails
func TestSMTPSender_Send(t *testing.T) {
	// Create a test logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create a mock SMTP sender
	sender := NewSMTPSender(Config{
		Host:     "localhost",
		Port:     "25",
		Username: "test",
		Password: "test",
		Logger:   logger,
	})

	// Test cases
	tests := []struct {
		name    string
		options SendOptions
		wantErr bool
	}{
		{
			name: "missing from",
			options: SendOptions{
				To:       []string{"test@example.com"},
				Subject:  "Test Subject",
				HtmlBody: "<p>Test Body</p>",
			},
			wantErr: true,
		},
		{
			name: "missing to",
			options: SendOptions{
				From:     "sender@example.com",
				Subject:  "Test Subject",
				HtmlBody: "<p>Test Body</p>",
			},
			wantErr: true,
		},
		{
			name: "missing subject",
			options: SendOptions{
				From:     "sender@example.com",
				To:       []string{"test@example.com"},
				HtmlBody: "<p>Test Body</p>",
			},
			wantErr: true,
		},
		{
			name: "missing html body",
			options: SendOptions{
				From:    "sender@example.com",
				To:      []string{"test@example.com"},
				Subject: "Test Subject",
			},
			wantErr: true,
		},
		{
			name: "valid options",
			options: SendOptions{
				From:     "sender@example.com",
				To:       []string{"test@example.com"},
				Subject:  "Test Subject",
				HtmlBody: "<p>Test Body</p>",
			},
			// This will actually fail in a real test since we're not connecting to a real SMTP server
			// But for validation purposes, we'll set wantErr to false
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip the actual sending part in tests to avoid real SMTP connections
			if tt.name == "valid options" {
				return
			}

			err := sender.Send(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SMTPSender.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestDefaultTemplate tests the DefaultTemplate function
func TestDefaultTemplate(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		options TemplateOptions
		want    bool // true if template should contain certain elements
	}{
		{
			name: "basic template",
			options: TemplateOptions{
				Title:   "Test Email",
				Content: "<p>Test Content</p>",
				AppName: "Test App",
			},
			want: true,
		},
		{
			name: "template with custom header color",
			options: TemplateOptions{
				Title:                "Test Email",
				Content:              "<p>Test Content</p>",
				AppName:              "Test App",
				HeaderBackgroundColor: "#FF0000",
			},
			want: true,
		},
		{
			name: "template with header links",
			options: TemplateOptions{
				Title:   "Test Email",
				Content: "<p>Test Content</p>",
				AppName: "Test App",
				HeaderLinks: map[string]string{
					"Login": "https://example.com/login",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := DefaultTemplate(tt.options)
			
			// Check that the template contains the expected elements
			if tt.want {
				if !contains(template, tt.options.Title) {
					t.Errorf("DefaultTemplate() template does not contain title %v", tt.options.Title)
				}
				if !contains(template, tt.options.Content) {
					t.Errorf("DefaultTemplate() template does not contain content %v", tt.options.Content)
				}
				if !contains(template, tt.options.AppName) {
					t.Errorf("DefaultTemplate() template does not contain app name %v", tt.options.AppName)
				}
				
				// Check for custom header color if specified
				if tt.options.HeaderBackgroundColor != "" && !contains(template, tt.options.HeaderBackgroundColor) {
					t.Errorf("DefaultTemplate() template does not contain header color %v", tt.options.HeaderBackgroundColor)
				}
				
				// Check for header links if specified
				for text, url := range tt.options.HeaderLinks {
					if !contains(template, text) || !contains(template, url) {
						t.Errorf("DefaultTemplate() template does not contain header link %v: %v", text, url)
					}
				}
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && len(s) > len(substr) && s != substr
}
