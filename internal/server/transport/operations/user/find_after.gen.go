package user_operations

import (
	user "github.com/bakeable/bkry/internal/server/models/entities/user"
)

func afterFind(entity user.User) user.User {
	return entity
}