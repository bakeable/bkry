package email_operations

import (
	email "github.com/bakeable/bkry/data/entities/email"
)

func afterGet(emailID string, entity email.Email) email.Email {
	return entity
}