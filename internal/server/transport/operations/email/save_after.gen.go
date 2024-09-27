package email_operations

import (
	email "github.com/bakeable/bkry/internal/server/models/entities/email"
)

func afterSave(entity email.Email, editorID *string) {}