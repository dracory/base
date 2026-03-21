package blogtheme

import (
	"github.com/dracory/blockeditor"
	"github.com/dracory/hb"
	"github.com/dracory/ui"
	"github.com/samber/lo"
)

type theme struct {
	blocks          []ui.BlockInterface
	tree            *blockeditor.FlatTree
	availableBlocks []struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}
}

func New(blocksJSON string) (*theme, error) {
	blocks, err := ui.UnmarshalJsonToBlocks(blocksJSON)

	if err != nil {
		return nil, err
	}

	tree := blockeditor.NewFlatTree(blocks)

	t := &theme{
		blocks: blocks,
		tree:   tree,
	}

	t.addBlockRenderer("heading", t.headingToHtml)
	t.addBlockRenderer("hyperlink", t.hyperlinkToHtml)
	t.addBlockRenderer("image", t.imageToHtml)
	t.addBlockRenderer("paragraph", t.paragraphToHtml)
	t.addBlockRenderer("raw", t.rawToHtml)
	t.addBlockRenderer("unordered_list", t.ulToHtml)
	t.addBlockRenderer("list_item", t.liToHtml)
	t.addBlockRenderer("ordered_list", t.olToHtml)

	return t, nil
}

func (t *theme) Style() string {
	style := `
.BlogTitle {
	font-family: Roboto, sans-serif;
}
.BlogContent {
	font-family: Roboto, sans-serif;
}
h1 { 
	margin-bottom: 20px;
	font-size: 48px;
}
h2 { 
	margin-bottom: 20px;
	font-size: 36px;
}
h3 {
	margin-bottom: 20px;
	font-size: 24px;
}
h4 {
	margin-bottom: 20px;
	font-size: 18px;
}
h5 {
	margin-bottom: 20px;
	font-size: 16px;
}
h6 {
	margin-bottom: 20px;
	font-size: 14px;
}
	`
	return style
}

func (t *theme) addBlockRenderer(blockType string, toTag func(block ui.BlockInterface) *hb.Tag) {
	t.availableBlocks = append(t.availableBlocks, struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}{
		Type:  blockType,
		ToTag: toTag,
	})
}

func (t *theme) ToHtml() string {
	wrap := hb.Wrap()

	for _, block := range t.blocks {
		childrenTags := lo.Map(block.Children(), func(block ui.BlockInterface, _ int) hb.TagInterface {
			return t.blockToTag(block)
		})
		blockTag := t.blockToTag(block).Children(childrenTags)
		wrap.Child(blockTag)
	}

	return wrap.ToHTML()
}

func (t *theme) isSupportedBlock(block ui.BlockInterface) bool {
	for _, availableBlock := range t.availableBlocks {
		if block.Type() == availableBlock.Type {
			return true
		}
	}

	return false
}

func (t *theme) blockToTag(block ui.BlockInterface) *hb.Tag {
	if !t.isSupportedBlock(block) {
		return hb.Div().
			Class("alert alert-warning").
			Text("Block " + block.Type() + " renderer does not exist")
	}

	renderer, found := lo.Find(t.availableBlocks, func(availableBlock struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}) bool {
		return availableBlock.Type == block.Type()
	})

	if !found {
		return hb.Div().
			Class("alert alert-warning").
			Text("Block " + block.Type() + " renderer does not exist")
	}

	return renderer.ToTag(block)
}
