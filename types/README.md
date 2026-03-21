# Generic Types

The types package provides generic type definitions commonly used across Go web applications.

## Features

- **FlashMessage**: Standard flash message structure for user notifications

## Installation

```go
import "github.com/dracory/base/types"
```

## Usage

### Flash Messages

Flash messages are used to display temporary notifications to users after actions like form submissions, redirects, or authentication events.

```go
message := types.FlashMessage{
    Type:    "success",
    Message: "Profile updated successfully",
    Url:     "/profile",
    Time:    time.Now().Format(time.RFC3339),
}

// In a redirect
http.Redirect(w, r, "/login?flash=success&message=Please+log+in", http.StatusTemporaryRedirect)
```

### Common Flash Message Types

- **"success"**: Successful operations (create, update, delete)
- **"error"**: Error conditions (validation, authentication, system errors)
- **"info"**: Informational messages (tips, notifications)
- **"warning"**: Warning messages (important notices, cautions)

## Types

### FlashMessage

Represents a flash message for web applications.

**Fields:**
- `Type` (string): Message type (e.g., "success", "error", "info", "warning")
- `Message` (string): The message content to display to users
- `Url` (string): Optional URL for redirect after message display
- `Time` (string): Timestamp when the message was created

**Example:**
```go
message := types.FlashMessage{
    Type:    "success",
    Message: "User created successfully",
    Url:     "/users",
    Time:    "2023-12-01T10:30:00Z",
}
```

## Examples

### Authentication Flow

```go
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Successful login
    message := types.FlashMessage{
        Type:    "success",
        Message: "Welcome back!",
        Url:     "/dashboard",
        Time:    time.Now().Format(time.RFC3339),
    }
    
    // Store message in session or cookie
    session.SetFlash(message)
    
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LoginFailedHandler(w http.ResponseWriter, r *http.Request) {
    // Failed login
    message := types.FlashMessage{
        Type:    "error",
        Message: "Invalid credentials",
        Url:     "/login",
        Time:    time.Now().Format(time.RFC3339),
    }
    
    session.SetFlash(message)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```

### Form Processing

```go
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    user := User{
        Name:  r.FormValue("name"),
        Email: r.FormValue("email"),
    }
    
    if err := user.Save(); err != nil {
        message := types.FlashMessage{
            Type:    "error",
            Message: "Failed to create user: " + err.Error(),
            Url:     "/users/new",
            Time:    time.Now().Format(time.RFC3339),
        }
        session.SetFlash(message)
        http.Redirect(w, r, "/users/new", http.StatusSeeOther)
        return
    }
    
    message := types.FlashMessage{
        Type:    "success",
        Message: "User created successfully",
        Url:     "/users",
        Time:    time.Now().Format(time.RFC3339),
    }
    session.SetFlash(message)
    http.Redirect(w, r, "/users", http.StatusSeeOther)
}
```

## Testing

The package includes comprehensive tests:

```bash
go test ./types
```

## Dependencies

- Go 1.16+
- No external dependencies
