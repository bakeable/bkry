package firestore

import (
	"context"
)



func ReadDocument(path string) (map[string]interface{}, error) {
	doc, err := firestoreClient.Doc(path).Get(context.Background())
	if err != nil {
		return nil, handleError(err, "Failed to read document", &path, nil)
	}

	data := doc.Data()
	return data, nil
}