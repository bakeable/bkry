package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity printing_order_item.PrintingOrderItem, editorID *string) *string {
	// Preprocess PrintingOrderItem
	entity = beforeSave(entity, editorID)

	// Save PrintingOrderItem
	id, err := repository.SavePrintingOrderItem(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving PrintingOrderItem entity
	afterSave(entity, editorID)

	// Return PrintingOrderItem
	return &id
}