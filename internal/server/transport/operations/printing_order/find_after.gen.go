package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/internal/server/models/entities/printing_order"
)

func afterFind(entity printing_order.PrintingOrder) printing_order.PrintingOrder {
	return entity
}