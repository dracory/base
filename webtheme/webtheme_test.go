package webtheme

import (
	"testing"

	"github.com/dracory/ui"
)

func TestNew(t *testing.T) {
	block := ui.NewBlock()
	block.SetType("paragraph")
	block.SetParameter("content", "Test paragraph content")

	blocks := []ui.BlockInterface{block}

	theme := New(blocks)
	if theme == nil {
		t.Fatal("Expected theme to be created, got nil")
	}
}

func TestThemeToHtml(t *testing.T) {
	paragraphBlock := ui.NewBlock()
	paragraphBlock.SetType("paragraph")
	paragraphBlock.SetParameter("content", "Test paragraph content")
	paragraphBlock.SetParameter("status", "published") // Add status to make it render

	headingBlock := ui.NewBlock()
	headingBlock.SetType("heading")
	headingBlock.SetParameter("level", "1")
	headingBlock.SetParameter("content", "Test Heading")
	headingBlock.SetParameter("status", "published") // Add status to make it render

	blocks := []ui.BlockInterface{paragraphBlock, headingBlock}

	theme := New(blocks)
	html := theme.ToHtml()

	if html == "" {
		t.Fatal("Expected HTML output, got empty string")
	}

	// Basic checks that the HTML contains expected elements
	if !contains(html, "<p>") || !contains(html, "</p>") {
		t.Errorf("Expected paragraph tags in HTML output. Got: %s", html)
	}

	if !contains(html, "<h1") || !contains(html, "</h1>") {
		t.Errorf("Expected h1 tags in HTML output. Got: %s", html)
	}
}

func TestBlockEditorDefinitions(t *testing.T) {
	definitions := BlockEditorDefinitions()

	if len(definitions) == 0 {
		t.Fatal("Expected block editor definitions, got empty slice")
	}

	// Check that common block types are defined
	expectedTypes := []string{
		TYPE_PARAGRAPH,
		TYPE_HEADING,
		TYPE_IMAGE,
		TYPE_CONTAINER,
		TYPE_ROW,
		TYPE_COLUMN,
	}

	typeMap := make(map[string]bool)
	for _, def := range definitions {
		typeMap[def.Type] = true
	}

	for _, expectedType := range expectedTypes {
		if !typeMap[expectedType] {
			t.Errorf("Expected block type %s to be defined", expectedType)
		}
	}
}

func TestBootstrapIconList(t *testing.T) {
	icons := bootstrapIconList()

	if len(icons) == 0 {
		t.Fatal("Expected bootstrap icon list, got empty slice")
	}

	// Check that it contains some common icons (use actual icons from the list)
	iconMap := make(map[string]bool)
	for _, icon := range icons {
		iconMap[icon.Icon] = true
	}

	// Use some icons that are actually in the bootstrap list
	commonIcons := []string{
		"bi bi-0-circle", // This is in the list
		"bi bi-1-circle", // This is in the list
		"bi bi-2-circle", // This is in the list
		"bi bi-3-circle", // This is in the list
	}

	for _, expectedIcon := range commonIcons {
		if !iconMap[expectedIcon] {
			// Let's print some actual icons to debug
			if len(icons) > 0 {
				t.Logf("First few available icons: %s, %s, %s, %s",
					icons[0].Icon, icons[1].Icon, icons[2].Icon, icons[3].Icon)
			}
			t.Errorf("Expected icon %s to be in bootstrap icon list", expectedIcon)
		}
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr) >= 0
}

func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
