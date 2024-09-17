package user_operations

import (
	user "github.com/bakeable/bkry/data/entities/user"
)

func afterGet(userID string, entity user.User) user.User {
	return entity
}