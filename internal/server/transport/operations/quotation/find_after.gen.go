package quotation_operations

import (
	quotation "github.com/bakeable/bkry/internal/server/models/entities/quotation"
)

func afterFind(entity quotation.Quotation) quotation.Quotation {
	return entity
}