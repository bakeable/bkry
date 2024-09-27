package email_operations

import (
	email "github.com/bakeable/bkry/internal/server/models/entities/email"
)

func beforeSave(entity email.Email, editorID *string) email.Email {
	// Return Email
	return entity
}