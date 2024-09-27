package user_operations

import (
	user "github.com/bakeable/bkry/internal/server/models/entities/user"
)


func afterGetAll(entities []user.User) []user.User {
	return entities
}