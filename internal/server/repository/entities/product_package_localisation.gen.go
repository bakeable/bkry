package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/product_package_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ProductPackageLocalisationRepo = repository.NewFirestoreRepository[*product_package_localisation.ProductPackageLocalisation]()

// GetProductPackageLocalisation retrieves a single ProductPackageLocalisation entity by its ID and productPackageLocalisationID.
func GetProductPackageLocalisation(ProductPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error) {
	entityMap, err := ProductPackageLocalisationRepo.Read(product_package_localisation.GetDocumentPath(ProductPackageID,  productPackageLocalisationID))
	return product_package_localisation.Decode(entityMap), err
}

// GetProductPackageLocalisationOrNew retrieves a single ProductPackageLocalisation entity by its ID and productPackageLocalisationID.
func GetProductPackageLocalisationOrNew(ProductPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error) {
	entityMap, err := ProductPackageLocalisationRepo.Read(product_package_localisation.GetDocumentPath(ProductPackageID,  productPackageLocalisationID))
	if err != nil || entityMap == nil {
		return product_package_localisation.ProductPackageLocalisation{}, err
	}
	return product_package_localisation.Decode(entityMap), err
}

// GetProductPackageLocalisation retrieves a single ProductPackageLocalisation entity by its document path.
func GetProductPackageLocalisationByPath(path string) (product_package_localisation.ProductPackageLocalisation, error) {
	entityMap, err := ProductPackageLocalisationRepo.Read(path)
	return product_package_localisation.Decode(entityMap), err
}

// FindProductPackageLocalisation retrieves a ProductPackageLocalisation entity according to given queries.
func FindProductPackageLocalisation(productPackageID string, queries []database.Query) (product_package_localisation.ProductPackageLocalisation, error) {
	entityMap, err := ProductPackageLocalisationRepo.Find(product_package_localisation.GetCollectionPath(productPackageID), queries)
	if err != nil || entityMap == nil {
		return product_package_localisation.ProductPackageLocalisation{}, err
	}
	return product_package_localisation.Decode(entityMap), err
}

// GetAllProductPackageLocalisations retrieves all ProductPackageLocalisation entities.
func GetAllProductPackageLocalisations(productPackageID string) ([]product_package_localisation.ProductPackageLocalisation, error) {
	entityMaps, err := ProductPackageLocalisationRepo.ReadAll(product_package_localisation.GetCollectionPath(productPackageID))
	if err != nil {
		return []product_package_localisation.ProductPackageLocalisation{}, err
	}
	return product_package_localisation.DecodeAll(entityMaps), nil
}


// GetAllProductPackageLocalisationsPaginated retrieves all ProductPackageLocalisation entities in a paginated manner.
func GetAllProductPackageLocalisationsPaginated(productPackageID string, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := ProductPackageLocalisationRepo.ReadPaginated(product_package_localisation.GetCollectionPath(productPackageID), pagination)
	if err != nil {
		return []product_package_localisation.ProductPackageLocalisation{}, pagination, err
	}
	return product_package_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryProductPackageLocalisations retrieves all ProductPackageLocalisation entities according to given queries.
func QueryProductPackageLocalisations(productPackageID string, queries []database.Query, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, error) {
	entityMaps, err := ProductPackageLocalisationRepo.Query(product_package_localisation.GetCollectionPath(productPackageID), queries, pagination)
	if err != nil {
		return []product_package_localisation.ProductPackageLocalisation{}, err
	}
	return product_package_localisation.DecodeAll(entityMaps), nil
}

// QueryProductPackageLocalisationsGroup retrieves all ProductPackageLocalisation entities according to given queries.
func QueryProductPackageLocalisationsGroup(queries []database.Query) ([]product_package_localisation.ProductPackageLocalisation, error) {
	entityMaps, err := ProductPackageLocalisationRepo.QueryGroup("product_package_localisations", queries)
	if err != nil {
		return []product_package_localisation.ProductPackageLocalisation{}, err
	}
	return product_package_localisation.DecodeAll(entityMaps), nil
}

// CreateProductPackageLocalisation creates a new ProductPackageLocalisation entity.
func CreateProductPackageLocalisation(ProductPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	entity.ProductPackageID = ProductPackageID
	return ProductPackageLocalisationRepo.Create(&entity, editorID)
}

// UpdateProductPackageLocalisation updates an existing ProductPackageLocalisation entity.
func UpdateProductPackageLocalisation(ProductPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	entity.ProductPackageID = ProductPackageID
	return ProductPackageLocalisationRepo.Update(&entity, editorID)
}

// SaveProductPackageLocalisation saves a ProductPackageLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveProductPackageLocalisation(ProductPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateProductPackageLocalisation(ProductPackageID, entity, editorID)
	} else {
		return UpdateProductPackageLocalisation(ProductPackageID, entity, editorID)
	}
}

// DeleteProductPackageLocalisation deletes a ProductPackageLocalisation entity by its parents' path and productPackageLocalisationID.
func DeleteProductPackageLocalisation(ProductPackageID string, productPackageLocalisationID string) error {
	return ProductPackageLocalisationRepo.Delete(product_package_localisation.GetDocumentPath(ProductPackageID, productPackageLocalisationID))
}
