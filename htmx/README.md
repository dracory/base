# HTMX Utilities

This package provides utility functions for working with HTMX requests in Go applications.

## Functions

### Request Detection

- `IsHtmx(r *http.Request) bool` - Checks if the request is an HTMX request
- `IsHxBoosted(r *http.Request) bool` - Checks if the request was boosted
- `IsHxHistoryRestoreRequest(r *http.Request) bool` - Checks if it's a history restore request
- `IsHxRequest(r *http.Request) bool` - Checks if it's an HTMX request (validates value is "true")
- `IsHxTrigger(r *http.Request) bool` - Checks if the request was triggered by an event

### Header Value Extraction

- `HxPrompt(r *http.Request) string` - Gets the prompt message
- `HxTarget(r *http.Request) string` - Gets the target element
- `HxTriggerName(r *http.Request) string` - Gets the trigger name

### CSS Utilities

- `HxHideIndicatorCSS() string` - Returns CSS for hiding HTMX indicators

## Usage

```go
package main

import (
    "net/http"
    "github.com/dracory/base/htmx"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if htmx.IsHtmx(r) {
        // Handle HTMX request
        target := htmx.HxTarget(r)
        trigger := htmx.HxTriggerName(r)
        
        // Process HTMX-specific logic
    }
    
    // Regular request handling
}
```

## HTMX Headers Supported

- `HX-Request` - Indicates an HTMX request
- `HX-Boosted` - Indicates the request was boosted
- `HX-History-Restore-Request` - Indicates a history restore request
- `HX-Trigger` - Indicates the request was triggered by an event
- `HX-Trigger-Name` - The name of the trigger
- `HX-Target` - The target element
- `HX-Prompt` - The prompt message

## Dependencies

This package depends on the `github.com/dracory/base/req` package for header value extraction.
