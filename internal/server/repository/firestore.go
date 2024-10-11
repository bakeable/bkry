package repository

import (
	"log"

	"github.com/bakeable/bkry/internal/server/database"
	"github.com/bakeable/bkry/third_party/gcloud/firebase/firestore"
)

func NewFirestoreRepository[T Entity]() *FirestoreRepository[T] {
	return &FirestoreRepository[T]{}
}

// Read retrieves an entity from the data storage.
func (repo FirestoreRepository[T]) Read(path string) (map[string]interface{}, error) {
	doc, err := firestore.ReadDocument(path)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Find returns a  map[string]interface{} representing a single document.
func (repo *FirestoreRepository[T]) Find(path string, queries []database.Query) (map[string]interface{}, error) {
    docs, _, err := firestore.QueryCollection(path, queries, database.Pagination{
		PageSize: 1,
		PageNumber: 1,
		OrderBy: "id",
		Ascending: true,
	})
    if err != nil || len(docs) == 0 {
        return nil, err
    }
	
    return docs[0], nil
}


// ReadAll returns a slice of map[string]interface{} representing multiple documents.
func (repo *FirestoreRepository[T]) ReadAll(path string) ([]map[string]interface{}, error) {
    docs, err := firestore.ReadDocuments(path)
    if err != nil {
        return nil, err
    }
	
    return docs, nil
}

// ReadPaginated returns a slice of map[string]interface{} representing multiple paginated documents.
func (repo *FirestoreRepository[T]) ReadPaginated(path string, pagination database.Pagination) ([]map[string]interface{}, database.Pagination, error) {
    docs, pagination, err := firestore.ReadDocumentsPaginated(path, pagination)
    if err != nil {
        return nil, pagination, err
    }
	
    return docs, pagination, nil
}

// Query returns a slice of map[string]interface{} representing multiple documents.
func (repo *FirestoreRepository[T]) Query(path string, queries []database.Query, pagination database.Pagination) ([]map[string]interface{}, database.Pagination, error) {
    docs, pagination, err := firestore.QueryCollection(path, queries, pagination)
    if err != nil {
        return nil, pagination, err
    }
	
    return docs, pagination, nil
}

// QueryGroup returns a slice of map[string]interface{} representing multiple documents.
func (repo *FirestoreRepository[T]) QueryGroup(collectionGroup string, queries []database.Query) ([]map[string]interface{}, error) {
    docs, err := firestore.QueryCollectionGroup(collectionGroup, queries)
    if err != nil {
        return nil, err
    }
	
    return docs, nil
}

// Create adds an entity to the data storage.
func (repo *FirestoreRepository[T]) Create(entity T, editorId *string) (string, error) {
	// Get encoded entity
	encodedEntity := entity.ToMap()

	if entity.GetID() == "" {
		id, err := firestore.CreateDocument(entity.GetCollectionPath(), encodedEntity, editorId)
		if err != nil {
			return "", err
		}
		log.Println("Created document with ID: ", id)
		return id, nil
	} else {
		id, err := firestore.CreateDocumentWithId(entity.GetDocumentPath(), encodedEntity, editorId)
		if err != nil {
			return "", err
		}
		return id, nil
	}
}

// Update modifies an entity in the data storage.
func (repo *FirestoreRepository[T]) Update(entity T, editorId *string) (string, error) {
	// Get encoded entity
	encodedEntity := entity.ToMap()

	// Update entity
	err := firestore.UpdateDocument(entity.GetDocumentPath(), encodedEntity, editorId)
	if err != nil {
		return "", err
	}
	return entity.GetID(), nil
}

// Delete removes an entity from the data storage.
func (repo *FirestoreRepository[T]) Delete(path string) error {
	return firestore.DeleteDocument(path)
}
