package firestore

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func DeleteCollection(path string) error {
	fmt.Println("Deleting collection: " + path)

	// Get documents in collection
	docs, err := firestoreClient.Collection(path).Documents(context.Background()).GetAll()
	if err != nil {
		return handleError(err, "Failed to read documents", &path, nil)
	}

	// Loop through documents
	for _, doc := range docs {
		// Get subcollections of document
		subcollections, err := doc.Ref.Collections(context.Background()).GetAll()
		if err != nil {
			log.Printf("Failed to get subcollections: %v", err)
		}

		// Delete document
		_, err = doc.Ref.Delete(context.Background())
		if err != nil {
			log.Printf("Failed to delete document: %v", err)
		}

		// Delete subcollections
		if len(subcollections) > 0 && subcollections != nil {
			for _, collection := range subcollections {
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

	}

	// Update metadata
	onCollectionRemoved(path)


	return nil
}