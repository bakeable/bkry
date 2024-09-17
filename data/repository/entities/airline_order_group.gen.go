package repo

import (
	"github.com/bakeable/bkry/data/entities/airline_order_group"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AirlineOrderGroupRepo = repository.NewRepository[*airline_order_group.AirlineOrderGroup]()

// GetAirlineOrderGroup retrieves a single AirlineOrderGroup entity by its ID and airlineOrderGroupID.
func GetAirlineOrderGroup(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error) {
	entityMap, err := AirlineOrderGroupRepo.Read(airline_order_group.GetDocumentPath( airlineOrderGroupID))
	return airline_order_group.Decode(entityMap), err
}

// GetAirlineOrderGroupOrNew retrieves a single AirlineOrderGroup entity by its ID and airlineOrderGroupID.
func GetAirlineOrderGroupOrNew(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error) {
	entityMap, err := AirlineOrderGroupRepo.Read(airline_order_group.GetDocumentPath( airlineOrderGroupID))
	if err != nil || entityMap == nil {
		return airline_order_group.AirlineOrderGroup{}, err
	}
	return airline_order_group.Decode(entityMap), err
}

// GetAirlineOrderGroup retrieves a single AirlineOrderGroup entity by its document path.
func GetAirlineOrderGroupByPath(path string) (airline_order_group.AirlineOrderGroup, error) {
	entityMap, err := AirlineOrderGroupRepo.Read(path)
	return airline_order_group.Decode(entityMap), err
}

// FindAirlineOrderGroup retrieves a AirlineOrderGroup entity according to given queries.
func FindAirlineOrderGroup(queries []datastore.Query) (airline_order_group.AirlineOrderGroup, error) {
	entityMap, err := AirlineOrderGroupRepo.Find(airline_order_group.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return airline_order_group.AirlineOrderGroup{}, err
	}
	return airline_order_group.Decode(entityMap), err
}

// GetAllAirlineOrderGroups retrieves all AirlineOrderGroup entities.
func GetAllAirlineOrderGroups() ([]airline_order_group.AirlineOrderGroup, error) {
	entityMaps, err := AirlineOrderGroupRepo.ReadAll(airline_order_group.GetCollectionPath())
	if err != nil {
		return []airline_order_group.AirlineOrderGroup{}, err
	}
	return airline_order_group.DecodeAll(entityMaps), nil
}


// GetAllAirlineOrderGroupsPaginated retrieves all AirlineOrderGroup entities in a paginated manner.
func GetAllAirlineOrderGroupsPaginated(pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, datastore.Pagination, error) {
	entityMaps, pagination, err := AirlineOrderGroupRepo.ReadPaginated(airline_order_group.GetCollectionPath(), pagination)
	if err != nil {
		return []airline_order_group.AirlineOrderGroup{}, pagination, err
	}
	return airline_order_group.DecodeAll(entityMaps), pagination, nil
}

// QueryAirlineOrderGroups retrieves all AirlineOrderGroup entities according to given queries.
func QueryAirlineOrderGroups(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, error) {
	entityMaps, err := AirlineOrderGroupRepo.Query(airline_order_group.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []airline_order_group.AirlineOrderGroup{}, err
	}
	return airline_order_group.DecodeAll(entityMaps), nil
}

// QueryAirlineOrderGroupsGroup retrieves all AirlineOrderGroup entities according to given queries.
func QueryAirlineOrderGroupsGroup(queries []datastore.Query) ([]airline_order_group.AirlineOrderGroup, error) {
	entityMaps, err := AirlineOrderGroupRepo.QueryGroup("airline_order_groups", queries)
	if err != nil {
		return []airline_order_group.AirlineOrderGroup{}, err
	}
	return airline_order_group.DecodeAll(entityMaps), nil
}

// CreateAirlineOrderGroup creates a new AirlineOrderGroup entity.
func CreateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	return AirlineOrderGroupRepo.Create(&entity, editorID)
}

// UpdateAirlineOrderGroup updates an existing AirlineOrderGroup entity.
func UpdateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	return AirlineOrderGroupRepo.Update(&entity, editorID)
}

// SaveAirlineOrderGroup saves a AirlineOrderGroup entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAirlineOrderGroup(entity, editorID)
	} else {
		return UpdateAirlineOrderGroup(entity, editorID)
	}
}

// DeleteAirlineOrderGroup deletes a AirlineOrderGroup entity by its parents' path and airlineOrderGroupID.
func DeleteAirlineOrderGroup(airlineOrderGroupID string) error {
	return AirlineOrderGroupRepo.Delete(airline_order_group.GetDocumentPath(airlineOrderGroupID))
}
