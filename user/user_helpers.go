package user

import (
	"log"
	"strings"

	"github.com/dracory/userstore"
	"github.com/samber/lo"
)

// DisplayNameFull returns the full display name for a user.
// It combines first and last name, falling back to email if both are empty.
func DisplayNameFull(user userstore.UserInterface) string {
	if user == nil {
		return "n/a"
	}

	displayName := user.FirstName() + " " + user.LastName()

	if strings.TrimSpace(displayName) == "" {
		return user.Email()
	}

	return displayName
}

// IsClient checks if a user is marked as a client based on metadata.
// Returns true if the user has "is_client" metadata set to "yes".
func IsClient(user userstore.UserInterface) bool {
	if user == nil {
		return false
	}
	return user.Meta("is_client") == "yes"
}

// SetIsClient sets or removes the client status for a user.
// When isClient is true, sets "is_client" metadata to "yes", otherwise "no".
func SetIsClient(user userstore.UserInterface, isClient bool) userstore.UserInterface {
	if user == nil {
		return nil
	}
	value := lo.Ternary(isClient, "yes", "no")
	if err := user.SetMeta("is_client", value); err != nil {
		log.Println("Failed to set is_client meta", err)
	}
	return user
}
