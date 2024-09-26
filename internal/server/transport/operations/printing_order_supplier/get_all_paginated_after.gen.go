package printing_order_supplier_operations

import (
	printing_order_supplier "github.com/bakeable/bkry/data/entities/printing_order_supplier"
)


func afterGetAllPaginated(entities []printing_order_supplier.PrintingOrderSupplier, pageSize int, pageNumber int, orderBy string, ascending bool) []printing_order_supplier.PrintingOrderSupplier {
	return entities
}