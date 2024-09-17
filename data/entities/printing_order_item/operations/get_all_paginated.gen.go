package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/data/entities/printing_order_item"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, datastore.Pagination) {
	// Get PrintingOrderItems
	entities, pagination, err := repository.GetAllPrintingOrderItemsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process PrintingOrderItem entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return PrintingOrderItems
	return entities, pagination
}