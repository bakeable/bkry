package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// PutEntityById updates part of a single entity in Datastore by kind and id.
func PutEntityById(kind string, id interface{}, data map[string]interface{}, ancestor *datastore.Key) (*datastore.Key, error) {
	// Get context
	ctx := context.Background()

	// Start a new transaction
	tx, err := DatastoreClient.NewTransaction(datastoreClient, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	// Create key
	var key *datastore.Key
	if ancestor != nil {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), ancestor)
	} else {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), nil)
	}

	// Get current entity
	var currentEntity datastore.PropertyList
	if err := tx.Get(key, &currentEntity); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get current entity: %v", err)
	}

	// Convert current entity back to map[string]interface{}
	currentData := ConvertPropertiesToMap(currentEntity)

	// Update entity with new data
	for k, v := range data {
		currentData[k] = v
	}

	// Save the updated entity
	if _, err := tx.Put(key, ConvertMapToProperties(currentData)); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to save updated entity: %v", err)
	}

	// Commit transaction
	if _, err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Return the key
	return key, nil
}

// PutEntitiesByIds partially updates multiple entities in Datastore by kind and ids.
func PutEntitiesByIds(kind string, ids []interface{}, data []map[string]interface{}, ancestor *datastore.Key) ([]*datastore.Key, error) {
	// Get context
	ctx := context.Background()

	// Start a new transaction
	tx, err := DatastoreClient.NewTransaction(datastoreClient, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	// Create keys
	keys := make([]*datastore.Key, len(ids))
	for i := range ids {
		if ancestor != nil {
			keys[i] = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(ids[i])), ancestor)
		} else {
			keys[i] = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(ids[i])), nil)
		}
	}

	// Get current entities
	currentEntities := make([]datastore.PropertyList, len(ids))
	for i, key := range keys {
		if err := tx.Get(key, &currentEntities[i]); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to get current entity: %v", err)
		}
	}

	// Update entities with new data
	updatedEntities := make([]*datastore.PropertyList, len(data))
	for i := range data {
		currentData := ConvertPropertiesToMap(currentEntities[i])

		for k, v := range data[i] {
			currentData[k] = v
		}

		updatedEntity := ConvertMapToPropertyList(currentData)
		updatedEntities[i] = &updatedEntity
	}

	// Save the updated entities
	for i, key := range keys {
		if _, err := tx.Put(key, updatedEntities[i]); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to save updated entity: %v", err)
		}
	}

	// Commit transaction
	if _, err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Return the keys
	return keys, nil
}
