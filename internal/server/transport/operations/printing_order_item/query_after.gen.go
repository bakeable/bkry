package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)


func afterQuery(entities []printing_order_item.PrintingOrderItem) []printing_order_item.PrintingOrderItem {
	return entities
}