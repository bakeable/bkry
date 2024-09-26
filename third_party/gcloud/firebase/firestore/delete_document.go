package db

import (
	"context"
)

func DeleteDocument(path string) error {
	_, err := firestoreClient.Doc(path).Delete(context.Background())
	if err != nil {
		return handleError(err, "Failed to delete document", &path, nil)
	}

	return nil
}