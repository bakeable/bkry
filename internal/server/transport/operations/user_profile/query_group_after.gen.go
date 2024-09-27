package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/internal/server/models/entities/user_profile"
)


func afterQueryGroup(entities []user_profile.UserProfile) []user_profile.UserProfile {
	return entities
}