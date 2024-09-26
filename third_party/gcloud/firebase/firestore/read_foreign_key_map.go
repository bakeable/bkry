package db

import (
	"context"
	"log"
)

func ReadForeignKeyMap(documentPath string) (map[string]interface{}, error) {
	doc, err := firestoreClient.Doc(documentPath).Get(context.Background())
	if err != nil {
		return nil, handleError(err, "Failed to read document", &documentPath, nil)
	}

	data := doc.Data()
	value, ok := data["foreign"]
	if !ok {
		log.Printf("Field 'foreign' not found in document")
		return nil, nil
	}

	if value == nil {
		return nil, nil
	}

	return value.(map[string]interface{}), nil
}