package user_operations

import (
	user "github.com/bakeable/bkry/internal/server/models/entities/user"
)

func afterSave(entity user.User, editorID *string) {}