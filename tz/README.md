# Timezone Package

The `tz` package provides timezone conversion utilities for the Dracory framework.
It offers functions for converting UTC dates, times, and datetimes to different timezones.

## Overview

This package includes utilities for:

1. **Date Conversion**: Convert UTC dates to different timezones
2. **Time Conversion**: Convert UTC times to different timezones
3. **Datetime Conversion**: Convert UTC datetimes to different timezones

## Key Features

- **Simple API**: Easy-to-use functions for timezone conversion
- **Flexible Format Support**: Handles various date and time formats
- **Error Handling**: Proper error handling for invalid inputs
- **Carbon Integration**: Built on top of the Carbon library for robust datetime handling

## Usage Examples

### Converting UTC Date to Timezone

```go
import "github.com/dracory/base/tz"

// Convert a UTC date to New York timezone
dateInNewYork, err := tz.UTCDateToTz("2022-01-01", "America/New_York")
if err != nil {
    // Handle error
}
fmt.Println(dateInNewYork) // "2021-12-31" (if UTC date is Jan 1, it might be Dec 31 in NY)
```

### Converting UTC Time to Timezone

```go
import "github.com/dracory/base/tz"

// Convert a UTC time to Tokyo timezone
timeInTokyo := tz.UTCTimeToTz("10:00", "Asia/Tokyo")
fmt.Println(timeInTokyo) // "19:00" (Tokyo is 9 hours ahead of UTC)
```

### Converting UTC Datetime to Timezone

```go
import "github.com/dracory/base/tz"

// Convert a UTC datetime to London timezone
datetimeInLondon, err := tz.UTCDatetimeToTz("2022-01-01 10:00", "Europe/London")
if err != nil {
    // Handle error
}
fmt.Println(datetimeInLondon) // "2022-01-01 10:00" (London is at UTC+0 in winter)
```

## Supported Timezone Formats

The package supports standard IANA timezone identifiers, such as:

- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`
- `Africa/Cairo`
- `Pacific/Auckland`

## Best Practices

1. **Always Check for Errors**: The date and datetime conversion functions return errors that should be handled
2. **Use Valid Timezone Identifiers**: Ensure you're using valid IANA timezone identifiers
3. **Consider Daylight Saving Time**: Be aware that timezone conversions may be affected by daylight saving time
4. **Handle Edge Cases**: Be prepared for edge cases, such as date changes when converting between timezones

## Implementation Details

The package uses the [Carbon](https://github.com/dromara/carbon) library for datetime handling, which provides robust timezone conversion capabilities. The functions in this package are thin wrappers around Carbon's functionality, providing a simplified API for common timezone conversion tasks.

## License

This package is part of the dracory/base project and is licensed under the same terms. 