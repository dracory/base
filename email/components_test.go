package email

import (
	"strings"
	"testing"
)

func TestComponents(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		contains []string
	}{
		{
			name: "Heading1",
			html: Heading("Hello World", 1).ToHTML(),
			contains: []string{
				"<h1",
				"Hello World",
				StyleHeading1,
			},
		},
		{
			name: "Heading2",
			html: Heading("Subheading", 2).ToHTML(),
			contains: []string{
				"<h2",
				"Subheading",
				StyleHeading2,
			},
		},
		{
			name: "Paragraph",
			html: Paragraph("This is a paragraph.").ToHTML(),
			contains: []string{
				"<p",
				"This is a paragraph.",
				StyleParagraph,
			},
		},
		{
			name: "Button",
			html: Button("Click Me", "https://example.com").ToHTML(),
			contains: []string{
				"<a",
				"href=\"https://example.com\"",
				"Click Me",
				StyleButtonPrimary,
			},
		},
		{
			name: "Alert",
			html: Alert("Warning message", "warning").ToHTML(),
			contains: []string{
				"<div",
				"Warning message",
				StyleAlertWarning,
			},
		},
		{
			name: "Divider",
			html: Divider().ToHTML(),
			contains: []string{
				"<div",
				StyleDivider,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, substr := range tt.contains {
				if !strings.Contains(tt.html, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.html)
				}
			}
		})
	}
}

func TestComplexComponents(t *testing.T) {
	// Test Card with children
	card := Card(
		Heading("Card Title", 2),
		Paragraph("Card content goes here."),
		Button("Action", "#"),
	).ToHTML()

	expectedSubstrings := []string{
		StyleCard,
		"Card Title",
		"Card content goes here.",
		"Action",
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(card, substr) {
			t.Errorf("Card should contain %q", substr)
		}
	}
}

func TestTableComponent(t *testing.T) {
	headers := []string{"Name", "Value"}
	rows := [][]string{
		{"Item 1", "$10"},
		{"Item 2", "$20"},
	}

	table := Table(headers, rows).ToHTML()

	expectedSubstrings := []string{
		"<table",
		StyleTable,
		"Name",
		"Value",
		StyleTableHead,
		"Item 1",
		"$10",
		"Item 2",
		"$20",
		StyleTableCell,
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(table, substr) {
			t.Errorf("Table should contain %q", substr)
		}
	}
}
