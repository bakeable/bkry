package quotation_operations

import (
	quotation "github.com/bakeable/bkry/internal/server/models/entities/quotation"
)


func afterGetAllPaginated(entities []quotation.Quotation, pageSize int, pageNumber int, orderBy string, ascending bool) []quotation.Quotation {
	return entities
}