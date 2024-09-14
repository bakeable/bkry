package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// DeleteEntityById deletes a single entity from Datastore by kind and id.
func DeleteEntityById(kind string, id interface{}, ancestor *datastore.Key) error {
	// Get context
	ctx := context.Background()

	// Create key
	var key *datastore.Key
	if ancestor != nil {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), ancestor)
	} else {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), nil)
	}

	// Delete the entity
	if err := DatastoreClient.Delete(datastoreClient, ctx, key); err != nil {
		return fmt.Errorf("failed to delete entity: %v", err)
	}

	return nil
}

// DeleteEntitiesByIds deletes multiple entities from Datastore by kind and ids.
func DeleteEntitiesByIds(kind string, ids []interface{}, ancestor *datastore.Key) error {
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

	// Delete the entities
	if err := DatastoreClient.DeleteMulti(datastoreClient, ctx, keys); err != nil {
		return fmt.Errorf("failed to delete entities: %v", err)
	}

	return nil
}

// DeleteEntitiesByQueries deletes multiple entities from Datastore by kind and queries.
func DeleteEntitiesByQueries(kind string, queries []Query, ancestor *datastore.Key) error {
	// Get context
	ctx := context.Background()

	// Create a query
	q := datastore.NewQuery(kind)

	// Apply ancestor filter if present
	if ancestor != nil {
		q = q.Ancestor(ancestor)
	}

	// Apply filters using FilterField
	for _, query := range queries {
		q = q.FilterField(query.Field, string(query.Operator), query.Value)
	}

	// Run the query to get entities
	var entities []datastore.PropertyList
	keys, err := DatastoreClient.GetAll(datastoreClient, ctx, q, &entities)
	if err != nil {
		return fmt.Errorf("failed to run query: %v", err)
	}

	// Delete the entities
	if err := DatastoreClient.DeleteMulti(datastoreClient, ctx, keys); err != nil {
		return fmt.Errorf("failed to delete entities: %v", err)
	}

	return nil
}