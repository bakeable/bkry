package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)

func beforeSave(entity printing_order_item.PrintingOrderItem, editorID *string) printing_order_item.PrintingOrderItem {
	// Return PrintingOrderItem
	return entity
}