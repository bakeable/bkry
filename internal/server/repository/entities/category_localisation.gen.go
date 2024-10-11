package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var CategoryLocalisationRepo = repository.NewFirestoreRepository[*category_localisation.CategoryLocalisation]()

// GetCategoryLocalisation retrieves a single CategoryLocalisation entity by its ID and categoryLocalisationID.
func GetCategoryLocalisation(CategoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error) {
	entityMap, err := CategoryLocalisationRepo.Read(category_localisation.GetDocumentPath(CategoryID,  categoryLocalisationID))
	return category_localisation.Decode(entityMap), err
}

// GetCategoryLocalisationOrNew retrieves a single CategoryLocalisation entity by its ID and categoryLocalisationID.
func GetCategoryLocalisationOrNew(CategoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error) {
	entityMap, err := CategoryLocalisationRepo.Read(category_localisation.GetDocumentPath(CategoryID,  categoryLocalisationID))
	if err != nil || entityMap == nil {
		return category_localisation.CategoryLocalisation{}, err
	}
	return category_localisation.Decode(entityMap), err
}

// GetCategoryLocalisation retrieves a single CategoryLocalisation entity by its document path.
func GetCategoryLocalisationByPath(path string) (category_localisation.CategoryLocalisation, error) {
	entityMap, err := CategoryLocalisationRepo.Read(path)
	return category_localisation.Decode(entityMap), err
}

// FindCategoryLocalisation retrieves a CategoryLocalisation entity according to given queries.
func FindCategoryLocalisation(categoryID string, queries []database.Query) (category_localisation.CategoryLocalisation, error) {
	entityMap, err := CategoryLocalisationRepo.Find(category_localisation.GetCollectionPath(categoryID), queries)
	if err != nil || entityMap == nil {
		return category_localisation.CategoryLocalisation{}, err
	}
	return category_localisation.Decode(entityMap), err
}

// GetAllCategoryLocalisations retrieves all CategoryLocalisation entities.
func GetAllCategoryLocalisations(categoryID string) ([]category_localisation.CategoryLocalisation, error) {
	entityMaps, err := CategoryLocalisationRepo.ReadAll(category_localisation.GetCollectionPath(categoryID))
	if err != nil {
		return []category_localisation.CategoryLocalisation{}, err
	}
	return category_localisation.DecodeAll(entityMaps), nil
}


// GetAllCategoryLocalisationsPaginated retrieves all CategoryLocalisation entities in a paginated manner.
func GetAllCategoryLocalisationsPaginated(categoryID string, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := CategoryLocalisationRepo.ReadPaginated(category_localisation.GetCollectionPath(categoryID), pagination)
	if err != nil {
		return []category_localisation.CategoryLocalisation{}, pagination, err
	}
	return category_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryCategoryLocalisations retrieves all CategoryLocalisation entities according to given queries.
func QueryCategoryLocalisations(categoryID string, queries []database.Query, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, error) {
	entityMaps, err := CategoryLocalisationRepo.Query(category_localisation.GetCollectionPath(categoryID), queries, pagination)
	if err != nil {
		return []category_localisation.CategoryLocalisation{}, err
	}
	return category_localisation.DecodeAll(entityMaps), nil
}

// QueryCategoryLocalisationsGroup retrieves all CategoryLocalisation entities according to given queries.
func QueryCategoryLocalisationsGroup(queries []database.Query) ([]category_localisation.CategoryLocalisation, error) {
	entityMaps, err := CategoryLocalisationRepo.QueryGroup("category_localisations", queries)
	if err != nil {
		return []category_localisation.CategoryLocalisation{}, err
	}
	return category_localisation.DecodeAll(entityMaps), nil
}

// CreateCategoryLocalisation creates a new CategoryLocalisation entity.
func CreateCategoryLocalisation(CategoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	entity.CategoryID = CategoryID
	return CategoryLocalisationRepo.Create(&entity, editorID)
}

// UpdateCategoryLocalisation updates an existing CategoryLocalisation entity.
func UpdateCategoryLocalisation(CategoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	entity.CategoryID = CategoryID
	return CategoryLocalisationRepo.Update(&entity, editorID)
}

// SaveCategoryLocalisation saves a CategoryLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveCategoryLocalisation(CategoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateCategoryLocalisation(CategoryID, entity, editorID)
	} else {
		return UpdateCategoryLocalisation(CategoryID, entity, editorID)
	}
}

// DeleteCategoryLocalisation deletes a CategoryLocalisation entity by its parents' path and categoryLocalisationID.
func DeleteCategoryLocalisation(CategoryID string, categoryLocalisationID string) error {
	return CategoryLocalisationRepo.Delete(category_localisation.GetDocumentPath(CategoryID, categoryLocalisationID))
}
