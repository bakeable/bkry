package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)


func afterGetAll(entities []printing_order_item.PrintingOrderItem) []printing_order_item.PrintingOrderItem {
	return entities
}