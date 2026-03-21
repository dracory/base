package types

// FlashMessage represents a flash message for web applications
type FlashMessage struct {
	Type    string // Message type (e.g., "success", "error", "info", "warning")
	Message string // The message content
	Url     string // Optional URL for redirect after message display
	Time    string // Timestamp when the message was created
}
