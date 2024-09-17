package email_operations

import (
	email "github.com/bakeable/bkry/data/entities/email"
)

func beforeSave(entity email.Email, editorID *string) email.Email {
	// Return Email
	return entity
}