package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/airline_settings"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AirlineSettingsRepo = repository.NewRepository[*airline_settings.AirlineSettings]()

// GetAirlineSettings retrieves a single AirlineSettings entity by its ID and airlineSettingsID.
func GetAirlineSettings(airlineSettingsID string) (airline_settings.AirlineSettings, error) {
	entityMap, err := AirlineSettingsRepo.Read(airline_settings.GetDocumentPath( airlineSettingsID))
	return airline_settings.Decode(entityMap), err
}

// GetAirlineSettingsOrNew retrieves a single AirlineSettings entity by its ID and airlineSettingsID.
func GetAirlineSettingsOrNew(airlineSettingsID string) (airline_settings.AirlineSettings, error) {
	entityMap, err := AirlineSettingsRepo.Read(airline_settings.GetDocumentPath( airlineSettingsID))
	if err != nil || entityMap == nil {
		return airline_settings.AirlineSettings{}, err
	}
	return airline_settings.Decode(entityMap), err
}

// GetAirlineSettings retrieves a single AirlineSettings entity by its document path.
func GetAirlineSettingsByPath(path string) (airline_settings.AirlineSettings, error) {
	entityMap, err := AirlineSettingsRepo.Read(path)
	return airline_settings.Decode(entityMap), err
}

// FindAirlineSettings retrieves a AirlineSettings entity according to given queries.
func FindAirlineSettings(queries []datastore.Query) (airline_settings.AirlineSettings, error) {
	entityMap, err := AirlineSettingsRepo.Find(airline_settings.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return airline_settings.AirlineSettings{}, err
	}
	return airline_settings.Decode(entityMap), err
}

// GetAllAirlineSettings retrieves all AirlineSettings entities.
func GetAllAirlineSettings() ([]airline_settings.AirlineSettings, error) {
	entityMaps, err := AirlineSettingsRepo.ReadAll(airline_settings.GetCollectionPath())
	if err != nil {
		return []airline_settings.AirlineSettings{}, err
	}
	return airline_settings.DecodeAll(entityMaps), nil
}


// GetAllAirlineSettingsPaginated retrieves all AirlineSettings entities in a paginated manner.
func GetAllAirlineSettingsPaginated(pagination datastore.Pagination) ([]airline_settings.AirlineSettings, datastore.Pagination, error) {
	entityMaps, pagination, err := AirlineSettingsRepo.ReadPaginated(airline_settings.GetCollectionPath(), pagination)
	if err != nil {
		return []airline_settings.AirlineSettings{}, pagination, err
	}
	return airline_settings.DecodeAll(entityMaps), pagination, nil
}

// QueryAirlineSettings retrieves all AirlineSettings entities according to given queries.
func QueryAirlineSettings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_settings.AirlineSettings, error) {
	entityMaps, err := AirlineSettingsRepo.Query(airline_settings.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []airline_settings.AirlineSettings{}, err
	}
	return airline_settings.DecodeAll(entityMaps), nil
}

// QueryAirlineSettingsGroup retrieves all AirlineSettings entities according to given queries.
func QueryAirlineSettingsGroup(queries []datastore.Query) ([]airline_settings.AirlineSettings, error) {
	entityMaps, err := AirlineSettingsRepo.QueryGroup("airline_settings", queries)
	if err != nil {
		return []airline_settings.AirlineSettings{}, err
	}
	return airline_settings.DecodeAll(entityMaps), nil
}

// CreateAirlineSettings creates a new AirlineSettings entity.
func CreateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	return AirlineSettingsRepo.Create(&entity, editorID)
}

// UpdateAirlineSettings updates an existing AirlineSettings entity.
func UpdateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	return AirlineSettingsRepo.Update(&entity, editorID)
}

// SaveAirlineSettings saves a AirlineSettings entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAirlineSettings(entity, editorID)
	} else {
		return UpdateAirlineSettings(entity, editorID)
	}
}

// DeleteAirlineSettings deletes a AirlineSettings entity by its parents' path and airlineSettingsID.
func DeleteAirlineSettings(airlineSettingsID string) error {
	return AirlineSettingsRepo.Delete(airline_settings.GetDocumentPath(airlineSettingsID))
}
