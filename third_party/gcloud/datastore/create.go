package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// ParseID is a placeholder for a function that parses an ID. You should implement this based on your logic.
func ParseID(id interface{}) interface{} {
	// Add your parsing logic here
	return id
}

// CreateEntity creates a single entity in Datastore by kind and id.
func CreateEntity(kind string, data map[string]interface{}, ancestor *datastore.Key) (*datastore.Key, error) {
	// Get context
	ctx := context.Background()
	
	// Create key
	var key *datastore.Key
	if ancestor != nil {
		key = datastore.NameKey(kind, "", ancestor)
	} else {
		key = datastore.NameKey(kind, "", nil)
	}

	// Add key to the data
	data["id"] = key.ID

	// Create entity
	entity := &datastore.Entity{
		Key:   key,
		Properties: ConvertMapToProperties(data),
	}

	// Upsert
	if _, err := DatastoreClient.Put(datastoreClient, ctx, key, entity); err != nil {
		return nil, fmt.Errorf("failed to save entity: %v", err)
	}

	// Return ID
	return key, nil
}

// CreateEntities creates multiple entities in Datastore by kind and ids.
func CreateEntities(kind string, data []map[string]interface{}, ancestor *datastore.Key) ([]*datastore.Key, error) {
	// Get context
	ctx := context.Background()

	// Create keys
	keys := make([]*datastore.Key, len(data))
	for i := range data {
		if ancestor != nil {
			keys[i] = datastore.NameKey(kind, "", ancestor)
		} else {
			keys[i] = datastore.NameKey(kind, "", nil)
		}
	}

	// Entity slice
	entities := make([]*datastore.Entity, len(data))
	for i, entityData := range data {
		// Add key to the data
		entityData["id"] = keys[i].ID

		entities[i] = &datastore.Entity{
			Key:   keys[i],
			Properties: ConvertMapToProperties(entityData),
		}
	}

	// Upsert
	if _, err := DatastoreClient.PutMulti(datastoreClient, ctx, keys, entities); err != nil {
		return nil, fmt.Errorf("failed to save entities: %v", err)
	}

	// Return IDs
	return keys, nil
}
