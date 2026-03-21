# User Package

The `user` package provides utility functions for working with user objects, particularly focused on display name formatting and client status management. It offers framework-agnostic user helper functions that work with any `userstore` implementation.

## Features

- **Display Name Formatting**: Intelligent display name generation with fallback to email
- **Client Status Management**: Check and set client status for users
- **Framework Agnostic**: Works with any userstore implementation
- **Metadata Integration**: Seamless integration with user metadata

## Usage

### Display Name Formatting

```go
import (
    "github.com/dracory/base/user"
    "github.com/dracory/userstore"
)

// Create or retrieve a user
user := userstore.NewUser()
user.SetFirstName("John")
user.SetLastName("Doe")
user.SetEmail("john.doe@example.com")

// Get formatted display name
displayName := user.DisplayNameFull(user)
// Returns: "John Doe"

// Fallback to email when names are empty
user.SetFirstName("")
user.SetLastName("")
displayName = user.DisplayNameFull(user)
// Returns: "john.doe@example.com"
```

### Client Status Management

```go
// Check if user is marked as a client
isClient := user.IsClient(user)
if isClient {
    // Apply client-specific logic
    fmt.Println("User is a client")
}

// Mark user as a client
user = user.SetIsClient(user, true)
if user.IsClient(user) {
    fmt.Println("User is now marked as a client")
}

// Remove client status
user = user.SetIsClient(user, false)
if !user.IsClient(user) {
    fmt.Println("User is no longer marked as a client")
}
```

## Function Details

### DisplayNameFull(user userstore.UserInterface) string

Generates a full display name from a user object with intelligent fallback logic:

1. **Primary**: Combines first and last name: "John Doe"
2. **Fallback**: If combined name is empty or whitespace, returns email
3. **Default**: If user is nil, returns "n/a"

**Behavior Examples:**
```go
user.SetFirstName("John"), user.SetLastName("Doe")     → "John Doe"
user.SetFirstName("John"), user.SetLastName("")       → "John "  
user.SetFirstName(""), user.SetLastName("Doe")         → " Doe"
user.SetFirstName(""), user.SetLastName("")            → "user@example.com"
user = nil                                            → "n/a"
```

### IsClient(user userstore.UserInterface) bool

Checks if a user is marked as a client by looking for the `"is_client"` metadata key:

- Returns `true` if `user.Meta("is_client") == "yes"`
- Returns `false` for any other value or if user is nil
- Case-sensitive comparison

### SetIsClient(user userstore.UserInterface, isClient bool) userstore.UserInterface

Sets or removes the client status for a user:

- When `isClient` is `true`: Sets `"is_client"` metadata to `"yes"`
- When `isClient` is `false`: Sets `"is_client"` metadata to `"no"`
- Returns the same user instance (or nil if input is nil)
- Logs errors if metadata setting fails

## Integration Examples

### User Profile Display

```go
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
    currentUser := session.GetAuthUser(r)
    if currentUser == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    profile := map[string]interface{}{
        "displayName": user.DisplayNameFull(currentUser),
        "email": currentUser.Email(),
        "isClient": user.IsClient(currentUser),
        "role": currentUser.Role(),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profile)
}
```

### Client-Specific Features

```go
func FeatureAccessMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        currentUser := session.GetAuthUser(r)
        
        // Check if user is a client for client-specific features
        if user.IsClient(currentUser) {
            // Grant access to client features
            r = r.WithContext(context.WithValue(r.Context(), "client_features", true))
        }
        
        next.ServeHTTP(w, r)
    })
}
```

### User Management Interface

```go
func UserManagementHandler(w http.ResponseWriter, r *http.Request) {
    users, err := userStore.UserFindAll(r.Context())
    if err != nil {
        http.Error(w, "Failed to load users", http.StatusInternalServerError)
        return
    }

    type UserInfo struct {
        ID          string `json:"id"`
        DisplayName string `json:"displayName"`
        Email       string `json:"email"`
        IsClient    bool   `json:"isClient"`
        Status      string `json:"status"`
    }

    var userInfos []UserInfo
    for _, u := range users {
        userInfos = append(userInfos, UserInfo{
            ID:          u.ID(),
            DisplayName: user.DisplayNameFull(u),
            Email:       u.Email(),
            IsClient:    user.IsClient(u),
            Status:      u.Status(),
        })
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userInfos)
}
```

### Client Status Toggle

```go
func ToggleClientStatusHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(w, "User ID required", http.StatusBadRequest)
        return
    }

    user, err := userStore.UserFindByID(r.Context(), userID)
    if err != nil || user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Toggle client status
    currentStatus := user.IsClient(user)
    user = user.SetIsClient(user, !currentStatus)

    // Save user
    err = userStore.UserUpdate(r.Context(), user)
    if err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "userID": user.ID(),
        "isClient": user.IsClient(user),
        "message": fmt.Sprintf("Client status %s", map[bool]string{true: "enabled", false: "disabled"}[user.IsClient(user)]),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
```

## Dependencies

- `github.com/dracory/userstore`: User store interface and implementations
- `github.com/samber/lo`: Functional utilities (for ternary operations)

## Error Handling

The package is designed to be safe and defensive:

- **Nil User Handling**: All functions gracefully handle nil user objects
- **Metadata Errors**: `SetIsClient` logs errors but doesn't fail the operation
- **Type Safety**: Uses proper interface types for framework compatibility

## Best Practices

### Display Names

- Always use `DisplayNameFull()` for user-facing display names
- Don't manually concatenate names and emails
- The function handles edge cases like empty names and whitespace

### Client Status

- Use `IsClient()` to check client status before applying client-specific logic
- Use `SetIsClient()` to manage client status programmatically
- Client status is stored as metadata, making it flexible and extensible

### Error Handling

- Always check for nil users before calling helper functions
- The functions return safe defaults ("n/a", false) for edge cases
- Metadata setting errors are logged but don't prevent operation

## Testing

The package includes comprehensive tests covering:

- Display name formatting with various name combinations
- Client status checking and setting
- Nil user handling
- Edge cases and error conditions

Run tests with:
```bash
go test ./user/...
```

## License

This package is part of the Dracory base package and follows the same licensing terms.
