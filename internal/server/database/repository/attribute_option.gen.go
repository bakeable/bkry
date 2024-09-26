package repo

import (
	"github.com/bakeable/bkry/data/entities/attribute_option"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AttributeOptionRepo = repository.NewRepository[*attribute_option.AttributeOption]()

// GetAttributeOption retrieves a single AttributeOption entity by its ID and attributeOptionID.
func GetAttributeOption(attributeOptionID string) (attribute_option.AttributeOption, error) {
	entityMap, err := AttributeOptionRepo.Read(attribute_option.GetDocumentPath( attributeOptionID))
	return attribute_option.Decode(entityMap), err
}

// GetAttributeOptionOrNew retrieves a single AttributeOption entity by its ID and attributeOptionID.
func GetAttributeOptionOrNew(attributeOptionID string) (attribute_option.AttributeOption, error) {
	entityMap, err := AttributeOptionRepo.Read(attribute_option.GetDocumentPath( attributeOptionID))
	if err != nil || entityMap == nil {
		return attribute_option.AttributeOption{}, err
	}
	return attribute_option.Decode(entityMap), err
}

// GetAttributeOption retrieves a single AttributeOption entity by its document path.
func GetAttributeOptionByPath(path string) (attribute_option.AttributeOption, error) {
	entityMap, err := AttributeOptionRepo.Read(path)
	return attribute_option.Decode(entityMap), err
}

// FindAttributeOption retrieves a AttributeOption entity according to given queries.
func FindAttributeOption(queries []datastore.Query) (attribute_option.AttributeOption, error) {
	entityMap, err := AttributeOptionRepo.Find(attribute_option.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return attribute_option.AttributeOption{}, err
	}
	return attribute_option.Decode(entityMap), err
}

// GetAllAttributeOptions retrieves all AttributeOption entities.
func GetAllAttributeOptions() ([]attribute_option.AttributeOption, error) {
	entityMaps, err := AttributeOptionRepo.ReadAll(attribute_option.GetCollectionPath())
	if err != nil {
		return []attribute_option.AttributeOption{}, err
	}
	return attribute_option.DecodeAll(entityMaps), nil
}


// GetAllAttributeOptionsPaginated retrieves all AttributeOption entities in a paginated manner.
func GetAllAttributeOptionsPaginated(pagination datastore.Pagination) ([]attribute_option.AttributeOption, datastore.Pagination, error) {
	entityMaps, pagination, err := AttributeOptionRepo.ReadPaginated(attribute_option.GetCollectionPath(), pagination)
	if err != nil {
		return []attribute_option.AttributeOption{}, pagination, err
	}
	return attribute_option.DecodeAll(entityMaps), pagination, nil
}

// QueryAttributeOptions retrieves all AttributeOption entities according to given queries.
func QueryAttributeOptions(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option.AttributeOption, error) {
	entityMaps, err := AttributeOptionRepo.Query(attribute_option.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []attribute_option.AttributeOption{}, err
	}
	return attribute_option.DecodeAll(entityMaps), nil
}

// QueryAttributeOptionsGroup retrieves all AttributeOption entities according to given queries.
func QueryAttributeOptionsGroup(queries []datastore.Query) ([]attribute_option.AttributeOption, error) {
	entityMaps, err := AttributeOptionRepo.QueryGroup("attribute_options", queries)
	if err != nil {
		return []attribute_option.AttributeOption{}, err
	}
	return attribute_option.DecodeAll(entityMaps), nil
}

// CreateAttributeOption creates a new AttributeOption entity.
func CreateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	return AttributeOptionRepo.Create(&entity, editorID)
}

// UpdateAttributeOption updates an existing AttributeOption entity.
func UpdateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	return AttributeOptionRepo.Update(&entity, editorID)
}

// SaveAttributeOption saves a AttributeOption entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAttributeOption(entity, editorID)
	} else {
		return UpdateAttributeOption(entity, editorID)
	}
}

// DeleteAttributeOption deletes a AttributeOption entity by its parents' path and attributeOptionID.
func DeleteAttributeOption(attributeOptionID string) error {
	return AttributeOptionRepo.Delete(attribute_option.GetDocumentPath(attributeOptionID))
}
