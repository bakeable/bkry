package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []datastore.Query) []printing_order_item.PrintingOrderItem {
	// Query PrintingOrderItems group
	entities, err := repository.QueryPrintingOrderItemsGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process PrintingOrderItem entities
	entities = afterQueryGroup(entities)

	// Return PrintingOrderItems
	return entities
}