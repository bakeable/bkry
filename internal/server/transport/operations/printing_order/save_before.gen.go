package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/data/entities/printing_order"
)

func beforeSave(entity printing_order.PrintingOrder, editorID *string) printing_order.PrintingOrder {
	// Return PrintingOrder
	return entity
}