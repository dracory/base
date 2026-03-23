package email

import "github.com/dracory/hb"

// Heading creates a styled heading (level 1-3).
// If level is outside 1-3, it defaults to Heading1.
func Heading(text string, level int) *hb.Tag {
	var h *hb.Tag
	switch level {
	case 1:
		h = hb.Heading1().Style(StyleHeading1)
	case 2:
		h = hb.Heading2().Style(StyleHeading2)
	case 3:
		h = hb.Heading3().Style(StyleHeading3)
	default:
		h = hb.Heading1().Style(StyleHeading1)
	}
	return h.HTML(text)
}

// Paragraph creates a styled paragraph.
func Paragraph(text string) *hb.Tag {
	return hb.Paragraph().Style(StyleParagraph).HTML(text)
}

// Small creates a styled small paragraph.
func Small(text string) *hb.Tag {
	return hb.Paragraph().Style(StyleSmall).HTML(text)
}

// Button creates a styled primary brand button.
func Button(text, url string) *hb.Tag {
	return hb.Hyperlink().
		Href(url).
		Style(StyleButtonPrimary).
		HTML(text)
}

// ButtonSecondary creates a styled secondary (outlined) button.
func ButtonSecondary(text, url string) *hb.Tag {
	return hb.Hyperlink().
		Href(url).
		Style(StyleButtonSecondary).
		HTML(text)
}

// ButtonSuccess creates a styled success button.
func ButtonSuccess(text, url string) *hb.Tag {
	return hb.Hyperlink().
		Href(url).
		Style(StyleButtonSuccess).
		HTML(text)
}

// ButtonDanger creates a styled danger button.
func ButtonDanger(text, url string) *hb.Tag {
	return hb.Hyperlink().
		Href(url).
		Style(StyleButtonDanger).
		HTML(text)
}

// ButtonSmall creates a smaller primary button.
func ButtonSmall(text, url string) *hb.Tag {
	return hb.Hyperlink().
		Href(url).
		Style(StyleButtonSmall).
		HTML(text)
}

// Alert creates a styled alert box.
// Supported variants: info, success, warning, danger.
func Alert(text, variant string) *hb.Tag {
	alert := hb.Div().HTML(text)
	switch variant {
	case "success":
		alert.Style(StyleAlertSuccess)
	case "warning":
		alert.Style(StyleAlertWarning)
	case "danger":
		alert.Style(StyleAlertDanger)
	default:
		alert.Style(StyleAlertInfo)
	}
	return alert
}

// Card creates a styled card container.
func Card(children ...hb.TagInterface) *hb.Tag {
	return hb.Div().Style(StyleCard).Children(children)
}

// Section creates a styled section with a light background.
func Section(children ...hb.TagInterface) *hb.Tag {
	return hb.Div().Style(StyleSection).Children(children)
}

// Container creates a styled container with maximum width.
func Container(children ...hb.TagInterface) *hb.Tag {
	return hb.Div().Style(StyleContainer).Children(children)
}

// Divider creates a horizontal rule/separator.
func Divider() *hb.Tag {
	return hb.Div().Style(StyleDivider)
}

// Table creates a styled data table from headers and rows.
func Table(headers []string, rows [][]string) *hb.Tag {
	table := hb.Table().Style(StyleTable)

	// Create table header
	trHead := hb.TR()
	for _, h := range headers {
		trHead.Child(hb.TH().Style(StyleTableHead).HTML(h))
	}
	table.Child(trHead)

	// Create table rows
	for _, row := range rows {
		trRow := hb.TR()
		for _, cell := range row {
			trRow.Child(hb.TD().Style(StyleTableCell).HTML(cell))
		}
		table.Child(trRow)
	}

	return table
}
