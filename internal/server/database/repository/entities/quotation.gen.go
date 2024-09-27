package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/quotation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var QuotationRepo = repository.NewRepository[*quotation.Quotation]()

// GetQuotation retrieves a single Quotation entity by its ID and quotationID.
func GetQuotation(quotationID string) (quotation.Quotation, error) {
	entityMap, err := QuotationRepo.Read(quotation.GetDocumentPath( quotationID))
	return quotation.Decode(entityMap), err
}

// GetQuotationOrNew retrieves a single Quotation entity by its ID and quotationID.
func GetQuotationOrNew(quotationID string) (quotation.Quotation, error) {
	entityMap, err := QuotationRepo.Read(quotation.GetDocumentPath( quotationID))
	if err != nil || entityMap == nil {
		return quotation.Quotation{}, err
	}
	return quotation.Decode(entityMap), err
}

// GetQuotation retrieves a single Quotation entity by its document path.
func GetQuotationByPath(path string) (quotation.Quotation, error) {
	entityMap, err := QuotationRepo.Read(path)
	return quotation.Decode(entityMap), err
}

// FindQuotation retrieves a Quotation entity according to given queries.
func FindQuotation(queries []datastore.Query) (quotation.Quotation, error) {
	entityMap, err := QuotationRepo.Find(quotation.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return quotation.Quotation{}, err
	}
	return quotation.Decode(entityMap), err
}

// GetAllQuotations retrieves all Quotation entities.
func GetAllQuotations() ([]quotation.Quotation, error) {
	entityMaps, err := QuotationRepo.ReadAll(quotation.GetCollectionPath())
	if err != nil {
		return []quotation.Quotation{}, err
	}
	return quotation.DecodeAll(entityMaps), nil
}


// GetAllQuotationsPaginated retrieves all Quotation entities in a paginated manner.
func GetAllQuotationsPaginated(pagination datastore.Pagination) ([]quotation.Quotation, datastore.Pagination, error) {
	entityMaps, pagination, err := QuotationRepo.ReadPaginated(quotation.GetCollectionPath(), pagination)
	if err != nil {
		return []quotation.Quotation{}, pagination, err
	}
	return quotation.DecodeAll(entityMaps), pagination, nil
}

// QueryQuotations retrieves all Quotation entities according to given queries.
func QueryQuotations(queries []datastore.Query, pagination datastore.Pagination) ([]quotation.Quotation, error) {
	entityMaps, err := QuotationRepo.Query(quotation.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []quotation.Quotation{}, err
	}
	return quotation.DecodeAll(entityMaps), nil
}

// QueryQuotationsGroup retrieves all Quotation entities according to given queries.
func QueryQuotationsGroup(queries []datastore.Query) ([]quotation.Quotation, error) {
	entityMaps, err := QuotationRepo.QueryGroup("quotations", queries)
	if err != nil {
		return []quotation.Quotation{}, err
	}
	return quotation.DecodeAll(entityMaps), nil
}

// CreateQuotation creates a new Quotation entity.
func CreateQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	return QuotationRepo.Create(&entity, editorID)
}

// UpdateQuotation updates an existing Quotation entity.
func UpdateQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	return QuotationRepo.Update(&entity, editorID)
}

// SaveQuotation saves a Quotation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateQuotation(entity, editorID)
	} else {
		return UpdateQuotation(entity, editorID)
	}
}

// DeleteQuotation deletes a Quotation entity by its parents' path and quotationID.
func DeleteQuotation(quotationID string) error {
	return QuotationRepo.Delete(quotation.GetDocumentPath(quotationID))
}
