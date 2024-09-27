package quotation_operations

import (
	quotation "github.com/bakeable/bkry/internal/server/models/entities/quotation"
)

func beforeSave(entity quotation.Quotation, editorID *string) quotation.Quotation {
	// Return Quotation
	return entity
}