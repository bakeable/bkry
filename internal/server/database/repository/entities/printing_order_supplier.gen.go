package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/printing_order_supplier"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PrintingOrderSupplierRepo = repository.NewRepository[*printing_order_supplier.PrintingOrderSupplier]()

// GetPrintingOrderSupplier retrieves a single PrintingOrderSupplier entity by its ID and printingOrderSupplierID.
func GetPrintingOrderSupplier(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error) {
	entityMap, err := PrintingOrderSupplierRepo.Read(printing_order_supplier.GetDocumentPath( printingOrderSupplierID))
	return printing_order_supplier.Decode(entityMap), err
}

// GetPrintingOrderSupplierOrNew retrieves a single PrintingOrderSupplier entity by its ID and printingOrderSupplierID.
func GetPrintingOrderSupplierOrNew(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error) {
	entityMap, err := PrintingOrderSupplierRepo.Read(printing_order_supplier.GetDocumentPath( printingOrderSupplierID))
	if err != nil || entityMap == nil {
		return printing_order_supplier.PrintingOrderSupplier{}, err
	}
	return printing_order_supplier.Decode(entityMap), err
}

// GetPrintingOrderSupplier retrieves a single PrintingOrderSupplier entity by its document path.
func GetPrintingOrderSupplierByPath(path string) (printing_order_supplier.PrintingOrderSupplier, error) {
	entityMap, err := PrintingOrderSupplierRepo.Read(path)
	return printing_order_supplier.Decode(entityMap), err
}

// FindPrintingOrderSupplier retrieves a PrintingOrderSupplier entity according to given queries.
func FindPrintingOrderSupplier(queries []datastore.Query) (printing_order_supplier.PrintingOrderSupplier, error) {
	entityMap, err := PrintingOrderSupplierRepo.Find(printing_order_supplier.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return printing_order_supplier.PrintingOrderSupplier{}, err
	}
	return printing_order_supplier.Decode(entityMap), err
}

// GetAllPrintingOrderSuppliers retrieves all PrintingOrderSupplier entities.
func GetAllPrintingOrderSuppliers() ([]printing_order_supplier.PrintingOrderSupplier, error) {
	entityMaps, err := PrintingOrderSupplierRepo.ReadAll(printing_order_supplier.GetCollectionPath())
	if err != nil {
		return []printing_order_supplier.PrintingOrderSupplier{}, err
	}
	return printing_order_supplier.DecodeAll(entityMaps), nil
}


// GetAllPrintingOrderSuppliersPaginated retrieves all PrintingOrderSupplier entities in a paginated manner.
func GetAllPrintingOrderSuppliersPaginated(pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, datastore.Pagination, error) {
	entityMaps, pagination, err := PrintingOrderSupplierRepo.ReadPaginated(printing_order_supplier.GetCollectionPath(), pagination)
	if err != nil {
		return []printing_order_supplier.PrintingOrderSupplier{}, pagination, err
	}
	return printing_order_supplier.DecodeAll(entityMaps), pagination, nil
}

// QueryPrintingOrderSuppliers retrieves all PrintingOrderSupplier entities according to given queries.
func QueryPrintingOrderSuppliers(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, error) {
	entityMaps, err := PrintingOrderSupplierRepo.Query(printing_order_supplier.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []printing_order_supplier.PrintingOrderSupplier{}, err
	}
	return printing_order_supplier.DecodeAll(entityMaps), nil
}

// QueryPrintingOrderSuppliersGroup retrieves all PrintingOrderSupplier entities according to given queries.
func QueryPrintingOrderSuppliersGroup(queries []datastore.Query) ([]printing_order_supplier.PrintingOrderSupplier, error) {
	entityMaps, err := PrintingOrderSupplierRepo.QueryGroup("printing_order_suppliers", queries)
	if err != nil {
		return []printing_order_supplier.PrintingOrderSupplier{}, err
	}
	return printing_order_supplier.DecodeAll(entityMaps), nil
}

// CreatePrintingOrderSupplier creates a new PrintingOrderSupplier entity.
func CreatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	return PrintingOrderSupplierRepo.Create(&entity, editorID)
}

// UpdatePrintingOrderSupplier updates an existing PrintingOrderSupplier entity.
func UpdatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	return PrintingOrderSupplierRepo.Update(&entity, editorID)
}

// SavePrintingOrderSupplier saves a PrintingOrderSupplier entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePrintingOrderSupplier(entity, editorID)
	} else {
		return UpdatePrintingOrderSupplier(entity, editorID)
	}
}

// DeletePrintingOrderSupplier deletes a PrintingOrderSupplier entity by its parents' path and printingOrderSupplierID.
func DeletePrintingOrderSupplier(printingOrderSupplierID string) error {
	return PrintingOrderSupplierRepo.Delete(printing_order_supplier.GetDocumentPath(printingOrderSupplierID))
}
