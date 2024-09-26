package db

import (
	"context"
)

func QueryCollection(collectionPath string, queries []Query) ([]map[string]interface{}, error) {
	// Query documents by fields
	query := firestoreClient.Collection(collectionPath).Where(queries[0].Field, queries[0].Operator, queries[0].Value)

	// Build query
	for i := 1; i < len(queries); i++ {
		query = query.Where(queries[i].Field, queries[i].Operator, queries[i].Value)
	}

	// Perform query
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, handleError(err, "Failed to query documents", &collectionPath, collectQueries(queries))
	}


	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	return data, nil
}

