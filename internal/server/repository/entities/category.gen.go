package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/category"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var CategoryRepo = repository.NewFirestoreRepository[*category.Category]()

// GetCategory retrieves a single Category entity by its ID and categoryID.
func GetCategory(categoryID string) (category.Category, error) {
	entityMap, err := CategoryRepo.Read(category.GetDocumentPath( categoryID))
	return category.Decode(entityMap), err
}

// GetCategoryOrNew retrieves a single Category entity by its ID and categoryID.
func GetCategoryOrNew(categoryID string) (category.Category, error) {
	entityMap, err := CategoryRepo.Read(category.GetDocumentPath( categoryID))
	if err != nil || entityMap == nil {
		return category.Category{}, err
	}
	return category.Decode(entityMap), err
}

// GetCategory retrieves a single Category entity by its document path.
func GetCategoryByPath(path string) (category.Category, error) {
	entityMap, err := CategoryRepo.Read(path)
	return category.Decode(entityMap), err
}

// FindCategory retrieves a Category entity according to given queries.
func FindCategory(queries []database.Query) (category.Category, error) {
	entityMap, err := CategoryRepo.Find(category.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return category.Category{}, err
	}
	return category.Decode(entityMap), err
}

// GetAllCategories retrieves all Category entities.
func GetAllCategories() ([]category.Category, error) {
	entityMaps, err := CategoryRepo.ReadAll(category.GetCollectionPath())
	if err != nil {
		return []category.Category{}, err
	}
	return category.DecodeAll(entityMaps), nil
}


// GetAllCategoriesPaginated retrieves all Category entities in a paginated manner.
func GetAllCategoriesPaginated(pagination database.Pagination) ([]category.Category, database.Pagination, error) {
	entityMaps, pagination, err := CategoryRepo.ReadPaginated(category.GetCollectionPath(), pagination)
	if err != nil {
		return []category.Category{}, pagination, err
	}
	return category.DecodeAll(entityMaps), pagination, nil
}

// QueryCategories retrieves all Category entities according to given queries.
func QueryCategories(queries []database.Query, pagination database.Pagination) ([]category.Category, error) {
	entityMaps, err := CategoryRepo.Query(category.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []category.Category{}, err
	}
	return category.DecodeAll(entityMaps), nil
}

// QueryCategoriesGroup retrieves all Category entities according to given queries.
func QueryCategoriesGroup(queries []database.Query) ([]category.Category, error) {
	entityMaps, err := CategoryRepo.QueryGroup("categories", queries)
	if err != nil {
		return []category.Category{}, err
	}
	return category.DecodeAll(entityMaps), nil
}

// CreateCategory creates a new Category entity.
func CreateCategory(entity category.Category, editorID *string) (string, error) {
	return CategoryRepo.Create(&entity, editorID)
}

// UpdateCategory updates an existing Category entity.
func UpdateCategory(entity category.Category, editorID *string) (string, error) {
	return CategoryRepo.Update(&entity, editorID)
}

// SaveCategory saves a Category entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveCategory(entity category.Category, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateCategory(entity, editorID)
	} else {
		return UpdateCategory(entity, editorID)
	}
}

// DeleteCategory deletes a Category entity by its parents' path and categoryID.
func DeleteCategory(categoryID string) error {
	return CategoryRepo.Delete(category.GetDocumentPath(categoryID))
}
