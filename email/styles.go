package email

// BrandColor allows projects to customize email colors to match their brand.
// These variables can be overridden in a project's init() function.
var (
	// ColorPrimary is the main brand color (default: blue)
	// Override example: email.ColorPrimary = "#6f42c1" // purple
	ColorPrimary = "#007BFF"

	// ColorSecondary is for secondary elements (default: gray)
	ColorSecondary = "#6C757D"

	// ColorSuccess is for positive/success states (default: green)
	ColorSuccess = "#28A745"

	// ColorDanger is for errors/danger states (default: red)
	ColorDanger = "#DC3545"

	// ColorWarning is for warnings (default: yellow/orange)
	ColorWarning = "#FFC107"

	// ColorInfo is for informational messages (default: teal/cyan)
	ColorInfo = "#17A2B8"

	// ColorLight is for light backgrounds (default: light gray)
	ColorLight = "#F8F9FA"

	// ColorDark is for dark elements (default: dark gray)
	ColorDark = "#343A40"
)

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
// These variables define inline CSS styles for button/link elements.
// They use BrandColor variables which can be customized per-project.
// Use these for call-to-action links in emails.
var (
	// StyleButtonPrimary is the main call-to-action button (uses ColorPrimary)
	StyleButtonPrimary = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: " + ColorPrimary + "; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid " + ColorPrimary + ";"

	// StyleButtonSecondary is a secondary action button (outlined, uses ColorPrimary)
	StyleButtonSecondary = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: " + ColorPrimary + "; background-color: transparent; text-align: center; text-decoration: none; border-radius: 6px; border: 2px solid " + ColorPrimary + ";"

	// StyleButtonSuccess is for positive actions (uses ColorSuccess)
	StyleButtonSuccess = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: " + ColorSuccess + "; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid " + ColorSuccess + ";"

	// StyleButtonDanger is for destructive actions (uses ColorDanger)
	StyleButtonDanger = "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: white; background-color: " + ColorDanger + "; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid " + ColorDanger + ";"

	// StyleButtonSmall is a smaller button variant (uses ColorPrimary)
	StyleButtonSmall = "display: inline-block; padding: 8px 16px; font-size: 14px; font-weight:600; color: white; background-color: " + ColorPrimary + "; text-align: center; text-decoration: none; border-radius: 4px; border: 1px solid " + ColorPrimary + ";"
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
	// StyleAlertInfo is for informational messages (uses ColorInfo)
	StyleAlertInfo = "padding: 12px 16px; background-color: #D1ECF1; border: 1px solid #BEE5EB; border-radius: 4px; color: #0C5460; margin: 10px 0px;"

	// StyleAlertSuccess is for success messages (uses ColorSuccess)
	StyleAlertSuccess = "padding: 12px 16px; background-color: #D4EDDA; border: 1px solid #C3E6CB; border-radius: 4px; color: #155724; margin: 10px 0px;"

	// StyleAlertWarning is for warning messages (uses ColorWarning)
	StyleAlertWarning = "padding: 12px 16px; background-color: #FFF3CD; border: 1px solid #FFEAA7; border-radius: 4px; color: #856404; margin: 10px 0px;"

	// StyleAlertDanger is for error/danger messages (uses ColorDanger)
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
// These variables define inline CSS utility styles for common formatting needs.
// They use BrandColor variables which can be customized per-project.
var (
	// StyleTextCenter centers text horizontally
	StyleTextCenter = "text-align: center;"

	// StyleTextRight aligns text to the right
	StyleTextRight = "text-align: right;"

	// StyleTextMuted is for muted/secondary text (gray)
	StyleTextMuted = "color: #6c757d;"

	// StyleTextPrimary is for primary brand color text (uses ColorPrimary)
	StyleTextPrimary = "color: " + ColorPrimary + ";"

	// StyleTextSuccess is for success state text (uses ColorSuccess)
	StyleTextSuccess = "color: " + ColorSuccess + ";"

	// StyleTextDanger is for danger/error state text (uses ColorDanger)
	StyleTextDanger = "color: " + ColorDanger + ";"

	// StyleTextWarning is for warning state text (uses ColorWarning)
	StyleTextWarning = "color: " + ColorWarning + ";"

	// StyleBgLight is for light background color (uses ColorLight)
	StyleBgLight = "background-color: " + ColorLight + ";"

	// StyleBgDark is for dark background color (uses ColorDark)
	StyleBgDark = "background-color: " + ColorDark + ";"
)

// StyleBuilder provides functions to generate custom styles with specific colors.
// These are useful when you need styles with colors different from the BrandColor variables.

// ButtonStyle generates a custom button style with the specified colors.
// Use this when you need a button with a specific brand color not covered by BrandColor variables.
//
// Example:
//
//	customButton := email.ButtonStyle("#6f42c1", "#5a32a3", "white")
//	// Creates a purple button with darker purple border and white text
func ButtonStyle(backgroundColor, borderColor, textColor string) string {
	return "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: " + textColor + "; background-color: " + backgroundColor + "; text-align: center; text-decoration: none; border-radius: 6px; border: 1px solid " + borderColor + ";"
}

// ButtonStyleSecondary generates a custom outlined button style.
// Use this for secondary actions with custom brand colors.
//
// Example:
//
//	customSecondaryBtn := email.ButtonStyleSecondary("#6f42c1")
//	// Creates an outlined purple button
func ButtonStyleSecondary(color string) string {
	return "display: inline-block; padding: 12px 24px; font-size: 16px; font-weight:600; color: " + color + "; background-color: transparent; text-align: center; text-decoration: none; border-radius: 6px; border: 2px solid " + color + ";"
}

// AlertStyle generates a custom alert style with specified colors.
// Use this when you need alerts with custom color schemes.
//
// Example:
//
//	customAlert := email.AlertStyle("#6f42c1", "#e9d5ff", "#3b0764")
//	// Creates a purple-themed alert
func AlertStyle(backgroundColor, borderColor, textColor string) string {
	return "padding: 12px 16px; background-color: " + backgroundColor + "; border: 1px solid " + borderColor + "; border-radius: 4px; color: " + textColor + "; margin: 10px 0px;"
}

// SetBrandColors allows setting all brand colors at once.
// Call this in your project's init() function to customize email colors.
//
// Example:
//
//	func init() {
//	    email.SetBrandColors(
//	        "#6f42c1", // Primary (purple)
//	        "#6C757D", // Secondary (gray)
//	        "#28A745", // Success (green)
//	        "#DC3545", // Danger (red)
//	        "#FFC107", // Warning (yellow)
//	        "#17A2B8", // Info (teal)
//	        "#F8F9FA", // Light
//	        "#343A40", // Dark
//	    )
//	}
func SetBrandColors(primary, secondary, success, danger, warning, info, light, dark string) {
	ColorPrimary = primary
	ColorSecondary = secondary
	ColorSuccess = success
	ColorDanger = danger
	ColorWarning = warning
	ColorInfo = info
	ColorLight = light
	ColorDark = dark
}
