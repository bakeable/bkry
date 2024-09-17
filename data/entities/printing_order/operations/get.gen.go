package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/data/entities/printing_order"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, printingOrderID string) printing_order.PrintingOrder {
	// Get PrintingOrder
	entity, err := repository.GetPrintingOrder(printingOrderID)
	if err != nil {
		panic(err)
	}

	// Process PrintingOrder entity
	entity = afterGet(printingOrderID, entity)

	// Return PrintingOrder
	return entity
}