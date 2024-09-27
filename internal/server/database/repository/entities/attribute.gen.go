package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/attribute"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AttributeRepo = repository.NewRepository[*attribute.Attribute]()

// GetAttribute retrieves a single Attribute entity by its ID and attributeID.
func GetAttribute(attributeID string) (attribute.Attribute, error) {
	entityMap, err := AttributeRepo.Read(attribute.GetDocumentPath( attributeID))
	return attribute.Decode(entityMap), err
}

// GetAttributeOrNew retrieves a single Attribute entity by its ID and attributeID.
func GetAttributeOrNew(attributeID string) (attribute.Attribute, error) {
	entityMap, err := AttributeRepo.Read(attribute.GetDocumentPath( attributeID))
	if err != nil || entityMap == nil {
		return attribute.Attribute{}, err
	}
	return attribute.Decode(entityMap), err
}

// GetAttribute retrieves a single Attribute entity by its document path.
func GetAttributeByPath(path string) (attribute.Attribute, error) {
	entityMap, err := AttributeRepo.Read(path)
	return attribute.Decode(entityMap), err
}

// FindAttribute retrieves a Attribute entity according to given queries.
func FindAttribute(queries []datastore.Query) (attribute.Attribute, error) {
	entityMap, err := AttributeRepo.Find(attribute.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return attribute.Attribute{}, err
	}
	return attribute.Decode(entityMap), err
}

// GetAllAttributes retrieves all Attribute entities.
func GetAllAttributes() ([]attribute.Attribute, error) {
	entityMaps, err := AttributeRepo.ReadAll(attribute.GetCollectionPath())
	if err != nil {
		return []attribute.Attribute{}, err
	}
	return attribute.DecodeAll(entityMaps), nil
}


// GetAllAttributesPaginated retrieves all Attribute entities in a paginated manner.
func GetAllAttributesPaginated(pagination datastore.Pagination) ([]attribute.Attribute, datastore.Pagination, error) {
	entityMaps, pagination, err := AttributeRepo.ReadPaginated(attribute.GetCollectionPath(), pagination)
	if err != nil {
		return []attribute.Attribute{}, pagination, err
	}
	return attribute.DecodeAll(entityMaps), pagination, nil
}

// QueryAttributes retrieves all Attribute entities according to given queries.
func QueryAttributes(queries []datastore.Query, pagination datastore.Pagination) ([]attribute.Attribute, error) {
	entityMaps, err := AttributeRepo.Query(attribute.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []attribute.Attribute{}, err
	}
	return attribute.DecodeAll(entityMaps), nil
}

// QueryAttributesGroup retrieves all Attribute entities according to given queries.
func QueryAttributesGroup(queries []datastore.Query) ([]attribute.Attribute, error) {
	entityMaps, err := AttributeRepo.QueryGroup("attributes", queries)
	if err != nil {
		return []attribute.Attribute{}, err
	}
	return attribute.DecodeAll(entityMaps), nil
}

// CreateAttribute creates a new Attribute entity.
func CreateAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	return AttributeRepo.Create(&entity, editorID)
}

// UpdateAttribute updates an existing Attribute entity.
func UpdateAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	return AttributeRepo.Update(&entity, editorID)
}

// SaveAttribute saves a Attribute entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAttribute(entity, editorID)
	} else {
		return UpdateAttribute(entity, editorID)
	}
}

// DeleteAttribute deletes a Attribute entity by its parents' path and attributeID.
func DeleteAttribute(attributeID string) error {
	return AttributeRepo.Delete(attribute.GetDocumentPath(attributeID))
}
