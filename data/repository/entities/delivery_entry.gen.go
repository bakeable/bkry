package repo

import (
	"github.com/bakeable/bkry/data/entities/delivery_entry"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var DeliveryEntryRepo = repository.NewRepository[*delivery_entry.DeliveryEntry]()

// GetDeliveryEntry retrieves a single DeliveryEntry entity by its ID and deliveryEntryID.
func GetDeliveryEntry(deliveryEntryID string) (delivery_entry.DeliveryEntry, error) {
	entityMap, err := DeliveryEntryRepo.Read(delivery_entry.GetDocumentPath( deliveryEntryID))
	return delivery_entry.Decode(entityMap), err
}

// GetDeliveryEntryOrNew retrieves a single DeliveryEntry entity by its ID and deliveryEntryID.
func GetDeliveryEntryOrNew(deliveryEntryID string) (delivery_entry.DeliveryEntry, error) {
	entityMap, err := DeliveryEntryRepo.Read(delivery_entry.GetDocumentPath( deliveryEntryID))
	if err != nil || entityMap == nil {
		return delivery_entry.DeliveryEntry{}, err
	}
	return delivery_entry.Decode(entityMap), err
}

// GetDeliveryEntry retrieves a single DeliveryEntry entity by its document path.
func GetDeliveryEntryByPath(path string) (delivery_entry.DeliveryEntry, error) {
	entityMap, err := DeliveryEntryRepo.Read(path)
	return delivery_entry.Decode(entityMap), err
}

// FindDeliveryEntry retrieves a DeliveryEntry entity according to given queries.
func FindDeliveryEntry(queries []datastore.Query) (delivery_entry.DeliveryEntry, error) {
	entityMap, err := DeliveryEntryRepo.Find(delivery_entry.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return delivery_entry.DeliveryEntry{}, err
	}
	return delivery_entry.Decode(entityMap), err
}

// GetAllDeliveryEntries retrieves all DeliveryEntry entities.
func GetAllDeliveryEntries() ([]delivery_entry.DeliveryEntry, error) {
	entityMaps, err := DeliveryEntryRepo.ReadAll(delivery_entry.GetCollectionPath())
	if err != nil {
		return []delivery_entry.DeliveryEntry{}, err
	}
	return delivery_entry.DecodeAll(entityMaps), nil
}


// GetAllDeliveryEntriesPaginated retrieves all DeliveryEntry entities in a paginated manner.
func GetAllDeliveryEntriesPaginated(pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, datastore.Pagination, error) {
	entityMaps, pagination, err := DeliveryEntryRepo.ReadPaginated(delivery_entry.GetCollectionPath(), pagination)
	if err != nil {
		return []delivery_entry.DeliveryEntry{}, pagination, err
	}
	return delivery_entry.DecodeAll(entityMaps), pagination, nil
}

// QueryDeliveryEntries retrieves all DeliveryEntry entities according to given queries.
func QueryDeliveryEntries(queries []datastore.Query, pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, error) {
	entityMaps, err := DeliveryEntryRepo.Query(delivery_entry.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []delivery_entry.DeliveryEntry{}, err
	}
	return delivery_entry.DecodeAll(entityMaps), nil
}

// QueryDeliveryEntriesGroup retrieves all DeliveryEntry entities according to given queries.
func QueryDeliveryEntriesGroup(queries []datastore.Query) ([]delivery_entry.DeliveryEntry, error) {
	entityMaps, err := DeliveryEntryRepo.QueryGroup("delivery_entries", queries)
	if err != nil {
		return []delivery_entry.DeliveryEntry{}, err
	}
	return delivery_entry.DecodeAll(entityMaps), nil
}

// CreateDeliveryEntry creates a new DeliveryEntry entity.
func CreateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	return DeliveryEntryRepo.Create(&entity, editorID)
}

// UpdateDeliveryEntry updates an existing DeliveryEntry entity.
func UpdateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	return DeliveryEntryRepo.Update(&entity, editorID)
}

// SaveDeliveryEntry saves a DeliveryEntry entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateDeliveryEntry(entity, editorID)
	} else {
		return UpdateDeliveryEntry(entity, editorID)
	}
}

// DeleteDeliveryEntry deletes a DeliveryEntry entity by its parents' path and deliveryEntryID.
func DeleteDeliveryEntry(deliveryEntryID string) error {
	return DeliveryEntryRepo.Delete(delivery_entry.GetDocumentPath(deliveryEntryID))
}
