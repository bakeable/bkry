package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/airline_order"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AirlineOrderRepo = repository.NewRepository[*airline_order.AirlineOrder]()

// GetAirlineOrder retrieves a single AirlineOrder entity by its ID and airlineOrderID.
func GetAirlineOrder(airlineOrderID string) (airline_order.AirlineOrder, error) {
	entityMap, err := AirlineOrderRepo.Read(airline_order.GetDocumentPath( airlineOrderID))
	return airline_order.Decode(entityMap), err
}

// GetAirlineOrderOrNew retrieves a single AirlineOrder entity by its ID and airlineOrderID.
func GetAirlineOrderOrNew(airlineOrderID string) (airline_order.AirlineOrder, error) {
	entityMap, err := AirlineOrderRepo.Read(airline_order.GetDocumentPath( airlineOrderID))
	if err != nil || entityMap == nil {
		return airline_order.AirlineOrder{}, err
	}
	return airline_order.Decode(entityMap), err
}

// GetAirlineOrder retrieves a single AirlineOrder entity by its document path.
func GetAirlineOrderByPath(path string) (airline_order.AirlineOrder, error) {
	entityMap, err := AirlineOrderRepo.Read(path)
	return airline_order.Decode(entityMap), err
}

// FindAirlineOrder retrieves a AirlineOrder entity according to given queries.
func FindAirlineOrder(queries []datastore.Query) (airline_order.AirlineOrder, error) {
	entityMap, err := AirlineOrderRepo.Find(airline_order.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return airline_order.AirlineOrder{}, err
	}
	return airline_order.Decode(entityMap), err
}

// GetAllAirlineOrders retrieves all AirlineOrder entities.
func GetAllAirlineOrders() ([]airline_order.AirlineOrder, error) {
	entityMaps, err := AirlineOrderRepo.ReadAll(airline_order.GetCollectionPath())
	if err != nil {
		return []airline_order.AirlineOrder{}, err
	}
	return airline_order.DecodeAll(entityMaps), nil
}


// GetAllAirlineOrdersPaginated retrieves all AirlineOrder entities in a paginated manner.
func GetAllAirlineOrdersPaginated(pagination datastore.Pagination) ([]airline_order.AirlineOrder, datastore.Pagination, error) {
	entityMaps, pagination, err := AirlineOrderRepo.ReadPaginated(airline_order.GetCollectionPath(), pagination)
	if err != nil {
		return []airline_order.AirlineOrder{}, pagination, err
	}
	return airline_order.DecodeAll(entityMaps), pagination, nil
}

// QueryAirlineOrders retrieves all AirlineOrder entities according to given queries.
func QueryAirlineOrders(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order.AirlineOrder, error) {
	entityMaps, err := AirlineOrderRepo.Query(airline_order.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []airline_order.AirlineOrder{}, err
	}
	return airline_order.DecodeAll(entityMaps), nil
}

// QueryAirlineOrdersGroup retrieves all AirlineOrder entities according to given queries.
func QueryAirlineOrdersGroup(queries []datastore.Query) ([]airline_order.AirlineOrder, error) {
	entityMaps, err := AirlineOrderRepo.QueryGroup("airline_orders", queries)
	if err != nil {
		return []airline_order.AirlineOrder{}, err
	}
	return airline_order.DecodeAll(entityMaps), nil
}

// CreateAirlineOrder creates a new AirlineOrder entity.
func CreateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	return AirlineOrderRepo.Create(&entity, editorID)
}

// UpdateAirlineOrder updates an existing AirlineOrder entity.
func UpdateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	return AirlineOrderRepo.Update(&entity, editorID)
}

// SaveAirlineOrder saves a AirlineOrder entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAirlineOrder(entity, editorID)
	} else {
		return UpdateAirlineOrder(entity, editorID)
	}
}

// DeleteAirlineOrder deletes a AirlineOrder entity by its parents' path and airlineOrderID.
func DeleteAirlineOrder(airlineOrderID string) error {
	return AirlineOrderRepo.Delete(airline_order.GetDocumentPath(airlineOrderID))
}
