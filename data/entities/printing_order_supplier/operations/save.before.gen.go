package printing_order_supplier_operations

import (
	printing_order_supplier "github.com/bakeable/bkry/data/entities/printing_order_supplier"
)

func beforeSave(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) printing_order_supplier.PrintingOrderSupplier {
	// Return PrintingOrderSupplier
	return entity
}