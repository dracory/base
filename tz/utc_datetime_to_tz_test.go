package tz_test

import (
	"testing"

	"github.com/dracory/base/tz"
)

func TestUTCDatetimeToTz(t *testing.T) {
	testCases := []struct {
		name              string
		utcDatetimeString string
		timezone          string
		expected          string
		expectError       bool
	}{
		{
			name:              "Valid Datetime and Timezone",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/New_York",
			expected:          "2023-10-27 06:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - Different Datetime",
			utcDatetimeString: "2024-01-15 15:30",
			timezone:          "Europe/London",
			expected:          "2024-01-15 15:30:00",
			expectError:       false,
		},
		{
			name:              "Invalid Datetime Format",
			utcDatetimeString: "2023/10/27 10:00",
			timezone:          "America/New_York",
			expected:          "2023-10-27 06:00:00",
			expectError:       false,
		},
		{
			name:              "Invalid Timezone",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "Invalid/Timezone",
			expected:          "",
			expectError:       true,
		},
		{
			name:              "Empty Datetime String",
			utcDatetimeString: "",
			timezone:          "America/New_York",
			expected:          "",
			expectError:       true,
		},
		{
			name:              "Empty Timezone",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "",
			expected:          "",
			expectError:       true,
		},
		{
			name:              "Valid Datetime and Timezone - Asia/Tokyo",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "Asia/Tokyo",
			expected:          "2023-10-27 19:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - UTC",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "UTC",
			expected:          "2023-10-27 10:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - Europe/Berlin",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "Europe/Berlin",
			expected:          "2023-10-27 12:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Los_Angeles",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Los_Angeles",
			expected:          "2023-10-27 03:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Chicago",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Chicago",
			expected:          "2023-10-27 05:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Denver",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Denver",
			expected:          "2023-10-27 04:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Phoenix",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Phoenix",
			expected:          "2023-10-27 03:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Anchorage",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Anchorage",
			expected:          "2023-10-27 02:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - America/Adak",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "America/Adak",
			expected:          "2023-10-27 01:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - Pacific/Honolulu",
			utcDatetimeString: "2023-10-27 10:00",
			timezone:          "Pacific/Honolulu",
			expected:          "2023-10-27 00:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - Daylight Saving Time",
			utcDatetimeString: "2023-03-12 05:00",
			timezone:          "America/New_York",
			expected:          "2023-03-12 00:00:00",
			expectError:       false,
		},
		{
			name:              "Valid Datetime and Timezone - Standard Time",
			utcDatetimeString: "2023-11-05 05:00",
			timezone:          "America/New_York",
			expected:          "2023-11-05 01:00:00",
			expectError:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tz.UTCDatetimeToTz(tc.utcDatetimeString, tc.timezone)

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
				t.Errorf("UTCDatetimeToTz(%q, %q) = %q, want %q", tc.utcDatetimeString, tc.timezone, result, tc.expected)
			}
		})
	}
}
