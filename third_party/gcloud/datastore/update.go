package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// UpdateEntityById updates a single entity in Datastore by kind and id.
func UpdateEntityById(kind string, id interface{}, data map[string]interface{}, ancestor *datastore.Key) (*datastore.Key, error) {
	// Get context
	ctx := context.Background()

	// Create key
	var key *datastore.Key
	if ancestor != nil {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), ancestor)
	} else {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), nil)
	}

	// Create entity
	entity := &datastore.Entity{
		Key:        key,
		Properties: ConvertMapToProperties(data),
	}

	// Upsert
	if _, err := DatastoreClient.Put(datastoreClient, ctx, key, entity); err != nil {
		return nil, fmt.Errorf("failed to update entity: %v", err)
	}

	// Return updated key
	return key, nil
}

// UpdateEntitiesByIds updates multiple entities in Datastore by kind and ids.
func UpdateEntitiesByIds(kind string, ids []interface{}, data []map[string]interface{}, ancestor *datastore.Key) ([]*datastore.Key, error) {
	if len(ids) != len(data) {
		return nil, fmt.Errorf("ids and data slices must have the same length")
	}

	// Get context
	ctx := context.Background()

	// Create keys
	keys := make([]*datastore.Key, len(ids))
	for i := range ids {
		if ancestor != nil {
			keys[i] = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(ids[i])), ancestor)
		} else {
			keys[i] = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(ids[i])), nil)
		}
	}

	// Entity slice
	entities := make([]*datastore.Entity, len(data))
	for i := range data {
		entities[i] = &datastore.Entity{
			Key:        keys[i],
			Properties: ConvertMapToProperties(data[i]),
		}
	}

	// Upsert
	if _, err := DatastoreClient.PutMulti(datastoreClient, ctx, keys, entities); err != nil {
		return nil, fmt.Errorf("failed to update entities: %v", err)
	}

	// Return updated keys
	return keys, nil
}
