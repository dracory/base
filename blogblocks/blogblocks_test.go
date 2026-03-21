package blogblocks

import (
	"testing"

	"github.com/dracory/blockeditor"
	"github.com/dracory/form"
	"github.com/dracory/ui"
)

func TestBlockEditorDefinitions(t *testing.T) {
	definitions := BlockEditorDefinitions()

	if len(definitions) == 0 {
		t.Fatal("Expected block editor definitions, got empty slice")
	}

	// Check that expected block types are defined
	expectedTypes := []string{
		"heading",
		"paragraph",
		"image",
		"hyperlink",
		"unordered_list",
		"ordered_list",
		"list_item",
	}

	typeMap := make(map[string]bool)
	for _, def := range definitions {
		typeMap[def.Type] = true

		// Verify each definition has required properties
		if def.Icon == nil {
			t.Errorf("Block definition %s should have an icon", def.Type)
		}

		if def.Type == "" {
			t.Error("Block definition should have a type")
		}
	}

	for _, expectedType := range expectedTypes {
		if !typeMap[expectedType] {
			t.Errorf("Expected block type %s to be defined", expectedType)
		}
	}
}

func TestHeadingBlockDefinition(t *testing.T) {
	definitions := BlockEditorDefinitions()

	var headingDef blockeditor.BlockDefinition
	found := false

	for _, def := range definitions {
		if def.Type == "heading" {
			headingDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected heading block definition to be found")
	}

	// Test heading block rendering
	testBlock := ui.NewBlock()
	testBlock.SetType("heading")
	testBlock.SetParameter("level", "2")
	testBlock.SetParameter("content", "Test Heading")

	if headingDef.ToTag == nil {
		t.Fatal("Expected heading block to have ToTag function")
	}

	result := headingDef.ToTag(testBlock)
	if result == nil {
		t.Fatal("Expected ToTag to return a result")
	}
}

func TestParagraphBlockDefinition(t *testing.T) {
	definitions := BlockEditorDefinitions()

	var paragraphDef blockeditor.BlockDefinition
	found := false

	for _, def := range definitions {
		if def.Type == "paragraph" {
			paragraphDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected paragraph block definition to be found")
	}

	// Test paragraph block rendering
	testBlock := ui.NewBlock()
	testBlock.SetType("paragraph")
	testBlock.SetParameter("content", "Test paragraph content")

	if paragraphDef.ToTag == nil {
		t.Fatal("Expected paragraph block to have ToTag function")
	}

	result := paragraphDef.ToTag(testBlock)
	if result == nil {
		t.Fatal("Expected ToTag to return a result")
	}
}

func TestImageBlockDefinition(t *testing.T) {
	definitions := BlockEditorDefinitions()

	var imageDef blockeditor.BlockDefinition
	found := false

	for _, def := range definitions {
		if def.Type == "image" {
			imageDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected image block definition to be found")
	}

	// Test image block rendering with URL
	testBlock := ui.NewBlock()
	testBlock.SetType("image")
	testBlock.SetParameter("image_url", "https://example.com/image.jpg")

	if imageDef.ToTag == nil {
		t.Fatal("Expected image block to have ToTag function")
	}

	result := imageDef.ToTag(testBlock)
	if result == nil {
		t.Fatal("Expected ToTag to return a result")
	}

	// Test image block rendering without URL (should use placeholder)
	testBlockEmpty := ui.NewBlock()
	testBlockEmpty.SetType("image")
	// Don't set image_url parameter

	resultEmpty := imageDef.ToTag(testBlockEmpty)
	if resultEmpty == nil {
		t.Fatal("Expected ToTag to return a result even without URL")
	}
}

func TestHyperlinkBlockDefinition(t *testing.T) {
	definitions := BlockEditorDefinitions()

	var hyperlinkDef blockeditor.BlockDefinition
	found := false

	for _, def := range definitions {
		if def.Type == "hyperlink" {
			hyperlinkDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected hyperlink block definition to be found")
	}

	// Test hyperlink block rendering
	testBlock := ui.NewBlock()
	testBlock.SetType("hyperlink")
	testBlock.SetParameter("url", "https://example.com")
	testBlock.SetParameter("content", "Click here")
	testBlock.SetParameter("target", "_blank")

	if hyperlinkDef.ToTag == nil {
		t.Fatal("Expected hyperlink block to have ToTag function")
	}

	result := hyperlinkDef.ToTag(testBlock)
	if result == nil {
		t.Fatal("Expected ToTag to return a result")
	}
}

func TestListBlockDefinitions(t *testing.T) {
	definitions := BlockEditorDefinitions()

	// Test unordered list
	var ulDef blockeditor.BlockDefinition
	found := false

	for _, def := range definitions {
		if def.Type == "unordered_list" {
			ulDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected unordered_list block definition to be found")
	}

	if !ulDef.AllowChildren {
		t.Error("Expected unordered list to allow children")
	}

	if len(ulDef.AllowedChildTypes) == 0 {
		t.Error("Expected unordered list to have allowed child types")
	}

	// Test ordered list
	var olDef blockeditor.BlockDefinition
	found = false

	for _, def := range definitions {
		if def.Type == "ordered_list" {
			olDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected ordered_list block definition to be found")
	}

	if !olDef.AllowChildren {
		t.Error("Expected ordered list to allow children")
	}

	// Test list item
	var liDef blockeditor.BlockDefinition
	found = false

	for _, def := range definitions {
		if def.Type == "list_item" {
			liDef = def
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected list_item block definition to be found")
	}

	// Test list item rendering
	testBlock := ui.NewBlock()
	testBlock.SetType("list_item")
	testBlock.SetParameter("content", "List item content")

	if liDef.ToTag == nil {
		t.Fatal("Expected list item block to have ToTag function")
	}

	result := liDef.ToTag(testBlock)
	if result == nil {
		t.Fatal("Expected ToTag to return a result")
	}
}

func TestBlockDefinitionFields(t *testing.T) {
	definitions := BlockEditorDefinitions()

	// Test that definitions have proper field structure
	for _, def := range definitions {
		// All definitions should have fields (even if empty)
		if def.Fields == nil {
			t.Errorf("Block definition %s should have fields array", def.Type)
		}

		// Check field types
		for _, field := range def.Fields {
			if field == nil {
				t.Errorf("Block definition %s has nil field", def.Type)
				continue
			}

			// Verify field has required properties
			if field.GetName() == "" && field.GetType() != form.FORM_FIELD_TYPE_RAW {
				t.Errorf("Field in block definition %s should have a name", def.Type)
			}

			if field.GetType() == "" {
				t.Errorf("Field in block definition %s should have a type", def.Type)
			}
		}
	}
}

func TestBlockDefinitionIcons(t *testing.T) {
	definitions := BlockEditorDefinitions()

	// Test that all definitions have icons
	for _, def := range definitions {
		if def.Icon == nil {
			t.Errorf("Block definition %s should have an icon", def.Type)
		}
	}
}
