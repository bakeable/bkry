package user_operations

import (
	user "github.com/bakeable/bkry/data/entities/user"
)


func afterGetAllPaginated(entities []user.User, pageSize int, pageNumber int, orderBy string, ascending bool) []user.User {
	return entities
}