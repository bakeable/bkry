package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/product_package"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ProductPackageRepo = repository.NewFirestoreRepository[*product_package.ProductPackage]()

// GetProductPackage retrieves a single ProductPackage entity by its ID and productPackageID.
func GetProductPackage(productPackageID string) (product_package.ProductPackage, error) {
	entityMap, err := ProductPackageRepo.Read(product_package.GetDocumentPath( productPackageID))
	return product_package.Decode(entityMap), err
}

// GetProductPackageOrNew retrieves a single ProductPackage entity by its ID and productPackageID.
func GetProductPackageOrNew(productPackageID string) (product_package.ProductPackage, error) {
	entityMap, err := ProductPackageRepo.Read(product_package.GetDocumentPath( productPackageID))
	if err != nil || entityMap == nil {
		return product_package.ProductPackage{}, err
	}
	return product_package.Decode(entityMap), err
}

// GetProductPackage retrieves a single ProductPackage entity by its document path.
func GetProductPackageByPath(path string) (product_package.ProductPackage, error) {
	entityMap, err := ProductPackageRepo.Read(path)
	return product_package.Decode(entityMap), err
}

// FindProductPackage retrieves a ProductPackage entity according to given queries.
func FindProductPackage(queries []database.Query) (product_package.ProductPackage, error) {
	entityMap, err := ProductPackageRepo.Find(product_package.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return product_package.ProductPackage{}, err
	}
	return product_package.Decode(entityMap), err
}

// GetAllProductPackages retrieves all ProductPackage entities.
func GetAllProductPackages() ([]product_package.ProductPackage, error) {
	entityMaps, err := ProductPackageRepo.ReadAll(product_package.GetCollectionPath())
	if err != nil {
		return []product_package.ProductPackage{}, err
	}
	return product_package.DecodeAll(entityMaps), nil
}


// GetAllProductPackagesPaginated retrieves all ProductPackage entities in a paginated manner.
func GetAllProductPackagesPaginated(pagination database.Pagination) ([]product_package.ProductPackage, database.Pagination, error) {
	entityMaps, pagination, err := ProductPackageRepo.ReadPaginated(product_package.GetCollectionPath(), pagination)
	if err != nil {
		return []product_package.ProductPackage{}, pagination, err
	}
	return product_package.DecodeAll(entityMaps), pagination, nil
}

// QueryProductPackages retrieves all ProductPackage entities according to given queries.
func QueryProductPackages(queries []database.Query, pagination database.Pagination) ([]product_package.ProductPackage, error) {
	entityMaps, err := ProductPackageRepo.Query(product_package.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []product_package.ProductPackage{}, err
	}
	return product_package.DecodeAll(entityMaps), nil
}

// QueryProductPackagesGroup retrieves all ProductPackage entities according to given queries.
func QueryProductPackagesGroup(queries []database.Query) ([]product_package.ProductPackage, error) {
	entityMaps, err := ProductPackageRepo.QueryGroup("product_packages", queries)
	if err != nil {
		return []product_package.ProductPackage{}, err
	}
	return product_package.DecodeAll(entityMaps), nil
}

// CreateProductPackage creates a new ProductPackage entity.
func CreateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	return ProductPackageRepo.Create(&entity, editorID)
}

// UpdateProductPackage updates an existing ProductPackage entity.
func UpdateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	return ProductPackageRepo.Update(&entity, editorID)
}

// SaveProductPackage saves a ProductPackage entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateProductPackage(entity, editorID)
	} else {
		return UpdateProductPackage(entity, editorID)
	}
}

// DeleteProductPackage deletes a ProductPackage entity by its parents' path and productPackageID.
func DeleteProductPackage(productPackageID string) error {
	return ProductPackageRepo.Delete(product_package.GetDocumentPath(productPackageID))
}
