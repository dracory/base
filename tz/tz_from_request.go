package tz

import (
	"net/http"

	basesession "github.com/dracory/base/session"
	"github.com/dracory/userstore"
)

// FromUser returns the timezone from the given user, or UTC if the user
// is nil or has no timezone set.
func FromUser(user userstore.UserInterface) string {
	defaultTimezone := "UTC"
	if user == nil {
		return defaultTimezone
	}

	userTimezone := user.GetTimezone()
	if userTimezone != "" {
		return userTimezone
	}

	return defaultTimezone
}

// FromRequest returns the timezone from the authenticated user in the
// request context, or UTC if the user is not authenticated or has no
// timezone set.
func FromRequest(r *http.Request) string {
	return FromUser(basesession.GetAuthUser(r))
}
