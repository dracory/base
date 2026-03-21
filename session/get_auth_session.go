package session

import (
	"context"
	"net/http"

	"github.com/dracory/sessionstore"
)

// AuthenticatedSessionContextKey is the context key used to store the authenticated session
type AuthenticatedSessionContextKey struct{}

// GetAuthSession returns the authenticated session from the request context.
// It retrieves the session that was previously stored in the context using
// the AuthenticatedSessionContextKey.
func GetAuthSession(r *http.Request) sessionstore.SessionInterface {
	if r == nil {
		return nil
	}

	value := r.Context().Value(AuthenticatedSessionContextKey{})

	if value == nil {
		return nil
	}

	session, ok := value.(sessionstore.SessionInterface)
	if !ok {
		return nil
	}

	return session
}

// GetAuthSessionFromContext returns the authenticated session from a context.
// This is useful when you have a context but not the original request.
func GetAuthSessionFromContext(ctx context.Context) sessionstore.SessionInterface {
	if ctx == nil {
		return nil
	}

	value := ctx.Value(AuthenticatedSessionContextKey{})

	if value == nil {
		return nil
	}

	session, ok := value.(sessionstore.SessionInterface)
	if !ok {
		return nil
	}

	return session
}

// SetAuthSession sets the authenticated session in the request context.
// This is typically used by authentication middleware to store the session
// after successful authentication.
func SetAuthSession(r *http.Request, session sessionstore.SessionInterface) *http.Request {
	if r == nil || session == nil {
		return r
	}

	ctx := SetAuthSessionInContext(r.Context(), session)
	return r.WithContext(ctx)
}

// SetAuthSessionInContext sets the authenticated session in the context.
// Returns a new context with the session stored.
func SetAuthSessionInContext(ctx context.Context, session sessionstore.SessionInterface) context.Context {
	if ctx == nil || session == nil {
		return ctx
	}

	return context.WithValue(ctx, AuthenticatedSessionContextKey{}, session)
}
