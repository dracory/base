package tz_test

import (
	"testing"

	"github.com/dracory/base/tz"
)

func TestUTCDateToTz(t *testing.T) {
	testCases := []struct {
		name          string
		utcDateString string
		timezone      string
		expected      string
		expectError   bool
	}{
		{
			name:          "Valid Date and Timezone",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/New_York",
			expected:      "2023-10-26", // 2023-10-26 20:00:00 America/New_York
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - Different Date",
			utcDateString: "2024-01-15", // 2024-01-15 00:00:00 UTC
			timezone:      "Europe/London",
			expected:      "2024-01-15", // 2024-01-15 00:00:00 Europe/London
			expectError:   false,
		},
		{
			name:          "Valid Date Format - Though Not Very Common",
			utcDateString: "2023/10/27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/New_York",
			expected:      "2023-10-26", // 2023-10-26 20:00:00 America/New_York
			expectError:   false,
		},
		{
			name:          "Invalid Timezone",
			utcDateString: "2023-10-27",
			timezone:      "Invalid/Timezone",
			expected:      "",
			expectError:   true,
		},
		{
			name:          "Empty Date String",
			utcDateString: "",
			timezone:      "America/New_York",
			expected:      "",
			expectError:   true,
		},
		{
			name:          "Empty Timezone",
			utcDateString: "2023-10-27",
			timezone:      "",
			expected:      "",
			expectError:   true,
		},
		{
			name:          "Valid Date and Timezone - Asia/Tokyo",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "Asia/Tokyo",
			expected:      "2023-10-27", // 2023-10-27 09:00:00 Asia/Tokyo
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - UTC",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "UTC",
			expected:      "2023-10-27", // 2023-10-27 00:00:00 UTC
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - Europe/Berlin",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "Europe/Berlin",
			expected:      "2023-10-27", // 2023-10-27 02:00:00 Europe/Berlin
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Los_Angeles",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Los_Angeles",
			expected:      "2023-10-26", // 2023-10-26 17:00:00 America/Los_Angeles
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Chicago",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Chicago",
			expected:      "2023-10-26", // 2023-10-26 19:00:00 America/Chicago
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Denver",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Denver",
			expected:      "2023-10-26", // 2023-10-26 18:00:00 America/Denver
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Phoenix",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Phoenix",
			expected:      "2023-10-26", // 2023-10-26 17:00:00 America/Phoenix
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Anchorage",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Anchorage",
			expected:      "2023-10-26", // 2023-10-26 16:00:00 America/Anchorage
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - America/Adak",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "America/Adak",
			expected:      "2023-10-26", // 2023-10-26 15:00:00 America/Adak
			expectError:   false,
		},
		{
			name:          "Valid Date and Timezone - Pacific/Honolulu",
			utcDateString: "2023-10-27", // 2023-10-27 00:00:00 UTC
			timezone:      "Pacific/Honolulu",
			expected:      "2023-10-26", // 2023-10-26 14:00:00 Pacific/Honolulu
			expectError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tz.UTCDateToTz(tc.utcDateString, tc.timezone)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error for %s, but got none", tc.name)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s: %v", tc.name, err)
				}
			}

			if result != tc.expected {
				t.Errorf("UTCDateToTz(%q, %q) = %q, want %q", tc.utcDateString, tc.timezone, result, tc.expected)
			}
		})
	}
}
