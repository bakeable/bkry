package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/data/entities/printing_order_item"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, printingOrderItemID string) printing_order_item.PrintingOrderItem {
	// Get PrintingOrderItem
	entity, err := repository.GetPrintingOrderItem(printingOrderItemID)
	if err != nil {
		panic(err)
	}

	// Process PrintingOrderItem entity
	entity = afterGet(printingOrderItemID, entity)

	// Return PrintingOrderItem
	return entity
}