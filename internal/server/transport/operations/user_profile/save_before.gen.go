package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/data/entities/user_profile"
)

func beforeSave(entity user_profile.UserProfile, editorID *string) user_profile.UserProfile {
	// Return UserProfile
	return entity
}