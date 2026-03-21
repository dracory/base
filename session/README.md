# Session Package

The `session` package provides utilities for managing authenticated sessions and users in HTTP applications. It offers a framework-agnostic approach to session and user management that works with any `sessionstore` and `userstore` implementation.

## Features

- **Session Management**: Retrieve, store, and extend authenticated sessions
- **User Management**: Access authenticated users from request context
- **Context Integration**: Seamless integration with HTTP request contexts
- **Security Validation**: IP address and user agent validation for session security
- **User Settings**: Store and retrieve user-specific settings in sessions
- **Framework Agnostic**: Works with any HTTP framework and store implementation

## Usage

### Session Management

```go
import (
    "github.com/dracory/base/session"
    "github.com/dracory/sessionstore"
    "net/http"
)

// Get authenticated session from request
session := session.GetAuthSession(r)
if session != nil {
    // Session is available
    userID := session.GetUserID()
    expiresAt := session.GetExpiresAt()
}

// Extend session expiration
err := session.ExtendSession(sessionStore, r, 3600) // Extend by 1 hour
if err != nil {
    log.Printf("Failed to extend session: %v", err)
}

// Set session in request context (typically done by middleware)
req = session.SetAuthSession(r, authenticatedSession)
```

### User Management

```go
// Get authenticated user from request
user := session.GetAuthUser(r)
if user != nil {
    displayName := session.DisplayNameFull(user)
    email := user.Email()
    userID := user.ID()
}

// Check if user is a client
isClient := session.IsClient(user)
if isClient {
    // Handle client-specific logic
}

// Set user as client
user = session.SetIsClient(user, true)
```

### User Settings

```go
// Store user-specific setting
err := session.UserSettingSet(sessionStore, r, "theme_preference", "dark")
if err != nil {
    log.Printf("Failed to store user setting: %v", err)
}

// Retrieve user-specific setting
theme := session.UserSettingGet(sessionStore, r, "theme_preference", "light")
if theme == "dark" {
    // Apply dark theme
}
```

### Context Management

```go
// Get session from context (useful when you have context but not request)
session := session.GetAuthSessionFromContext(ctx)

// Get user from context
user := session.GetAuthUserFromContext(ctx)

// Set session in context
ctx = session.SetAuthSessionInContext(ctx, session)

// Set user in context
ctx = session.SetAuthUserInContext(ctx, user)
```

## Context Keys

The package defines context keys for storing authenticated sessions and users:

- `AuthenticatedSessionContextKey{}`: Stores the authenticated session
- `AuthenticatedUserContextKey{}`: Stores the authenticated user
- `APIAuthenticatedUserContextKey{}`: Stores API authenticated user

## Security Features

### Session Extension Security

When extending a session, the `ExtendSession` function validates:

1. **Session Store**: Must not be nil
2. **Session Existence**: Session must exist in the context
3. **IP Address**: Session IP must match current request IP
4. **User Agent**: Session user agent must match current request user agent

### User Settings Security

When storing/retrieving user settings, the functions validate:

1. **User Ownership**: Settings belong to the authenticated user
2. **IP Address**: Session IP must match current request IP
3. **User Agent**: Session user agent must match current request user agent

## Integration Examples

### Authentication Middleware

```go
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Validate session token
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Find session in store
        session, err := sessionStore.SessionFindByToken(r.Context(), token)
        if err != nil || session == nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Find user
        user, err := userStore.UserFindByID(r.Context(), session.GetUserID())
        if err != nil || user == nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Store session and user in context
        r = session.SetAuthSession(r, session)
        r = session.SetAuthUser(r, user)

        next.ServeHTTP(w, r)
    })
}
```

### Session Extension Handler

```go
func ExtendSessionHandler(w http.ResponseWriter, r *http.Request) {
    // Extend session by 30 minutes
    err := session.ExtendSession(sessionStore, r, 1800)
    if err != nil {
        http.Error(w, "Failed to extend session", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Session extended"))
}
```

### User Preferences API

```go
func GetUserPreferenceHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        http.Error(w, "Key parameter required", http.StatusBadRequest)
        return
    }

    // Get user preference with default
    value := session.UserSettingGet(sessionStore, r, key, "default")
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "key": key,
        "value": value,
    })
}

func SetUserPreferenceHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Key   string `json:"key"`
        Value string `json:"value"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    err := session.UserSettingSet(sessionStore, r, req.Key, req.Value)
    if err != nil {
        http.Error(w, "Failed to store preference", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
```

## Dependencies

- `github.com/dracory/sessionstore`: Session store interface and implementations
- `github.com/dracory/userstore`: User store interface and implementations  
- `github.com/dracory/req`: HTTP request utilities (IP extraction)
- `github.com/dromara/carbon/v2`: Date/time utilities
- `github.com/spf13/cast`: Type casting utilities

## Error Handling

The package provides descriptive error messages for common failure scenarios:

- `"session store is nil"`: Session store is not initialized
- `"session not found"`: No session found in request context
- `"session ip address does not match request ip address"`: IP validation failed
- `"session user agent does not match request user agent"`: User agent validation failed
- `"auth user is nil"`: No authenticated user found
- `"session user id does not match auth user id"`: User ownership validation failed

## Testing

The package includes comprehensive tests covering:

- Session retrieval and storage
- User access and management
- Context manipulation
- Security validation
- Error handling
- User settings management

Run tests with:
```bash
go test ./session/...
```

## License

This package is part of the Dracory base package and follows the same licensing terms.
