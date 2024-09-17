package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/data/entities/printing_order_item"
)


func afterQueryGroup(entities []printing_order_item.PrintingOrderItem) []printing_order_item.PrintingOrderItem {
	return entities
}