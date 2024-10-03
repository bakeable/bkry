package firestore

import (
	"context"
	"fmt"
	"strings"
)

func DeleteDocument(path string) error {
	fmt.Println("Deleting document: " + path)
	// Delete child collections
	collections, _ := firestoreClient.Doc(path).Collections(context.Background()).GetAll()
	if len(collections) > 0 && collections != nil {
		for _, collection := range collections {
			if collection != nil {
				// Get the path after "/documents/"
				parts := strings.Split(collection.Path, "/documents/")
				if len(parts) > 1 {
					// Get the path after "/documents/"
					path := parts[1]
					DeleteCollection(path)
				}
			}
		}
	}

	// Delete document
	_, err := firestoreClient.Doc(path).Delete(context.Background())
	if err != nil {
		return handleError(err, "Failed to delete document", &path, nil)
	}

	// Update metadata
	onDocumentRemoved(path)

	return nil
}
