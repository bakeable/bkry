package repository

import (
	"fmt"

	"github.com/bakeable/bkry/internal/server/database"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
)

// NewDatastoreRepository creates a new repository instance.
func NewDatastoreRepository[T Entity]() *DatastoreRepository[T] {
	return &DatastoreRepository[T]{}
}

// Read retrieves an entity from the data storage.
func (repo DatastoreRepository[T]) Read(path string) (map[string]interface{}, error) {
	kind, id, ancestor, err := datastore.DeriveKeyAndAncestorKey(path)
	if err != nil {
		return nil, err
	}

	entity, err := datastore.GetEntityById(kind, id, ancestor)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Find returns a map[string]interface{} representing a single document.
func (repo *DatastoreRepository[T]) Find(path string, queries []database.Query) (map[string]interface{}, error) {
	kind, ancestor, err := datastore.DeriveKindAndAncestorKey(path)
	if err != nil {
		return nil, err
	}

	entity, err := datastore.QueryEntity(kind, queries, nil, ancestor)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// ReadAll returns a slice of map[string]interface{} representing multiple documents.
func (repo *DatastoreRepository[T]) ReadAll(path string) ([]map[string]interface{}, error) {
	kind, ancestor, err := datastore.DeriveKindAndAncestorKey(path)
	if err != nil {
		return nil, err
	}

	entities, err := datastore.GetAllEntitiesByKind(kind, ancestor)
	if err != nil {
		return nil, err
	}

	fmt.Println("Found entities:", len(entities))

	return entities, nil
}

// ReadPaginated returns a slice of map[string]interface{} representing multiple paginated documents.
func (repo *DatastoreRepository[T]) ReadPaginated(path string, pagination database.Pagination) ([]map[string]interface{}, database.Pagination, error) {
	kind, ancestor, err := datastore.DeriveKindAndAncestorKey(path)
	if err != nil {
		return nil, pagination, err
	}

	entities, cursor, err := datastore.GetPaginatedEntitiesByKind(kind, pagination, ancestor)
	if err != nil {
		return nil, pagination, err
	}

	// Update the pagination object with the cursor and count information
	pagination.TotalPages = len(entities) / pagination.PageSize
	if len(entities)%pagination.PageSize != 0 {
		pagination.TotalPages++
	}

	// Add cursor to pagination object
	pagination.Cursor = cursor.String()

	return entities, pagination, nil
}

// Query returns a slice of map[string]interface{} representing multiple documents.
func (repo *DatastoreRepository[T]) Query(path string, queries []database.Query, pagination database.Pagination) ([]map[string]interface{}, error) {
	kind, ancestor, err := datastore.DeriveKindAndAncestorKey(path)
	if err != nil {
		return nil, err
	}

	entities, err := datastore.QueryEntities(kind, queries, &pagination, ancestor)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// QueryGroup returns a slice of map[string]interface{} representing multiple documents.
func (repo *DatastoreRepository[T]) QueryGroup(collectionGroup string, queries []database.Query) ([]map[string]interface{}, error) {
	entities, err := datastore.QueryEntities(collectionGroup, queries, nil, nil)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// Create adds an entity to the data storage.
func (repo *DatastoreRepository[T]) Create(entity T, editorId *string) (string, error) {
	// Get encoded entity
	encodedEntity := entity.ToMap()

	// Derive path
	documentPath := entity.GetDocumentPath()
	kind, id, ancestor, err := datastore.DeriveKeyAndAncestorKey(documentPath)
	if err != nil {
		return "", err
	}

	if id == "" {
		// Create new entity
		key, err := datastore.CreateEntity(kind, encodedEntity, ancestor)
		if err != nil {
			return "", err
		}
		return key.Name, nil
	} else {
		// Create entity with specified ID
		key, err := datastore.UpdateEntityById(kind, id, encodedEntity, ancestor)
		if err != nil {
			return "", err
		}
		return key.Name, nil
	}
}

// Update modifies an entity in the data storage.
func (repo *DatastoreRepository[T]) Update(entity T, editorId *string) (string, error) {
	// Get encoded entity
	encodedEntity := entity.ToMap()

	// Derive path
	documentPath := entity.GetDocumentPath()
	kind, id, ancestor, err := datastore.DeriveKeyAndAncestorKey(documentPath)
	if err != nil {
		return "", err
	}

	// Update entity
	key, err := datastore.UpdateEntityById(kind, id, encodedEntity, ancestor)
	if err != nil {
		return "", err
	}
	return key.Name, nil
}

// Delete removes an entity from the data storage.
func (repo *DatastoreRepository[T]) Delete(path string) error {
	kind, id, ancestor, err := datastore.DeriveKeyAndAncestorKey(path)
	if err != nil {
		return err
	}

	return datastore.DeleteEntityById(kind, id, ancestor)
}
