package repo

import (
	"github.com/bakeable/bkry/data/entities/printing_order"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PrintingOrderRepo = repository.NewRepository[*printing_order.PrintingOrder]()

// GetPrintingOrder retrieves a single PrintingOrder entity by its ID and printingOrderID.
func GetPrintingOrder(printingOrderID string) (printing_order.PrintingOrder, error) {
	entityMap, err := PrintingOrderRepo.Read(printing_order.GetDocumentPath( printingOrderID))
	return printing_order.Decode(entityMap), err
}

// GetPrintingOrderOrNew retrieves a single PrintingOrder entity by its ID and printingOrderID.
func GetPrintingOrderOrNew(printingOrderID string) (printing_order.PrintingOrder, error) {
	entityMap, err := PrintingOrderRepo.Read(printing_order.GetDocumentPath( printingOrderID))
	if err != nil || entityMap == nil {
		return printing_order.PrintingOrder{}, err
	}
	return printing_order.Decode(entityMap), err
}

// GetPrintingOrder retrieves a single PrintingOrder entity by its document path.
func GetPrintingOrderByPath(path string) (printing_order.PrintingOrder, error) {
	entityMap, err := PrintingOrderRepo.Read(path)
	return printing_order.Decode(entityMap), err
}

// FindPrintingOrder retrieves a PrintingOrder entity according to given queries.
func FindPrintingOrder(queries []datastore.Query) (printing_order.PrintingOrder, error) {
	entityMap, err := PrintingOrderRepo.Find(printing_order.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return printing_order.PrintingOrder{}, err
	}
	return printing_order.Decode(entityMap), err
}

// GetAllPrintingOrders retrieves all PrintingOrder entities.
func GetAllPrintingOrders() ([]printing_order.PrintingOrder, error) {
	entityMaps, err := PrintingOrderRepo.ReadAll(printing_order.GetCollectionPath())
	if err != nil {
		return []printing_order.PrintingOrder{}, err
	}
	return printing_order.DecodeAll(entityMaps), nil
}


// GetAllPrintingOrdersPaginated retrieves all PrintingOrder entities in a paginated manner.
func GetAllPrintingOrdersPaginated(pagination datastore.Pagination) ([]printing_order.PrintingOrder, datastore.Pagination, error) {
	entityMaps, pagination, err := PrintingOrderRepo.ReadPaginated(printing_order.GetCollectionPath(), pagination)
	if err != nil {
		return []printing_order.PrintingOrder{}, pagination, err
	}
	return printing_order.DecodeAll(entityMaps), pagination, nil
}

// QueryPrintingOrders retrieves all PrintingOrder entities according to given queries.
func QueryPrintingOrders(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order.PrintingOrder, error) {
	entityMaps, err := PrintingOrderRepo.Query(printing_order.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []printing_order.PrintingOrder{}, err
	}
	return printing_order.DecodeAll(entityMaps), nil
}

// QueryPrintingOrdersGroup retrieves all PrintingOrder entities according to given queries.
func QueryPrintingOrdersGroup(queries []datastore.Query) ([]printing_order.PrintingOrder, error) {
	entityMaps, err := PrintingOrderRepo.QueryGroup("printing_orders", queries)
	if err != nil {
		return []printing_order.PrintingOrder{}, err
	}
	return printing_order.DecodeAll(entityMaps), nil
}

// CreatePrintingOrder creates a new PrintingOrder entity.
func CreatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	return PrintingOrderRepo.Create(&entity, editorID)
}

// UpdatePrintingOrder updates an existing PrintingOrder entity.
func UpdatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	return PrintingOrderRepo.Update(&entity, editorID)
}

// SavePrintingOrder saves a PrintingOrder entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePrintingOrder(entity, editorID)
	} else {
		return UpdatePrintingOrder(entity, editorID)
	}
}

// DeletePrintingOrder deletes a PrintingOrder entity by its parents' path and printingOrderID.
func DeletePrintingOrder(printingOrderID string) error {
	return PrintingOrderRepo.Delete(printing_order.GetDocumentPath(printingOrderID))
}
