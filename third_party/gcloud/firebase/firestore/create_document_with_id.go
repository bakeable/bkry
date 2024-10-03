package firestore

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
)

func CreateDocumentWithId(documentPath string, data map[string]interface{}, userId *string) (string, error) {	
	// Set id
	parts := strings.Split(documentPath, "/")
	data["id"] = parts[len(parts) - 1]

	// Set audit info
	data["created"] = audit(userId)
	data["modified"] = audit(userId)

	_, err := firestoreClient.Doc(documentPath).Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return "", handleError(err, "Failed to create document with ID", &documentPath, nil)
	}

	// Update metadata
	onDocumentAdded(documentPath)

	return data["id"].(string), nil
}