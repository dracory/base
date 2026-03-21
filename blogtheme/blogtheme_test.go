package blogtheme

import (
	"testing"

	"github.com/dracory/ui"
)

func TestNew(t *testing.T) {
	// Create test blocks
	blocks := []ui.BlockInterface{
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("paragraph")
			block.SetParameter("content", "Test paragraph content")
			return block
		}(),
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("heading")
			block.SetParameter("content", "Test Heading")
			block.SetParameter("level", "1")
			return block
		}(),
	}

	// Convert to JSON
	blocksJSON, err := ui.MarshalBlocksToJson(blocks)
	if err != nil {
		t.Fatalf("Failed to marshal blocks to JSON: %v", err)
	}

	theme, err := New(blocksJSON)
	if err != nil {
		t.Fatalf("Expected theme to be created, got error: %v", err)
	}

	if theme == nil {
		t.Fatal("Expected theme to be created, got nil")
	}
}

func TestThemeToHtml(t *testing.T) {
	// Create test blocks
	blocks := []ui.BlockInterface{
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("paragraph")
			block.SetParameter("content", "Test paragraph content")
			return block
		}(),
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("heading")
			block.SetParameter("content", "Test Heading")
			block.SetParameter("level", "1")
			return block
		}(),
	}

	// Convert to JSON
	blocksJSON, err := ui.MarshalBlocksToJson(blocks)
	if err != nil {
		t.Fatalf("Failed to marshal blocks to JSON: %v", err)
	}

	theme, err := New(blocksJSON)
	if err != nil {
		t.Fatalf("Expected theme to be created, got error: %v", err)
	}

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

func TestThemeStyle(t *testing.T) {
	// Create test blocks
	blocks := []ui.BlockInterface{
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("paragraph")
			block.SetParameter("content", "Test content")
			return block
		}(),
	}

	// Convert to JSON
	blocksJSON, err := ui.MarshalBlocksToJson(blocks)
	if err != nil {
		t.Fatalf("Failed to marshal blocks to JSON: %v", err)
	}

	theme, err := New(blocksJSON)
	if err != nil {
		t.Fatalf("Expected theme to be created, got error: %v", err)
	}

	style := theme.Style()

	if style == "" {
		t.Fatal("Expected style output, got empty string")
	}

	// Check for expected CSS classes
	if !contains(style, ".BlogTitle") {
		t.Error("Expected .BlogTitle class in style output")
	}

	if !contains(style, ".BlogContent") {
		t.Error("Expected .BlogContent class in style output")
	}

	if !contains(style, "h1") {
		t.Error("Expected h1 styles in style output")
	}
}

func TestSupportedBlockTypes(t *testing.T) {
	// Create test blocks for each supported type
	blockTypes := []string{
		"heading",
		"hyperlink",
		"image",
		"paragraph",
		"raw",
		"unordered_list",
		"list_item",
		"ordered_list",
	}

	for _, blockType := range blockTypes {
		blocks := []ui.BlockInterface{
			func() ui.BlockInterface {
				block := ui.NewBlock()
				block.SetType(blockType)
				block.SetParameter("content", "Test content")
				return block
			}(),
		}

		// Convert to JSON
		blocksJSON, err := ui.MarshalBlocksToJson(blocks)
		if err != nil {
			t.Fatalf("Failed to marshal blocks to JSON for block type %s: %v", blockType, err)
		}

		theme, err := New(blocksJSON)
		if err != nil {
			t.Fatalf("Expected theme to be created for block type %s, got error: %v", blockType, err)
		}

		html := theme.ToHtml()
		if html == "" {
			t.Errorf("Expected HTML output for block type %s, got empty string", blockType)
		}
	}
}

func TestUnsupportedBlockType(t *testing.T) {
	// Create test block with unsupported type
	blocks := []ui.BlockInterface{
		func() ui.BlockInterface {
			block := ui.NewBlock()
			block.SetType("unsupported_type")
			block.SetParameter("content", "Test content")
			return block
		}(),
	}

	// Convert to JSON
	blocksJSON, err := ui.MarshalBlocksToJson(blocks)
	if err != nil {
		t.Fatalf("Failed to marshal blocks to JSON: %v", err)
	}

	theme, err := New(blocksJSON)
	if err != nil {
		t.Fatalf("Expected theme to be created, got error: %v", err)
	}

	html := theme.ToHtml()

	// Should contain warning message for unsupported block
	if !contains(html, "alert alert-warning") {
		t.Error("Expected warning alert for unsupported block type")
	}

	if !contains(html, "unsupported_type") {
		t.Error("Expected unsupported block type in warning message")
	}
}

func TestInvalidJSON(t *testing.T) {
	invalidJSON := "{ invalid json }"

	theme, err := New(invalidJSON)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	if theme != nil {
		t.Error("Expected theme to be nil for invalid JSON, got theme")
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
