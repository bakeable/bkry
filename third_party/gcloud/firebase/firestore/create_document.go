package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// CreateDocument creates a new document in the specified collection with the given data.
// It returns an error if the document creation fails.
func CreateDocument(collectionPath string, data map[string]interface{}, userId *string) (string, error) {

	// Set audit info
	data["created"] = audit(userId)

	// Make sure timestamps are equal at creation
	data["modified"] = data["created"]

	docRef, _, err := firestoreClient.Collection(collectionPath).Add(context.Background(), data)
	if err != nil {
		return "", handleError(err, "Failed to create document", &collectionPath, nil)
	}
	log.Println("Document written with ID: ", docRef.ID)

	// Get document ID
	documentID := docRef.ID

	// Update metadata
	onDocumentAdded(collectionPath + "/" + documentID)

	// Write ID to document
	_, err = firestoreClient.Doc(collectionPath + "/" + documentID).Set(context.Background(), map[string]interface{}{"id": documentID}, firestore.MergeAll)
	if err != nil {
		return "", handleError(err, "Failed to write ID to document", &collectionPath, nil)
	}


	return documentID, nil
}