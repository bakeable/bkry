package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

func UpdateForeignKeyMap(documentPath string, foreignKeyMap map[string]interface{}) error {
	doc, err := firestoreClient.Doc(documentPath).Get(context.Background())
	if err != nil {
		return handleError(err, "Failed to read foreign key map", &documentPath, nil)
	}

	data := doc.Data()
	
	data["foreign"] = foreignKeyMap

	_, err = firestoreClient.Doc(documentPath).Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return handleError(err, "Failed to update foreign key map", &documentPath, nil)
	}

	return nil
}
