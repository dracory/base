package tz

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

// UTCDatetimeToTz converts a UTC datetime string to the specified timezone.
// It accepts datetime strings in "YYYY-MM-DD HH:MM" format.
//
// Example:
//
//	datetimeInNewYork, err := UTCDatetimeToTz("2022-01-01 10:00", "America/New_York")
//
// Parameters:
//   - utcDatetimeString: a string representing the UTC datetime in "YYYY-MM-DD HH:MM" format.
//   - timezone: a string representing the timezone (e.g., "America/New_York").
//
// Returns:
//   - a string representing the converted datetime in "YYYY-MM-DD HH:MM" format, or an empty string and an error if parsing fails.
func UTCDatetimeToTz(utcDatetimeString string, timezone string) (string, error) {
	if utcDatetimeString == "" {
		return "", fmt.Errorf("utcDatetimeString cannot be empty")
	}

	if timezone == "" {
		return "", fmt.Errorf("timezone cannot be empty")
	}

	parsedDatetime := carbon.Parse(utcDatetimeString, carbon.UTC)

	if parsedDatetime.Error != nil {
		return "", fmt.Errorf("invalid datetime format: %w", parsedDatetime.Error)
	}

	datetimeInTz := parsedDatetime.SetTimezone(timezone)

	return datetimeInTz.ToDateTimeString(), datetimeInTz.Error
}
