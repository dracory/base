package tz

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

// UTCDateToTz converts a UTC date string to the specified timezone.
// It accepts date strings in "YYYY-MM-DD" format.
//
// Example:
//
//	dateInNewYork, err := UTCDateToTz("2022-01-01", "America/New_York")
//
// Parameters:
//   - utcDateString: a string representing the UTC date in "YYYY-MM-DD" format.
//   - timezone: a string representing the timezone (e.g., "America/New_York").
//
// Returns:
//   - a string representing the converted date in "YYYY-MM-DD" format, or an empty string and an error if parsing fails.
func UTCDateToTz(utcDateString string, timezone string) (string, error) {
	if utcDateString == "" {
		return "", fmt.Errorf("utcDateString cannot be empty")
	}

	if timezone == "" {
		return "", fmt.Errorf("timezone cannot be empty")
	}

	parsedDate := carbon.Parse(utcDateString, carbon.UTC)

	if parsedDate.Error != nil {
		return "", fmt.Errorf("invalid date format: %w", parsedDate.Error)
	}

	dateInTz := parsedDate.SetTimezone(timezone)
	return dateInTz.ToDateString(), parsedDate.Error
}
