package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/examination_settings"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ExaminationSettingsRepo = repository.NewRepository[*examination_settings.ExaminationSettings]()

// GetExaminationSettings retrieves a single ExaminationSettings entity by its ID and examinationSettingsID.
func GetExaminationSettings(examinationSettingsID string) (examination_settings.ExaminationSettings, error) {
	entityMap, err := ExaminationSettingsRepo.Read(examination_settings.GetDocumentPath( examinationSettingsID))
	return examination_settings.Decode(entityMap), err
}

// GetExaminationSettingsOrNew retrieves a single ExaminationSettings entity by its ID and examinationSettingsID.
func GetExaminationSettingsOrNew(examinationSettingsID string) (examination_settings.ExaminationSettings, error) {
	entityMap, err := ExaminationSettingsRepo.Read(examination_settings.GetDocumentPath( examinationSettingsID))
	if err != nil || entityMap == nil {
		return examination_settings.ExaminationSettings{}, err
	}
	return examination_settings.Decode(entityMap), err
}

// GetExaminationSettings retrieves a single ExaminationSettings entity by its document path.
func GetExaminationSettingsByPath(path string) (examination_settings.ExaminationSettings, error) {
	entityMap, err := ExaminationSettingsRepo.Read(path)
	return examination_settings.Decode(entityMap), err
}

// FindExaminationSettings retrieves a ExaminationSettings entity according to given queries.
func FindExaminationSettings(queries []datastore.Query) (examination_settings.ExaminationSettings, error) {
	entityMap, err := ExaminationSettingsRepo.Find(examination_settings.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return examination_settings.ExaminationSettings{}, err
	}
	return examination_settings.Decode(entityMap), err
}

// GetAllExaminationSettings retrieves all ExaminationSettings entities.
func GetAllExaminationSettings() ([]examination_settings.ExaminationSettings, error) {
	entityMaps, err := ExaminationSettingsRepo.ReadAll(examination_settings.GetCollectionPath())
	if err != nil {
		return []examination_settings.ExaminationSettings{}, err
	}
	return examination_settings.DecodeAll(entityMaps), nil
}


// GetAllExaminationSettingsPaginated retrieves all ExaminationSettings entities in a paginated manner.
func GetAllExaminationSettingsPaginated(pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, datastore.Pagination, error) {
	entityMaps, pagination, err := ExaminationSettingsRepo.ReadPaginated(examination_settings.GetCollectionPath(), pagination)
	if err != nil {
		return []examination_settings.ExaminationSettings{}, pagination, err
	}
	return examination_settings.DecodeAll(entityMaps), pagination, nil
}

// QueryExaminationSettings retrieves all ExaminationSettings entities according to given queries.
func QueryExaminationSettings(queries []datastore.Query, pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, error) {
	entityMaps, err := ExaminationSettingsRepo.Query(examination_settings.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []examination_settings.ExaminationSettings{}, err
	}
	return examination_settings.DecodeAll(entityMaps), nil
}

// QueryExaminationSettingsGroup retrieves all ExaminationSettings entities according to given queries.
func QueryExaminationSettingsGroup(queries []datastore.Query) ([]examination_settings.ExaminationSettings, error) {
	entityMaps, err := ExaminationSettingsRepo.QueryGroup("examination_settings", queries)
	if err != nil {
		return []examination_settings.ExaminationSettings{}, err
	}
	return examination_settings.DecodeAll(entityMaps), nil
}

// CreateExaminationSettings creates a new ExaminationSettings entity.
func CreateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	return ExaminationSettingsRepo.Create(&entity, editorID)
}

// UpdateExaminationSettings updates an existing ExaminationSettings entity.
func UpdateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	return ExaminationSettingsRepo.Update(&entity, editorID)
}

// SaveExaminationSettings saves a ExaminationSettings entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateExaminationSettings(entity, editorID)
	} else {
		return UpdateExaminationSettings(entity, editorID)
	}
}

// DeleteExaminationSettings deletes a ExaminationSettings entity by its parents' path and examinationSettingsID.
func DeleteExaminationSettings(examinationSettingsID string) error {
	return ExaminationSettingsRepo.Delete(examination_settings.GetDocumentPath(examinationSettingsID))
}
