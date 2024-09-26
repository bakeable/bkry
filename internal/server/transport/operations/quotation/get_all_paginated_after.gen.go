package quotation_operations

import (
	quotation "github.com/bakeable/bkry/data/entities/quotation"
)


func afterGetAllPaginated(entities []quotation.Quotation, pageSize int, pageNumber int, orderBy string, ascending bool) []quotation.Quotation {
	return entities
}