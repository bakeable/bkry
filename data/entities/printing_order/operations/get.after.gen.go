package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/data/entities/printing_order"
)

func afterGet(printingOrderID string, entity printing_order.PrintingOrder) printing_order.PrintingOrder {
	return entity
}