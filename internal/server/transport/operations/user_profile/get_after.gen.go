package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/internal/server/models/entities/user_profile"
)

func afterGet(userProfileID string, entity user_profile.UserProfile) user_profile.UserProfile {
	return entity
}