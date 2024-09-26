package db

import (
	"context"
	"log"
)

func DeleteCollection(path string) error {
	// Get all documents in collection
	docs, err := ReadDocuments(path)
	if err != nil {
		log.Printf("Failed to read documents: %v", err)
		return err
	}

	// Loop through documents
	for _, doc := range docs {
		docRef := firestoreClient.Doc(path + "/" + doc["id"].(string))

		// Get subcollections of document
		subcollections, err := docRef.Collections(context.Background()).GetAll()
		if err != nil {
			log.Printf("Failed to get subcollections: %v", err)
		}

		// Delete document
		_, err = docRef.Delete(context.Background())
		if err != nil {
			log.Printf("Failed to delete document: %v", err)
		}

		// Delete subcollections
		for _, subcollection := range subcollections {
			err = DeleteCollection(path + "/" + doc["id"].(string) + "/" + subcollection.ID)
			if err != nil {
				log.Printf("Failed to delete subcollection: %v", err)
			}
		}
	}


	return nil
}