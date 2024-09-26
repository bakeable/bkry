package repo

import (
	"github.com/bakeable/bkry/data/entities/airline_pricing"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AirlinePricingRepo = repository.NewRepository[*airline_pricing.AirlinePricing]()

// GetAirlinePricing retrieves a single AirlinePricing entity by its ID and airlinePricingID.
func GetAirlinePricing(airlinePricingID string) (airline_pricing.AirlinePricing, error) {
	entityMap, err := AirlinePricingRepo.Read(airline_pricing.GetDocumentPath( airlinePricingID))
	return airline_pricing.Decode(entityMap), err
}

// GetAirlinePricingOrNew retrieves a single AirlinePricing entity by its ID and airlinePricingID.
func GetAirlinePricingOrNew(airlinePricingID string) (airline_pricing.AirlinePricing, error) {
	entityMap, err := AirlinePricingRepo.Read(airline_pricing.GetDocumentPath( airlinePricingID))
	if err != nil || entityMap == nil {
		return airline_pricing.AirlinePricing{}, err
	}
	return airline_pricing.Decode(entityMap), err
}

// GetAirlinePricing retrieves a single AirlinePricing entity by its document path.
func GetAirlinePricingByPath(path string) (airline_pricing.AirlinePricing, error) {
	entityMap, err := AirlinePricingRepo.Read(path)
	return airline_pricing.Decode(entityMap), err
}

// FindAirlinePricing retrieves a AirlinePricing entity according to given queries.
func FindAirlinePricing(queries []datastore.Query) (airline_pricing.AirlinePricing, error) {
	entityMap, err := AirlinePricingRepo.Find(airline_pricing.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return airline_pricing.AirlinePricing{}, err
	}
	return airline_pricing.Decode(entityMap), err
}

// GetAllAirlinePricings retrieves all AirlinePricing entities.
func GetAllAirlinePricings() ([]airline_pricing.AirlinePricing, error) {
	entityMaps, err := AirlinePricingRepo.ReadAll(airline_pricing.GetCollectionPath())
	if err != nil {
		return []airline_pricing.AirlinePricing{}, err
	}
	return airline_pricing.DecodeAll(entityMaps), nil
}


// GetAllAirlinePricingsPaginated retrieves all AirlinePricing entities in a paginated manner.
func GetAllAirlinePricingsPaginated(pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, datastore.Pagination, error) {
	entityMaps, pagination, err := AirlinePricingRepo.ReadPaginated(airline_pricing.GetCollectionPath(), pagination)
	if err != nil {
		return []airline_pricing.AirlinePricing{}, pagination, err
	}
	return airline_pricing.DecodeAll(entityMaps), pagination, nil
}

// QueryAirlinePricings retrieves all AirlinePricing entities according to given queries.
func QueryAirlinePricings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, error) {
	entityMaps, err := AirlinePricingRepo.Query(airline_pricing.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []airline_pricing.AirlinePricing{}, err
	}
	return airline_pricing.DecodeAll(entityMaps), nil
}

// QueryAirlinePricingsGroup retrieves all AirlinePricing entities according to given queries.
func QueryAirlinePricingsGroup(queries []datastore.Query) ([]airline_pricing.AirlinePricing, error) {
	entityMaps, err := AirlinePricingRepo.QueryGroup("airline_pricings", queries)
	if err != nil {
		return []airline_pricing.AirlinePricing{}, err
	}
	return airline_pricing.DecodeAll(entityMaps), nil
}

// CreateAirlinePricing creates a new AirlinePricing entity.
func CreateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	return AirlinePricingRepo.Create(&entity, editorID)
}

// UpdateAirlinePricing updates an existing AirlinePricing entity.
func UpdateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	return AirlinePricingRepo.Update(&entity, editorID)
}

// SaveAirlinePricing saves a AirlinePricing entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAirlinePricing(entity, editorID)
	} else {
		return UpdateAirlinePricing(entity, editorID)
	}
}

// DeleteAirlinePricing deletes a AirlinePricing entity by its parents' path and airlinePricingID.
func DeleteAirlinePricing(airlinePricingID string) error {
	return AirlinePricingRepo.Delete(airline_pricing.GetDocumentPath(airlinePricingID))
}
