package email

import (
	"strings"
	"testing"
)

func TestTypographyStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleHeading1",
			style: StyleHeading1,
			contains: []string{
				"font-size:24px",
				"font-weight:600",
				"color:#333333",
			},
		},
		{
			name:  "StyleHeading2",
			style: StyleHeading2,
			contains: []string{
				"font-size:20px",
				"font-weight:600",
			},
		},
		{
			name:  "StyleHeading3",
			style: StyleHeading3,
			contains: []string{
				"font-size:18px",
				"font-weight:600",
			},
		},
		{
			name:  "StyleParagraph",
			style: StyleParagraph,
			contains: []string{
				"font-size:16px",
				"line-height:1.6",
			},
		},
		{
			name:  "StyleSmall",
			style: StyleSmall,
			contains: []string{
				"font-size:14px",
				"color:#666666",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestButtonStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleButtonPrimary",
			style: StyleButtonPrimary,
			contains: []string{
				"display: inline-block",
				"background-color: #007BFF",
				"color: white",
				"border-radius: 6px",
			},
		},
		{
			name:  "StyleButtonSecondary",
			style: StyleButtonSecondary,
			contains: []string{
				"display: inline-block",
				"background-color: transparent",
				"color: #007BFF",
				"border: 2px solid #007BFF",
			},
		},
		{
			name:  "StyleButtonSuccess",
			style: StyleButtonSuccess,
			contains: []string{
				"background-color: #28A745",
				"color: white",
			},
		},
		{
			name:  "StyleButtonDanger",
			style: StyleButtonDanger,
			contains: []string{
				"background-color: #DC3545",
				"color: white",
			},
		},
		{
			name:  "StyleButtonSmall",
			style: StyleButtonSmall,
			contains: []string{
				"padding: 8px 16px",
				"font-size: 14px",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestLayoutStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleContainer",
			style: StyleContainer,
			contains: []string{
				"max-width: 600px",
				"margin: 0 auto",
			},
		},
		{
			name:  "StyleSection",
			style: StyleSection,
			contains: []string{
				"background-color: #f8f9fa",
				"border-radius: 6px",
			},
		},
		{
			name:  "StyleDivider",
			style: StyleDivider,
			contains: []string{
				"height: 1px",
				"background-color: #dee2e6",
			},
		},
		{
			name:  "StyleCard",
			style: StyleCard,
			contains: []string{
				"border: 1px solid #dee2e6",
				"border-radius: 8px",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestAlertStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleAlertInfo",
			style: StyleAlertInfo,
			contains: []string{
				"background-color: #D1ECF1",
				"border: 1px solid #BEE5EB",
				"color: #0C5460",
			},
		},
		{
			name:  "StyleAlertSuccess",
			style: StyleAlertSuccess,
			contains: []string{
				"background-color: #D4EDDA",
				"color: #155724",
			},
		},
		{
			name:  "StyleAlertWarning",
			style: StyleAlertWarning,
			contains: []string{
				"background-color: #FFF3CD",
				"color: #856404",
			},
		},
		{
			name:  "StyleAlertDanger",
			style: StyleAlertDanger,
			contains: []string{
				"background-color: #F8D7DA",
				"color: #721C24",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestListStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleListUnordered",
			style: StyleListUnordered,
			contains: []string{
				"padding-left: 20px",
			},
		},
		{
			name:  "StyleListOrdered",
			style: StyleListOrdered,
			contains: []string{
				"padding-left: 20px",
			},
		},
		{
			name:  "StyleListItem",
			style: StyleListItem,
			contains: []string{
				"line-height: 1.5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestTableStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:  "StyleTable",
			style: StyleTable,
			contains: []string{
				"width: 100%",
				"border-collapse: collapse",
			},
		},
		{
			name:  "StyleTableHead",
			style: StyleTableHead,
			contains: []string{
				"background-color: #f8f9fa",
				"font-weight: 600",
			},
		},
		{
			name:  "StyleTableCell",
			style: StyleTableCell,
			contains: []string{
				"border: 1px solid #dee2e6",
				"padding: 12px",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestUtilityStyles(t *testing.T) {
	tests := []struct {
		name     string
		style    string
		contains []string
	}{
		{
			name:     "StyleTextCenter",
			style:    StyleTextCenter,
			contains: []string{"text-align: center"},
		},
		{
			name:     "StyleTextRight",
			style:    StyleTextRight,
			contains: []string{"text-align: right"},
		},
		{
			name:     "StyleTextMuted",
			style:    StyleTextMuted,
			contains: []string{"color: #6c757d"},
		},
		{
			name:     "StyleTextPrimary",
			style:    StyleTextPrimary,
			contains: []string{"color: #007BFF"},
		},
		{
			name:     "StyleTextSuccess",
			style:    StyleTextSuccess,
			contains: []string{"color: #28A745"},
		},
		{
			name:     "StyleTextDanger",
			style:    StyleTextDanger,
			contains: []string{"color: #DC3545"},
		},
		{
			name:     "StyleTextWarning",
			style:    StyleTextWarning,
			contains: []string{"color: #FFC107"},
		},
		{
			name:     "StyleBgLight",
			style:    StyleBgLight,
			contains: []string{"background-color: #F8F9FA"},
		},
		{
			name:     "StyleBgDark",
			style:    StyleBgDark,
			contains: []string{"background-color: #343A40"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.style == "" {
				t.Errorf("%s is empty", tt.name)
			}
			for _, substr := range tt.contains {
				if !strings.Contains(tt.style, substr) {
					t.Errorf("%s should contain %q, got: %s", tt.name, substr, tt.style)
				}
			}
		})
	}
}

func TestStylesAreNotEmpty(t *testing.T) {
	allStyles := []struct {
		name  string
		value string
	}{
		{"StyleHeading1", StyleHeading1},
		{"StyleHeading2", StyleHeading2},
		{"StyleHeading3", StyleHeading3},
		{"StyleParagraph", StyleParagraph},
		{"StyleSmall", StyleSmall},
		{"StyleButtonPrimary", StyleButtonPrimary},
		{"StyleButtonSecondary", StyleButtonSecondary},
		{"StyleButtonSuccess", StyleButtonSuccess},
		{"StyleButtonDanger", StyleButtonDanger},
		{"StyleButtonSmall", StyleButtonSmall},
		{"StyleContainer", StyleContainer},
		{"StyleSection", StyleSection},
		{"StyleDivider", StyleDivider},
		{"StyleCard", StyleCard},
		{"StyleAlertInfo", StyleAlertInfo},
		{"StyleAlertSuccess", StyleAlertSuccess},
		{"StyleAlertWarning", StyleAlertWarning},
		{"StyleAlertDanger", StyleAlertDanger},
		{"StyleListUnordered", StyleListUnordered},
		{"StyleListOrdered", StyleListOrdered},
		{"StyleListItem", StyleListItem},
		{"StyleTable", StyleTable},
		{"StyleTableHead", StyleTableHead},
		{"StyleTableCell", StyleTableCell},
		{"StyleTextCenter", StyleTextCenter},
		{"StyleTextRight", StyleTextRight},
		{"StyleTextMuted", StyleTextMuted},
		{"StyleTextPrimary", StyleTextPrimary},
		{"StyleTextSuccess", StyleTextSuccess},
		{"StyleTextDanger", StyleTextDanger},
		{"StyleTextWarning", StyleTextWarning},
		{"StyleBgLight", StyleBgLight},
		{"StyleBgDark", StyleBgDark},
	}

	for _, style := range allStyles {
		t.Run(style.name, func(t *testing.T) {
			if style.value == "" {
				t.Errorf("%s is empty", style.name)
			}
		})
	}
}

func TestStyleConstantsCount(t *testing.T) {
	expectedCount := 33
	allStyles := []string{
		StyleHeading1, StyleHeading2, StyleHeading3, StyleParagraph, StyleSmall,
		StyleButtonPrimary, StyleButtonSecondary, StyleButtonSuccess, StyleButtonDanger, StyleButtonSmall,
		StyleContainer, StyleSection, StyleDivider, StyleCard,
		StyleAlertInfo, StyleAlertSuccess, StyleAlertWarning, StyleAlertDanger,
		StyleListUnordered, StyleListOrdered, StyleListItem,
		StyleTable, StyleTableHead, StyleTableCell,
		StyleTextCenter, StyleTextRight, StyleTextMuted, StyleTextPrimary,
		StyleTextSuccess, StyleTextDanger, StyleTextWarning,
		StyleBgLight, StyleBgDark,
	}

	if len(allStyles) != expectedCount {
		t.Errorf("Expected %d style constants, got %d", expectedCount, len(allStyles))
	}
}
