package firestore

import (
	"context"
	"fmt"
	"log"

	"github.com/bakeable/bkry/internal/server/database"
)

func QueryCollection(collectionPath string, queries []database.Query, pagination database.Pagination) ([]map[string]interface{}, error) {
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
		return nil, handleError(err, "Failed to query documents", &collectionPath, collectQueries(queries))
	}


	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	return data, nil
}

