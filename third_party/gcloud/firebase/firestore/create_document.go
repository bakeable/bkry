package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

// CreateDocument creates a new document in the specified collection with the given data.
// It returns an error if the document creation fails.
func CreateDocument(collectionPath string, data map[string]interface{}, userId *string) (string, error) {
	// Set timestamps
	data["updated"] = firestore.ServerTimestamp
	data["created"] = firestore.ServerTimestamp

	// Set id
	if userId == nil {
		data["updated_by"] = "system"
		data["created_by"] = "system"
	} else {
		data["updated_by"] = *userId
		data["created_by"] = *userId
	}


	docRef, _, err := firestoreClient.Collection(collectionPath).Add(context.Background(), data)
	if err != nil {
		return "", handleError(err, "Failed to create document", &collectionPath, nil)
	}

	documentID := docRef.ID

	// Write ID to document
	_, err = firestoreClient.Doc(collectionPath + "/" + documentID).Set(context.Background(), map[string]interface{}{"id": documentID}, firestore.MergeAll)
	if err != nil {
		return "", handleError(err, "Failed to write ID to document", &collectionPath, nil)
	}

	return documentID, nil
}