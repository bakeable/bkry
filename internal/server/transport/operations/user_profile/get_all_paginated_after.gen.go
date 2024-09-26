package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/data/entities/user_profile"
)


func afterGetAllPaginated(entities []user_profile.UserProfile, pageSize int, pageNumber int, orderBy string, ascending bool) []user_profile.UserProfile {
	return entities
}