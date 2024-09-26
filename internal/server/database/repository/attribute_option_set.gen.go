package repo

import (
	"github.com/bakeable/bkry/data/entities/attribute_option_set"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AttributeOptionSetRepo = repository.NewRepository[*attribute_option_set.AttributeOptionSet]()

// GetAttributeOptionSet retrieves a single AttributeOptionSet entity by its ID and attributeOptionSetID.
func GetAttributeOptionSet(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error) {
	entityMap, err := AttributeOptionSetRepo.Read(attribute_option_set.GetDocumentPath( attributeOptionSetID))
	return attribute_option_set.Decode(entityMap), err
}

// GetAttributeOptionSetOrNew retrieves a single AttributeOptionSet entity by its ID and attributeOptionSetID.
func GetAttributeOptionSetOrNew(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error) {
	entityMap, err := AttributeOptionSetRepo.Read(attribute_option_set.GetDocumentPath( attributeOptionSetID))
	if err != nil || entityMap == nil {
		return attribute_option_set.AttributeOptionSet{}, err
	}
	return attribute_option_set.Decode(entityMap), err
}

// GetAttributeOptionSet retrieves a single AttributeOptionSet entity by its document path.
func GetAttributeOptionSetByPath(path string) (attribute_option_set.AttributeOptionSet, error) {
	entityMap, err := AttributeOptionSetRepo.Read(path)
	return attribute_option_set.Decode(entityMap), err
}

// FindAttributeOptionSet retrieves a AttributeOptionSet entity according to given queries.
func FindAttributeOptionSet(queries []datastore.Query) (attribute_option_set.AttributeOptionSet, error) {
	entityMap, err := AttributeOptionSetRepo.Find(attribute_option_set.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return attribute_option_set.AttributeOptionSet{}, err
	}
	return attribute_option_set.Decode(entityMap), err
}

// GetAllAttributeOptionSets retrieves all AttributeOptionSet entities.
func GetAllAttributeOptionSets() ([]attribute_option_set.AttributeOptionSet, error) {
	entityMaps, err := AttributeOptionSetRepo.ReadAll(attribute_option_set.GetCollectionPath())
	if err != nil {
		return []attribute_option_set.AttributeOptionSet{}, err
	}
	return attribute_option_set.DecodeAll(entityMaps), nil
}


// GetAllAttributeOptionSetsPaginated retrieves all AttributeOptionSet entities in a paginated manner.
func GetAllAttributeOptionSetsPaginated(pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, datastore.Pagination, error) {
	entityMaps, pagination, err := AttributeOptionSetRepo.ReadPaginated(attribute_option_set.GetCollectionPath(), pagination)
	if err != nil {
		return []attribute_option_set.AttributeOptionSet{}, pagination, err
	}
	return attribute_option_set.DecodeAll(entityMaps), pagination, nil
}

// QueryAttributeOptionSets retrieves all AttributeOptionSet entities according to given queries.
func QueryAttributeOptionSets(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, error) {
	entityMaps, err := AttributeOptionSetRepo.Query(attribute_option_set.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []attribute_option_set.AttributeOptionSet{}, err
	}
	return attribute_option_set.DecodeAll(entityMaps), nil
}

// QueryAttributeOptionSetsGroup retrieves all AttributeOptionSet entities according to given queries.
func QueryAttributeOptionSetsGroup(queries []datastore.Query) ([]attribute_option_set.AttributeOptionSet, error) {
	entityMaps, err := AttributeOptionSetRepo.QueryGroup("attribute_option_sets", queries)
	if err != nil {
		return []attribute_option_set.AttributeOptionSet{}, err
	}
	return attribute_option_set.DecodeAll(entityMaps), nil
}

// CreateAttributeOptionSet creates a new AttributeOptionSet entity.
func CreateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	return AttributeOptionSetRepo.Create(&entity, editorID)
}

// UpdateAttributeOptionSet updates an existing AttributeOptionSet entity.
func UpdateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	return AttributeOptionSetRepo.Update(&entity, editorID)
}

// SaveAttributeOptionSet saves a AttributeOptionSet entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAttributeOptionSet(entity, editorID)
	} else {
		return UpdateAttributeOptionSet(entity, editorID)
	}
}

// DeleteAttributeOptionSet deletes a AttributeOptionSet entity by its parents' path and attributeOptionSetID.
func DeleteAttributeOptionSet(attributeOptionSetID string) error {
	return AttributeOptionSetRepo.Delete(attribute_option_set.GetDocumentPath(attributeOptionSetID))
}
