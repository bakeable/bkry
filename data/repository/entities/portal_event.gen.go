package repo

import (
	"github.com/bakeable/bkry/data/entities/portal_event"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PortalEventRepo = repository.NewRepository[*portal_event.PortalEvent]()

// GetPortalEvent retrieves a single PortalEvent entity by its ID and portalEventID.
func GetPortalEvent(portalEventID string) (portal_event.PortalEvent, error) {
	entityMap, err := PortalEventRepo.Read(portal_event.GetDocumentPath( portalEventID))
	return portal_event.Decode(entityMap), err
}

// GetPortalEventOrNew retrieves a single PortalEvent entity by its ID and portalEventID.
func GetPortalEventOrNew(portalEventID string) (portal_event.PortalEvent, error) {
	entityMap, err := PortalEventRepo.Read(portal_event.GetDocumentPath( portalEventID))
	if err != nil || entityMap == nil {
		return portal_event.PortalEvent{}, err
	}
	return portal_event.Decode(entityMap), err
}

// GetPortalEvent retrieves a single PortalEvent entity by its document path.
func GetPortalEventByPath(path string) (portal_event.PortalEvent, error) {
	entityMap, err := PortalEventRepo.Read(path)
	return portal_event.Decode(entityMap), err
}

// FindPortalEvent retrieves a PortalEvent entity according to given queries.
func FindPortalEvent(queries []datastore.Query) (portal_event.PortalEvent, error) {
	entityMap, err := PortalEventRepo.Find(portal_event.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return portal_event.PortalEvent{}, err
	}
	return portal_event.Decode(entityMap), err
}

// GetAllPortalEvents retrieves all PortalEvent entities.
func GetAllPortalEvents() ([]portal_event.PortalEvent, error) {
	entityMaps, err := PortalEventRepo.ReadAll(portal_event.GetCollectionPath())
	if err != nil {
		return []portal_event.PortalEvent{}, err
	}
	return portal_event.DecodeAll(entityMaps), nil
}


// GetAllPortalEventsPaginated retrieves all PortalEvent entities in a paginated manner.
func GetAllPortalEventsPaginated(pagination datastore.Pagination) ([]portal_event.PortalEvent, datastore.Pagination, error) {
	entityMaps, pagination, err := PortalEventRepo.ReadPaginated(portal_event.GetCollectionPath(), pagination)
	if err != nil {
		return []portal_event.PortalEvent{}, pagination, err
	}
	return portal_event.DecodeAll(entityMaps), pagination, nil
}

// QueryPortalEvents retrieves all PortalEvent entities according to given queries.
func QueryPortalEvents(queries []datastore.Query, pagination datastore.Pagination) ([]portal_event.PortalEvent, error) {
	entityMaps, err := PortalEventRepo.Query(portal_event.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []portal_event.PortalEvent{}, err
	}
	return portal_event.DecodeAll(entityMaps), nil
}

// QueryPortalEventsGroup retrieves all PortalEvent entities according to given queries.
func QueryPortalEventsGroup(queries []datastore.Query) ([]portal_event.PortalEvent, error) {
	entityMaps, err := PortalEventRepo.QueryGroup("portal_events", queries)
	if err != nil {
		return []portal_event.PortalEvent{}, err
	}
	return portal_event.DecodeAll(entityMaps), nil
}

// CreatePortalEvent creates a new PortalEvent entity.
func CreatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	return PortalEventRepo.Create(&entity, editorID)
}

// UpdatePortalEvent updates an existing PortalEvent entity.
func UpdatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	return PortalEventRepo.Update(&entity, editorID)
}

// SavePortalEvent saves a PortalEvent entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePortalEvent(entity, editorID)
	} else {
		return UpdatePortalEvent(entity, editorID)
	}
}

// DeletePortalEvent deletes a PortalEvent entity by its parents' path and portalEventID.
func DeletePortalEvent(portalEventID string) error {
	return PortalEventRepo.Delete(portal_event.GetDocumentPath(portalEventID))
}
