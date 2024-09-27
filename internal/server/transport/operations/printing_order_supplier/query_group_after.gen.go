package printing_order_supplier_operations

import (
	printing_order_supplier "github.com/bakeable/bkry/internal/server/models/entities/printing_order_supplier"
)


func afterQueryGroup(entities []printing_order_supplier.PrintingOrderSupplier) []printing_order_supplier.PrintingOrderSupplier {
	return entities
}