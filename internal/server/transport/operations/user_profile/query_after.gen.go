package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/data/entities/user_profile"
)


func afterQuery(entities []user_profile.UserProfile) []user_profile.UserProfile {
	return entities
}