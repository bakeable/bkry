package repo

import (
	"github.com/bakeable/bkry/data/entities/price_configuration"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PriceConfigurationRepo = repository.NewRepository[*price_configuration.PriceConfiguration]()

// GetPriceConfiguration retrieves a single PriceConfiguration entity by its ID and priceConfigurationID.
func GetPriceConfiguration(priceConfigurationID string) (price_configuration.PriceConfiguration, error) {
	entityMap, err := PriceConfigurationRepo.Read(price_configuration.GetDocumentPath( priceConfigurationID))
	return price_configuration.Decode(entityMap), err
}

// GetPriceConfigurationOrNew retrieves a single PriceConfiguration entity by its ID and priceConfigurationID.
func GetPriceConfigurationOrNew(priceConfigurationID string) (price_configuration.PriceConfiguration, error) {
	entityMap, err := PriceConfigurationRepo.Read(price_configuration.GetDocumentPath( priceConfigurationID))
	if err != nil || entityMap == nil {
		return price_configuration.PriceConfiguration{}, err
	}
	return price_configuration.Decode(entityMap), err
}

// GetPriceConfiguration retrieves a single PriceConfiguration entity by its document path.
func GetPriceConfigurationByPath(path string) (price_configuration.PriceConfiguration, error) {
	entityMap, err := PriceConfigurationRepo.Read(path)
	return price_configuration.Decode(entityMap), err
}

// FindPriceConfiguration retrieves a PriceConfiguration entity according to given queries.
func FindPriceConfiguration(queries []datastore.Query) (price_configuration.PriceConfiguration, error) {
	entityMap, err := PriceConfigurationRepo.Find(price_configuration.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return price_configuration.PriceConfiguration{}, err
	}
	return price_configuration.Decode(entityMap), err
}

// GetAllPriceConfigurations retrieves all PriceConfiguration entities.
func GetAllPriceConfigurations() ([]price_configuration.PriceConfiguration, error) {
	entityMaps, err := PriceConfigurationRepo.ReadAll(price_configuration.GetCollectionPath())
	if err != nil {
		return []price_configuration.PriceConfiguration{}, err
	}
	return price_configuration.DecodeAll(entityMaps), nil
}


// GetAllPriceConfigurationsPaginated retrieves all PriceConfiguration entities in a paginated manner.
func GetAllPriceConfigurationsPaginated(pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, datastore.Pagination, error) {
	entityMaps, pagination, err := PriceConfigurationRepo.ReadPaginated(price_configuration.GetCollectionPath(), pagination)
	if err != nil {
		return []price_configuration.PriceConfiguration{}, pagination, err
	}
	return price_configuration.DecodeAll(entityMaps), pagination, nil
}

// QueryPriceConfigurations retrieves all PriceConfiguration entities according to given queries.
func QueryPriceConfigurations(queries []datastore.Query, pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, error) {
	entityMaps, err := PriceConfigurationRepo.Query(price_configuration.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []price_configuration.PriceConfiguration{}, err
	}
	return price_configuration.DecodeAll(entityMaps), nil
}

// QueryPriceConfigurationsGroup retrieves all PriceConfiguration entities according to given queries.
func QueryPriceConfigurationsGroup(queries []datastore.Query) ([]price_configuration.PriceConfiguration, error) {
	entityMaps, err := PriceConfigurationRepo.QueryGroup("price_configurations", queries)
	if err != nil {
		return []price_configuration.PriceConfiguration{}, err
	}
	return price_configuration.DecodeAll(entityMaps), nil
}

// CreatePriceConfiguration creates a new PriceConfiguration entity.
func CreatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	return PriceConfigurationRepo.Create(&entity, editorID)
}

// UpdatePriceConfiguration updates an existing PriceConfiguration entity.
func UpdatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	return PriceConfigurationRepo.Update(&entity, editorID)
}

// SavePriceConfiguration saves a PriceConfiguration entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePriceConfiguration(entity, editorID)
	} else {
		return UpdatePriceConfiguration(entity, editorID)
	}
}

// DeletePriceConfiguration deletes a PriceConfiguration entity by its parents' path and priceConfigurationID.
func DeletePriceConfiguration(priceConfigurationID string) error {
	return PriceConfigurationRepo.Delete(price_configuration.GetDocumentPath(priceConfigurationID))
}
