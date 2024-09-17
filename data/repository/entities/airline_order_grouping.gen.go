package repo

import (
	"github.com/bakeable/bkry/data/entities/airline_order_grouping"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AirlineOrderGroupingRepo = repository.NewRepository[*airline_order_grouping.AirlineOrderGrouping]()

// GetAirlineOrderGrouping retrieves a single AirlineOrderGrouping entity by its ID and airlineOrderGroupingID.
func GetAirlineOrderGrouping(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error) {
	entityMap, err := AirlineOrderGroupingRepo.Read(airline_order_grouping.GetDocumentPath( airlineOrderGroupingID))
	return airline_order_grouping.Decode(entityMap), err
}

// GetAirlineOrderGroupingOrNew retrieves a single AirlineOrderGrouping entity by its ID and airlineOrderGroupingID.
func GetAirlineOrderGroupingOrNew(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error) {
	entityMap, err := AirlineOrderGroupingRepo.Read(airline_order_grouping.GetDocumentPath( airlineOrderGroupingID))
	if err != nil || entityMap == nil {
		return airline_order_grouping.AirlineOrderGrouping{}, err
	}
	return airline_order_grouping.Decode(entityMap), err
}

// GetAirlineOrderGrouping retrieves a single AirlineOrderGrouping entity by its document path.
func GetAirlineOrderGroupingByPath(path string) (airline_order_grouping.AirlineOrderGrouping, error) {
	entityMap, err := AirlineOrderGroupingRepo.Read(path)
	return airline_order_grouping.Decode(entityMap), err
}

// FindAirlineOrderGrouping retrieves a AirlineOrderGrouping entity according to given queries.
func FindAirlineOrderGrouping(queries []datastore.Query) (airline_order_grouping.AirlineOrderGrouping, error) {
	entityMap, err := AirlineOrderGroupingRepo.Find(airline_order_grouping.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return airline_order_grouping.AirlineOrderGrouping{}, err
	}
	return airline_order_grouping.Decode(entityMap), err
}

// GetAllAirlineOrderGroupings retrieves all AirlineOrderGrouping entities.
func GetAllAirlineOrderGroupings() ([]airline_order_grouping.AirlineOrderGrouping, error) {
	entityMaps, err := AirlineOrderGroupingRepo.ReadAll(airline_order_grouping.GetCollectionPath())
	if err != nil {
		return []airline_order_grouping.AirlineOrderGrouping{}, err
	}
	return airline_order_grouping.DecodeAll(entityMaps), nil
}


// GetAllAirlineOrderGroupingsPaginated retrieves all AirlineOrderGrouping entities in a paginated manner.
func GetAllAirlineOrderGroupingsPaginated(pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, datastore.Pagination, error) {
	entityMaps, pagination, err := AirlineOrderGroupingRepo.ReadPaginated(airline_order_grouping.GetCollectionPath(), pagination)
	if err != nil {
		return []airline_order_grouping.AirlineOrderGrouping{}, pagination, err
	}
	return airline_order_grouping.DecodeAll(entityMaps), pagination, nil
}

// QueryAirlineOrderGroupings retrieves all AirlineOrderGrouping entities according to given queries.
func QueryAirlineOrderGroupings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, error) {
	entityMaps, err := AirlineOrderGroupingRepo.Query(airline_order_grouping.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []airline_order_grouping.AirlineOrderGrouping{}, err
	}
	return airline_order_grouping.DecodeAll(entityMaps), nil
}

// QueryAirlineOrderGroupingsGroup retrieves all AirlineOrderGrouping entities according to given queries.
func QueryAirlineOrderGroupingsGroup(queries []datastore.Query) ([]airline_order_grouping.AirlineOrderGrouping, error) {
	entityMaps, err := AirlineOrderGroupingRepo.QueryGroup("airline_order_groupings", queries)
	if err != nil {
		return []airline_order_grouping.AirlineOrderGrouping{}, err
	}
	return airline_order_grouping.DecodeAll(entityMaps), nil
}

// CreateAirlineOrderGrouping creates a new AirlineOrderGrouping entity.
func CreateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	return AirlineOrderGroupingRepo.Create(&entity, editorID)
}

// UpdateAirlineOrderGrouping updates an existing AirlineOrderGrouping entity.
func UpdateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	return AirlineOrderGroupingRepo.Update(&entity, editorID)
}

// SaveAirlineOrderGrouping saves a AirlineOrderGrouping entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAirlineOrderGrouping(entity, editorID)
	} else {
		return UpdateAirlineOrderGrouping(entity, editorID)
	}
}

// DeleteAirlineOrderGrouping deletes a AirlineOrderGrouping entity by its parents' path and airlineOrderGroupingID.
func DeleteAirlineOrderGrouping(airlineOrderGroupingID string) error {
	return AirlineOrderGroupingRepo.Delete(airline_order_grouping.GetDocumentPath(airlineOrderGroupingID))
}
