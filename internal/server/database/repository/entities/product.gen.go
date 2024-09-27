package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/product"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ProductRepo = repository.NewRepository[*product.Product]()

// GetProduct retrieves a single Product entity by its ID and productID.
func GetProduct(productID string) (product.Product, error) {
	entityMap, err := ProductRepo.Read(product.GetDocumentPath( productID))
	return product.Decode(entityMap), err
}

// GetProductOrNew retrieves a single Product entity by its ID and productID.
func GetProductOrNew(productID string) (product.Product, error) {
	entityMap, err := ProductRepo.Read(product.GetDocumentPath( productID))
	if err != nil || entityMap == nil {
		return product.Product{}, err
	}
	return product.Decode(entityMap), err
}

// GetProduct retrieves a single Product entity by its document path.
func GetProductByPath(path string) (product.Product, error) {
	entityMap, err := ProductRepo.Read(path)
	return product.Decode(entityMap), err
}

// FindProduct retrieves a Product entity according to given queries.
func FindProduct(queries []datastore.Query) (product.Product, error) {
	entityMap, err := ProductRepo.Find(product.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return product.Product{}, err
	}
	return product.Decode(entityMap), err
}

// GetAllProducts retrieves all Product entities.
func GetAllProducts() ([]product.Product, error) {
	entityMaps, err := ProductRepo.ReadAll(product.GetCollectionPath())
	if err != nil {
		return []product.Product{}, err
	}
	return product.DecodeAll(entityMaps), nil
}


// GetAllProductsPaginated retrieves all Product entities in a paginated manner.
func GetAllProductsPaginated(pagination datastore.Pagination) ([]product.Product, datastore.Pagination, error) {
	entityMaps, pagination, err := ProductRepo.ReadPaginated(product.GetCollectionPath(), pagination)
	if err != nil {
		return []product.Product{}, pagination, err
	}
	return product.DecodeAll(entityMaps), pagination, nil
}

// QueryProducts retrieves all Product entities according to given queries.
func QueryProducts(queries []datastore.Query, pagination datastore.Pagination) ([]product.Product, error) {
	entityMaps, err := ProductRepo.Query(product.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []product.Product{}, err
	}
	return product.DecodeAll(entityMaps), nil
}

// QueryProductsGroup retrieves all Product entities according to given queries.
func QueryProductsGroup(queries []datastore.Query) ([]product.Product, error) {
	entityMaps, err := ProductRepo.QueryGroup("products", queries)
	if err != nil {
		return []product.Product{}, err
	}
	return product.DecodeAll(entityMaps), nil
}

// CreateProduct creates a new Product entity.
func CreateProduct(entity product.Product, editorID *string) (string, error) {
	return ProductRepo.Create(&entity, editorID)
}

// UpdateProduct updates an existing Product entity.
func UpdateProduct(entity product.Product, editorID *string) (string, error) {
	return ProductRepo.Update(&entity, editorID)
}

// SaveProduct saves a Product entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveProduct(entity product.Product, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateProduct(entity, editorID)
	} else {
		return UpdateProduct(entity, editorID)
	}
}

// DeleteProduct deletes a Product entity by its parents' path and productID.
func DeleteProduct(productID string) error {
	return ProductRepo.Delete(product.GetDocumentPath(productID))
}
