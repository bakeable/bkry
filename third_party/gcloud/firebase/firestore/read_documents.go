package firestore

import (
	"context"
	"log"
)


func ReadDocuments(collectionPath string) ([]map[string]interface{}, error) {
	log.Println("Reading documents from collection", collectionPath)
	if collectionPath == "" {
		return nil, handleError(nil, "Collection path is empty", &collectionPath, nil)
	}

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