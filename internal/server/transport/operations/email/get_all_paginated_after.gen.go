package email_operations

import (
	email "github.com/bakeable/bkry/internal/server/models/entities/email"
)


func afterGetAllPaginated(entities []email.Email, pageSize int, pageNumber int, orderBy string, ascending bool) []email.Email {
	return entities
}