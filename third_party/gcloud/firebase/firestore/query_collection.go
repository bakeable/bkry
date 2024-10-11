package firestore

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bakeable/bkry/internal/server/database"
)

func QueryCollection(collectionPath string, queries []database.Query, pagination database.Pagination) ([]map[string]interface{}, database.Pagination, error) {
	// Query documents by fields
	query := firestoreClient.Collection(collectionPath).Query

	// Build query
	for i := 0; i < len(queries); i++ {
		query = query.Where(queries[i].Field, string(queries[i].Operator), queries[i].Value)
	}
	fmt.Println("Querying collection", collectionPath, "with queries", queries, "and pagination", pagination)
	if pagination.PageSize > 0 {
		query = query.OrderBy(pagination.OrderBy, sortDirection(pagination.Ascending)).Offset(pagination.PageSize * (pagination.PageNumber - 1)).Limit(pagination.PageSize)
	}

	// Perform query
	docs, err := query.Documents(context.Background()).GetAll()
	log.Println("Found", len(docs), "documents in the query")
	if err != nil {
		return nil, pagination, handleError(err, "Failed to query documents", &collectionPath, collectQueries(queries))
	}


	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	// Get collection name
	collectionPathParts := strings.Split(collectionPath, "/")
	collectionName := collectionPathParts[len(collectionPathParts)-1]

	// Get parent reference
	parentRef := firestoreClient.Collection(collectionPath).Parent
	if parentRef == nil {
		parentRef = firestoreClient.Doc("general/metadata") // Root collection
	}

	// Get document
	doc, _ := parentRef.Get(context.Background())

	// Get total count
	d := doc.Data()
	if _metadata, ok := d["_metadata"]; ok {
		if metadata, ok := _metadata.(map[string]interface{}); ok {
			if documentCounts, ok := metadata["document_count"]; ok {
				if documentCounts, ok := documentCounts.(map[string]interface{}); ok {
					if count, ok := documentCounts[collectionName]; ok {
						pagination.Count = int(count.(int64))
					}
				}
			}
		}
	}

	// Calculate total pages
	pagination.TotalPages = int(float64(pagination.Count) / float64(pagination.PageSize))
	pagination.HasNextPage = pagination.PageNumber < pagination.TotalPages
	pagination.HasPreviousPage = pagination.PageNumber > 1

	return data, pagination, nil
}

