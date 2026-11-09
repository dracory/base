package widgets

// RegistryOptions configures the widgets included in WidgetRegistry.
type RegistryOptions struct {
	// IsEnvironment reports whether the given environment name is the
	// currently active one. Used by the "visible" widget.
	IsEnvironment func(environment string) bool

	// WidgetURL builds the editor endpoint URL for a widget alias.
	// Used by the blockeditor widget.
	WidgetURL func(alias string, params map[string]string) string
}

// WidgetRegistry returns the generic widgets provided by this package.
//
// Feature widgets (blog, contact, testimonials, …) live with their packages;
// append them to the returned list before registering shortcodes.
//
// Parameters:
//   - options: dependencies the widgets need (all optional)
//
// Returns:
//   - []Widget - A list of all generic widgets
func WidgetRegistry(options RegistryOptions) []Widget {
	return []Widget{
		NewAuthenticatedWidget(),
		// NewContactFormWidget(),
		// NewTermsOfUseWidget(),
		NewUnauthenticatedWidget(),
		NewVisibleWidget(options.IsEnvironment),
		NewPrintWidget(),
		NewBlockeditorWidget(options.WidgetURL),
	}
}
