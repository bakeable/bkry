package email_operations

import (
	email "github.com/bakeable/bkry/internal/server/models/entities/email"
)


func afterQueryGroup(entities []email.Email) []email.Email {
	return entities
}