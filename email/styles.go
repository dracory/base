package email

// Typography Styles
//
// These constants define inline CSS styles for text elements in HTML emails.
// Email clients require inline styles for proper rendering.
const (
	// StyleHeading1 is the primary heading style (24px, bold)
	StyleHeading1 = "margin:0px;padding:10px 0px;text-align:left;font-size:24px;font-weight:600;color:#333333;"

	// StyleHeading2 is the secondary heading style (20px, bold)
	StyleHeading2 = "margin:0px;padding:8px 0px;text-align:left;font-size:20px;font-weight:600;color:#333333;"

	// StyleHeading3 is the tertiary heading style (18px, bold)
	StyleHeading3 = "margin:0px;padding:6px 0px;text-align:left;font-size:18px;font-weight:600;color:#333333;"

	// StyleParagraph is the standard paragraph style (16px, normal line height)
	StyleParagraph = "margin:0px;padding:10px 0px;text-align:left;font-size:16px;line-height:1.6;color:#333333;"

	// StyleSmall is for smaller text like disclaimers or footnotes (14px)
	StyleSmall = "margin:0px;padding:5px 0px;text-align:left;font-size:14px;color:#666666;"
)

// Button Styles
//
// These constants define inline CSS styles for button/link elements.
// Use these for call-to-action links in emails.
const (
	// StyleButtonPrimary is the main call-to-action button (blue background)
	StyleButtonPrimary = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: #007BFF; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid #007BFF;"

	// StyleButtonSecondary is a secondary action button (outlined, no background)
	StyleButtonSecondary = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: #007BFF; background-color: transparent; text-align: center; text-decoration: none; border-radius: 6px; border: 2px solid #007BFF;"

	// StyleButtonSuccess is for positive actions (green background)
	StyleButtonSuccess = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: #28A745; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid #28A745;"

	// StyleButtonDanger is for destructive actions (red background)
	StyleButtonDanger = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: #DC3545; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid #DC3545;"

	// StyleButtonSmall is a smaller button variant (14px, less padding)
	StyleButtonSmall = "display: inline-block; padding: 8px 16px; font-size: 14px; font-weight:600; color: white; background-color: #007BFF; text-align: center; text-decoration: none; border-radius: 4px; border: 1px solid #007BFF;"
)

// Layout Styles
//
// These constants define inline CSS styles for layout containers and sections.
const (
	// StyleContainer is a centered container with max-width (600px)
	StyleContainer = "max-width: 600px; margin: 0 auto; padding: 20px; background-color: #ffffff;"

	// StyleSection is a content section with light background
	StyleSection = "margin: 20px 0px; padding: 15px; background-color: #f8f9fa; border-radius: 6px;"

	// StyleDivider is a horizontal rule/separator
	StyleDivider = "height: 1px; background-color: #dee2e6; margin: 20px 0px; border: none;"

	// StyleCard is a bordered card container
	StyleCard = "padding: 20px; background-color: #ffffff; border: 1px solid #dee2e6; border-radius: 8px; margin: 10px 0px;"
)

// Alert Styles
//
// These constants define inline CSS styles for alert/notification boxes.
const (
	// StyleAlertInfo is for informational messages (blue)
	StyleAlertInfo = "padding: 12px 16px; background-color: #D1ECF1; border: 1px solid #BEE5EB; border-radius: 4px; color: #0C5460; margin: 10px 0px;"

	// StyleAlertSuccess is for success messages (green)
	StyleAlertSuccess = "padding: 12px 16px; background-color: #D4EDDA; border: 1px solid #C3E6CB; border-radius: 4px; color: #155724; margin: 10px 0px;"

	// StyleAlertWarning is for warning messages (yellow)
	StyleAlertWarning = "padding: 12px 16px; background-color: #FFF3CD; border: 1px solid #FFEAA7; border-radius: 4px; color: #856404; margin: 10px 0px;"

	// StyleAlertDanger is for error/danger messages (red)
	StyleAlertDanger = "padding: 12px 16px; background-color: #F8D7DA; border: 1px solid #F5C6CB; border-radius: 4px; color: #721C24; margin: 10px 0px;"
)

// List Styles
//
// These constants define inline CSS styles for lists.
const (
	// StyleListUnordered is for unordered (bulleted) lists
	StyleListUnordered = "margin: 10px 0px; padding-left: 20px; color: #333333;"

	// StyleListOrdered is for ordered (numbered) lists
	StyleListOrdered = "margin: 10px 0px; padding-left: 20px; color: #333333;"

	// StyleListItem is for individual list items
	StyleListItem = "margin: 5px 0px; line-height: 1.5;"
)

// Table Styles
//
// These constants define inline CSS styles for tables.
const (
	// StyleTable is the base table style
	StyleTable = "width: 100%; border-collapse: collapse; margin: 15px 0px;"

	// StyleTableHead is for table header cells
	StyleTableHead = "background-color: #f8f9fa; border: 1px solid #dee2e6; padding: 12px; text-align: left; font-weight: 600;"

	// StyleTableCell is for table data cells
	StyleTableCell = "border: 1px solid #dee2e6; padding: 12px; text-align: left;"
)

// Utility Styles
//
// These constants define inline CSS utility styles for common formatting needs.
const (
	// StyleTextCenter centers text horizontally
	StyleTextCenter = "text-align: center;"

	// StyleTextRight aligns text to the right
	StyleTextRight = "text-align: right;"

	// StyleTextMuted is for muted/secondary text (gray)
	StyleTextMuted = "color: #6c757d;"

	// StyleTextPrimary is for primary brand color text (blue)
	StyleTextPrimary = "color: #007BFF;"

	// StyleTextSuccess is for success state text (green)
	StyleTextSuccess = "color: #28A745;"

	// StyleTextDanger is for danger/error state text (red)
	StyleTextDanger = "color: #DC3545;"

	// StyleTextWarning is for warning state text (yellow)
	StyleTextWarning = "color: #FFC107;"

	// StyleBgLight is for light background color
	StyleBgLight = "background-color: #f8f9fa;"

	// StyleBgDark is for dark background color
	StyleBgDark = "background-color: #343a40;"
)
