package session

import (
	"context"
	"net/http"

	"github.com/dracory/userstore"
)

// AuthenticatedUserContextKey is the context key used to store the authenticated user
type AuthenticatedUserContextKey struct{}

// APIAuthenticatedUserContextKey is the context key used to store the API authenticated user
type APIAuthenticatedUserContextKey struct{}

// GetAuthUser returns the authenticated user from the request context.
// It retrieves the user that was previously stored in the context using
// the AuthenticatedUserContextKey.
func GetAuthUser(r *http.Request) userstore.UserInterface {
	if r == nil {
		return nil
	}

	value := r.Context().Value(AuthenticatedUserContextKey{})

	if value == nil {
		return nil
	}

	user, ok := value.(userstore.UserInterface)
	if !ok {
		return nil
	}

	return user
}

// GetAPIAuthUser returns the authenticated user for API context.
// It retrieves the user that was previously stored in the context using
// the APIAuthenticatedUserContextKey.
func GetAPIAuthUser(r *http.Request) userstore.UserInterface {
	if r == nil {
		return nil
	}

	value := r.Context().Value(APIAuthenticatedUserContextKey{})

	if value == nil {
		return nil
	}

	user, ok := value.(userstore.UserInterface)
	if !ok {
		return nil
	}

	return user
}

// GetAuthUserFromContext returns the authenticated user from a context.
// This is useful when you have a context but not the original request.
func GetAuthUserFromContext(ctx context.Context) userstore.UserInterface {
	if ctx == nil {
		return nil
	}

	value := ctx.Value(AuthenticatedUserContextKey{})

	if value == nil {
		return nil
	}

	user, ok := value.(userstore.UserInterface)
	if !ok {
		return nil
	}

	return user
}

// GetAPIAuthUserFromContext returns the API authenticated user from a context.
// This is useful when you have a context but not the original request.
func GetAPIAuthUserFromContext(ctx context.Context) userstore.UserInterface {
	if ctx == nil {
		return nil
	}

	value := ctx.Value(APIAuthenticatedUserContextKey{})

	if value == nil {
		return nil
	}

	user, ok := value.(userstore.UserInterface)
	if !ok {
		return nil
	}

	return user
}

// SetAuthUser sets the authenticated user in the request context.
// This is typically used by authentication middleware to store the user
// after successful authentication.
func SetAuthUser(r *http.Request, user userstore.UserInterface) *http.Request {
	if r == nil || user == nil {
		return r
	}

	ctx := SetAuthUserInContext(r.Context(), user)
	return r.WithContext(ctx)
}

// SetAPIAuthUser sets the API authenticated user in the request context.
// This is typically used by API authentication middleware to store the user
// after successful authentication.
func SetAPIAuthUser(r *http.Request, user userstore.UserInterface) *http.Request {
	if r == nil || user == nil {
		return r
	}

	ctx := SetAPIAuthUserInContext(r.Context(), user)
	return r.WithContext(ctx)
}

// SetAuthUserInContext sets the authenticated user in the context.
// Returns a new context with the user stored.
func SetAuthUserInContext(ctx context.Context, user userstore.UserInterface) context.Context {
	if ctx == nil || user == nil {
		return ctx
	}

	return context.WithValue(ctx, AuthenticatedUserContextKey{}, user)
}

// SetAPIAuthUserInContext sets the API authenticated user in the context.
// Returns a new context with the user stored.
func SetAPIAuthUserInContext(ctx context.Context, user userstore.UserInterface) context.Context {
	if ctx == nil || user == nil {
		return ctx
	}

	return context.WithValue(ctx, APIAuthenticatedUserContextKey{}, user)
}
