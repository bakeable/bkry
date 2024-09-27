package user_operations

import (
	user "github.com/bakeable/bkry/internal/server/models/entities/user"
)

func beforeSave(entity user.User, editorID *string) user.User {
	// Return User
	return entity
}