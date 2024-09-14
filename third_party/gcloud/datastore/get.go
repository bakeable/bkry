package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// GetEntityById retrieves a single entity from Datastore by kind and id.
func GetEntityById(kind string, id interface{}, ancestor *datastore.Key) (map[string]interface{}, error) {
	ctx := context.Background()

	// Create key
	var key *datastore.Key
	if ancestor != nil {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), ancestor)
	} else {
		key = datastore.NameKey(kind, fmt.Sprintf("%v", ParseID(id)), nil)
	}

	// Get entity
	var entity datastore.PropertyList
	if err := DatastoreClient.Get(datastoreClient, ctx, key, &entity); err != nil {
		return nil, fmt.Errorf("failed to get entity: %v", err)
	}

	// Print the properties
	fmt.Println(entity)
	for _, prop := range entity {
		fmt.Println("Property", prop.Name, prop.Value)
	}

	// Convert properties to map and add ID
	data := ConvertPropertiesToMap(entity)
	data["id"] = ParseID(id)

	return data, nil
}

// GetEntitiesByIds retrieves multiple entities from Datastore by kind and ids.
func GetEntitiesByIds(kind string, ids []interface{}, ancestor *datastore.Key) ([]map[string]interface{}, error) {
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

	// Get entities
	entities := make([]datastore.PropertyList, len(keys))
	if err := DatastoreClient.GetMulti(datastoreClient, ctx, keys, entities); err != nil {
		return nil, fmt.Errorf("failed to get entities: %v", err)
	}

	// Convert properties to map and add IDs
	results := make([]map[string]interface{}, len(entities))
	for i, entity := range entities {
		data := ConvertPropertiesToMap(entity)
		data["id"] = ParseID(ids[i])
		results[i] = data
	}

	return results, nil
}

// GetAllEntitiesByKind retrieves all entities of a specific kind.
func GetAllEntitiesByKind(kind string, ancestor *datastore.Key) ([]map[string]interface{}, error) {
	ctx := context.Background()

	fmt.Println("kind", kind)

	// Create query
	q := datastore.NewQuery(kind)
	fmt.Println("ancestor", ancestor)
	if ancestor != nil {
		q = q.Ancestor(ancestor)
	}

	// Run query
	var entities []datastore.PropertyList
	fmt.Println("query", q, ctx, entities)
	keys, err := DatastoreClient.GetAll(datastoreClient, ctx, q, &entities)
	if err != nil {
		return nil, fmt.Errorf("failed to get all entities: %v", err)
	}
	fmt.Println("Found", len(entities), "entities")

	// Convert properties to map and add IDs
	results := make([]map[string]interface{}, len(entities))
	for i, entity := range entities {
		data := ConvertPropertiesToMap(entity)
		data["id"] = ParseID(keys[i].Name)
		fmt.Println(keys[i].Name)
		results[i] = data
	}

	return results, nil
}

// QueryEntity retrieves the first entity returned from Datastore by kind and queries.
func QueryEntity(kind string, queries []Query, pagination *Pagination, ancestor *datastore.Key) (map[string]interface{}, error) {
	ctx := context.Background()

	// Create query
	q := datastore.NewQuery(kind)
	if ancestor != nil {
		q = q.Ancestor(ancestor)
	}

	// Apply filters
	for _, query := range queries {
		q = q.FilterField(query.Field, string(query.Operator), query.Value)
	}

	// Apply ordering
	if pagination != nil {
		orderBy := pagination.OrderBy
		if !pagination.Ascending {
			orderBy = "-" + orderBy
		}
		q = q.Order(orderBy)
	}

	// Run query
	var entities []datastore.PropertyList
	keys, err := DatastoreClient.GetAll(datastoreClient, ctx, q.Limit(1), &entities)
	if err != nil || len(keys) == 0 {
		return nil, fmt.Errorf("failed to query entity: %v", err)
	}

	// Convert properties to map and add ID
	data := ConvertPropertiesToMap(entities[0])
	data["id"] = ParseID(keys[0].ID)

	return data, nil
}

// QueryEntities retrieves multiple entities from Datastore by kind and queries.
func QueryEntities(kind string, queries []Query, pagination *Pagination, ancestor *datastore.Key) ([]map[string]interface{}, error) {
	ctx := context.Background()

	// Create query
	q := datastore.NewQuery(kind)
	if ancestor != nil {
		q = q.Ancestor(ancestor)
	}

	// Apply filters
	for _, query := range queries {
		q = q.FilterField(query.Field, string(query.Operator), query.Value)
	}


	// Apply ordering
	if pagination != nil {
		orderBy := pagination.OrderBy
		if !pagination.Ascending {
			orderBy = "-" + orderBy
		}
		q = q.Order(orderBy)
	}

	// Run query
	var entities []datastore.PropertyList
	keys, err := DatastoreClient.GetAll(datastoreClient, ctx, q, &entities)
	if err != nil {
		return nil, fmt.Errorf("failed to query entities: %v", err)
	}

	// Convert properties to map and add IDs
	results := make([]map[string]interface{}, len(entities))
	for i, entity := range entities {
		data := ConvertPropertiesToMap(entity)
		data["id"] = ParseID(keys[i].ID)
		results[i] = data
	}

	return results, nil
}

// GetPaginatedEntitiesByKind retrieves paginated entities from Datastore by kind.
func GetPaginatedEntitiesByKind(kind string, pagination Pagination, ancestor *datastore.Key) ([]map[string]interface{}, *datastore.Cursor, error) {
	ctx := context.Background()

	// Create query
	q := datastore.NewQuery(kind).Order(pagination.OrderBy)
	if ancestor != nil {
		q = q.Ancestor(ancestor)
	}


	// Apply ordering
	orderBy := pagination.OrderBy
	if !pagination.Ascending {
		orderBy = "-" + orderBy
	}
	q = q.Order(orderBy)

	// Set pagination
	q = q.Limit(pagination.PageSize).Offset((pagination.PageNumber - 1) * pagination.PageSize)

	// Run query and get the cursor for the next page
	var entities []datastore.PropertyList
	it := DatastoreClient.Run(datastoreClient, ctx, q)
	keys := make([]*datastore.Key, 0)
	for {
		var entity datastore.PropertyList
		key, err := it.Next(&entity)
		if err == datastore.ErrNoSuchEntity {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("failed to iterate results: %v", err)
		}
		entities = append(entities, entity)
		keys = append(keys, key)
	}

	// Get the next page cursor
	cursor, err := it.Cursor()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get cursor: %v", err)
	}

	// Convert properties to map and add IDs
	results := make([]map[string]interface{}, len(entities))
	for i, entity := range entities {
		data := ConvertPropertiesToMap(entity)
		data["id"] = ParseID(keys[i].ID)
		results[i] = data
	}

	return results, &cursor, nil
}