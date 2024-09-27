package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/internal/server/models/entities/user_profile"
)

func afterSave(entity user_profile.UserProfile, editorID *string) {}