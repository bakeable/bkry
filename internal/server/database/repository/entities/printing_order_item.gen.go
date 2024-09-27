package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PrintingOrderItemRepo = repository.NewRepository[*printing_order_item.PrintingOrderItem]()

// GetPrintingOrderItem retrieves a single PrintingOrderItem entity by its ID and printingOrderItemID.
func GetPrintingOrderItem(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error) {
	entityMap, err := PrintingOrderItemRepo.Read(printing_order_item.GetDocumentPath( printingOrderItemID))
	return printing_order_item.Decode(entityMap), err
}

// GetPrintingOrderItemOrNew retrieves a single PrintingOrderItem entity by its ID and printingOrderItemID.
func GetPrintingOrderItemOrNew(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error) {
	entityMap, err := PrintingOrderItemRepo.Read(printing_order_item.GetDocumentPath( printingOrderItemID))
	if err != nil || entityMap == nil {
		return printing_order_item.PrintingOrderItem{}, err
	}
	return printing_order_item.Decode(entityMap), err
}

// GetPrintingOrderItem retrieves a single PrintingOrderItem entity by its document path.
func GetPrintingOrderItemByPath(path string) (printing_order_item.PrintingOrderItem, error) {
	entityMap, err := PrintingOrderItemRepo.Read(path)
	return printing_order_item.Decode(entityMap), err
}

// FindPrintingOrderItem retrieves a PrintingOrderItem entity according to given queries.
func FindPrintingOrderItem(queries []datastore.Query) (printing_order_item.PrintingOrderItem, error) {
	entityMap, err := PrintingOrderItemRepo.Find(printing_order_item.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return printing_order_item.PrintingOrderItem{}, err
	}
	return printing_order_item.Decode(entityMap), err
}

// GetAllPrintingOrderItems retrieves all PrintingOrderItem entities.
func GetAllPrintingOrderItems() ([]printing_order_item.PrintingOrderItem, error) {
	entityMaps, err := PrintingOrderItemRepo.ReadAll(printing_order_item.GetCollectionPath())
	if err != nil {
		return []printing_order_item.PrintingOrderItem{}, err
	}
	return printing_order_item.DecodeAll(entityMaps), nil
}


// GetAllPrintingOrderItemsPaginated retrieves all PrintingOrderItem entities in a paginated manner.
func GetAllPrintingOrderItemsPaginated(pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, datastore.Pagination, error) {
	entityMaps, pagination, err := PrintingOrderItemRepo.ReadPaginated(printing_order_item.GetCollectionPath(), pagination)
	if err != nil {
		return []printing_order_item.PrintingOrderItem{}, pagination, err
	}
	return printing_order_item.DecodeAll(entityMaps), pagination, nil
}

// QueryPrintingOrderItems retrieves all PrintingOrderItem entities according to given queries.
func QueryPrintingOrderItems(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, error) {
	entityMaps, err := PrintingOrderItemRepo.Query(printing_order_item.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []printing_order_item.PrintingOrderItem{}, err
	}
	return printing_order_item.DecodeAll(entityMaps), nil
}

// QueryPrintingOrderItemsGroup retrieves all PrintingOrderItem entities according to given queries.
func QueryPrintingOrderItemsGroup(queries []datastore.Query) ([]printing_order_item.PrintingOrderItem, error) {
	entityMaps, err := PrintingOrderItemRepo.QueryGroup("printing_order_items", queries)
	if err != nil {
		return []printing_order_item.PrintingOrderItem{}, err
	}
	return printing_order_item.DecodeAll(entityMaps), nil
}

// CreatePrintingOrderItem creates a new PrintingOrderItem entity.
func CreatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	return PrintingOrderItemRepo.Create(&entity, editorID)
}

// UpdatePrintingOrderItem updates an existing PrintingOrderItem entity.
func UpdatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	return PrintingOrderItemRepo.Update(&entity, editorID)
}

// SavePrintingOrderItem saves a PrintingOrderItem entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePrintingOrderItem(entity, editorID)
	} else {
		return UpdatePrintingOrderItem(entity, editorID)
	}
}

// DeletePrintingOrderItem deletes a PrintingOrderItem entity by its parents' path and printingOrderItemID.
func DeletePrintingOrderItem(printingOrderItemID string) error {
	return PrintingOrderItemRepo.Delete(printing_order_item.GetDocumentPath(printingOrderItemID))
}
