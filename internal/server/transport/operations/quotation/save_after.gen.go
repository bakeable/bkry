package quotation_operations

import (
	quotation "github.com/bakeable/bkry/internal/server/models/entities/quotation"
)

func afterSave(entity quotation.Quotation, editorID *string) {}