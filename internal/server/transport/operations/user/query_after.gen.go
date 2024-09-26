package user_operations

import (
	user "github.com/bakeable/bkry/data/entities/user"
)


func afterQuery(entities []user.User) []user.User {
	return entities
}