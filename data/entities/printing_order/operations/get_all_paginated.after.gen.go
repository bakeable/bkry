package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/data/entities/printing_order"
)


func afterGetAllPaginated(entities []printing_order.PrintingOrder, pageSize int, pageNumber int, orderBy string, ascending bool) []printing_order.PrintingOrder {
	return entities
}