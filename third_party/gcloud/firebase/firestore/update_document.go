package firestore

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
)


func UpdateDocument(path string, data map[string]interface{}, userId *string) error {
	// Set timestamps
	data["modified"] = audit(userId)

	_, err := firestoreClient.Doc(path).Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return handleError(err, "Failed to update document", &path, nil)
	}

	
	if _, ok := data["id"]; !ok || !strings.Contains(path, data["id"].(string)) {
		// Get ID from the document reference and update it
		documentID := path[strings.LastIndex(path, "/") + 1:]

		// Write ID to document
		_, err = firestoreClient.Doc(path).Set(context.Background(), map[string]interface{}{"id": documentID}, firestore.MergeAll)
		if err != nil {
			return handleError(err, "Failed to write ID to document", &path, nil)
		}

	}

	return nil
}