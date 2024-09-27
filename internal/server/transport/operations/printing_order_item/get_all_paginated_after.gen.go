package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)


func afterGetAllPaginated(entities []printing_order_item.PrintingOrderItem, pageSize int, pageNumber int, orderBy string, ascending bool) []printing_order_item.PrintingOrderItem {
	return entities
}