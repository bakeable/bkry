package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)

func afterGet(printingOrderItemID string, entity printing_order_item.PrintingOrderItem) printing_order_item.PrintingOrderItem {
	return entity
}