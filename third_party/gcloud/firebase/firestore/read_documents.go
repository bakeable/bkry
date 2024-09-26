package db

import (
	"context"
)


func ReadDocuments(collectionPath string) ([]map[string]interface{}, error) {
	docs, err := firestoreClient.Collection(collectionPath).Documents(context.Background()).GetAll()
	if err != nil {
		return nil, handleError(err, "Failed to read documents", &collectionPath, nil)
	}

	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	return data, nil
}