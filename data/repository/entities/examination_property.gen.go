package repo

import (
	"github.com/bakeable/bkry/data/entities/examination_property"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ExaminationPropertyRepo = repository.NewRepository[*examination_property.ExaminationProperty]()

// GetExaminationProperty retrieves a single ExaminationProperty entity by its ID and examinationPropertyID.
func GetExaminationProperty(examinationPropertyID string) (examination_property.ExaminationProperty, error) {
	entityMap, err := ExaminationPropertyRepo.Read(examination_property.GetDocumentPath( examinationPropertyID))
	return examination_property.Decode(entityMap), err
}

// GetExaminationPropertyOrNew retrieves a single ExaminationProperty entity by its ID and examinationPropertyID.
func GetExaminationPropertyOrNew(examinationPropertyID string) (examination_property.ExaminationProperty, error) {
	entityMap, err := ExaminationPropertyRepo.Read(examination_property.GetDocumentPath( examinationPropertyID))
	if err != nil || entityMap == nil {
		return examination_property.ExaminationProperty{}, err
	}
	return examination_property.Decode(entityMap), err
}

// GetExaminationProperty retrieves a single ExaminationProperty entity by its document path.
func GetExaminationPropertyByPath(path string) (examination_property.ExaminationProperty, error) {
	entityMap, err := ExaminationPropertyRepo.Read(path)
	return examination_property.Decode(entityMap), err
}

// FindExaminationProperty retrieves a ExaminationProperty entity according to given queries.
func FindExaminationProperty(queries []datastore.Query) (examination_property.ExaminationProperty, error) {
	entityMap, err := ExaminationPropertyRepo.Find(examination_property.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return examination_property.ExaminationProperty{}, err
	}
	return examination_property.Decode(entityMap), err
}

// GetAllExaminationProperties retrieves all ExaminationProperty entities.
func GetAllExaminationProperties() ([]examination_property.ExaminationProperty, error) {
	entityMaps, err := ExaminationPropertyRepo.ReadAll(examination_property.GetCollectionPath())
	if err != nil {
		return []examination_property.ExaminationProperty{}, err
	}
	return examination_property.DecodeAll(entityMaps), nil
}


// GetAllExaminationPropertiesPaginated retrieves all ExaminationProperty entities in a paginated manner.
func GetAllExaminationPropertiesPaginated(pagination datastore.Pagination) ([]examination_property.ExaminationProperty, datastore.Pagination, error) {
	entityMaps, pagination, err := ExaminationPropertyRepo.ReadPaginated(examination_property.GetCollectionPath(), pagination)
	if err != nil {
		return []examination_property.ExaminationProperty{}, pagination, err
	}
	return examination_property.DecodeAll(entityMaps), pagination, nil
}

// QueryExaminationProperties retrieves all ExaminationProperty entities according to given queries.
func QueryExaminationProperties(queries []datastore.Query, pagination datastore.Pagination) ([]examination_property.ExaminationProperty, error) {
	entityMaps, err := ExaminationPropertyRepo.Query(examination_property.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []examination_property.ExaminationProperty{}, err
	}
	return examination_property.DecodeAll(entityMaps), nil
}

// QueryExaminationPropertiesGroup retrieves all ExaminationProperty entities according to given queries.
func QueryExaminationPropertiesGroup(queries []datastore.Query) ([]examination_property.ExaminationProperty, error) {
	entityMaps, err := ExaminationPropertyRepo.QueryGroup("examination_properties", queries)
	if err != nil {
		return []examination_property.ExaminationProperty{}, err
	}
	return examination_property.DecodeAll(entityMaps), nil
}

// CreateExaminationProperty creates a new ExaminationProperty entity.
func CreateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	return ExaminationPropertyRepo.Create(&entity, editorID)
}

// UpdateExaminationProperty updates an existing ExaminationProperty entity.
func UpdateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	return ExaminationPropertyRepo.Update(&entity, editorID)
}

// SaveExaminationProperty saves a ExaminationProperty entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateExaminationProperty(entity, editorID)
	} else {
		return UpdateExaminationProperty(entity, editorID)
	}
}

// DeleteExaminationProperty deletes a ExaminationProperty entity by its parents' path and examinationPropertyID.
func DeleteExaminationProperty(examinationPropertyID string) error {
	return ExaminationPropertyRepo.Delete(examination_property.GetDocumentPath(examinationPropertyID))
}
