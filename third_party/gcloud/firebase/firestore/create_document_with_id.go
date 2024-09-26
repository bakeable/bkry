package db

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
)

func CreateDocumentWithId(documentPath string, data map[string]interface{}, userId *string) (string, error) {	
	// Set timestamps
	data["updated"] = firestore.ServerTimestamp
	data["created"] = firestore.ServerTimestamp

	// Set id
	parts := strings.Split(documentPath, "/")
	data["id"] = parts[len(parts) - 1]

	// Set editor
	if userId == nil {
		data["updated_by"] = "system"
		data["created_by"] = "system"
	} else {
		data["updated_by"] = *userId
		data["created_by"] = *userId
	}

	_, err := firestoreClient.Doc(documentPath).Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return "", handleError(err, "Failed to create document with ID", &documentPath, nil)
	}

	return data["id"].(string), nil
}