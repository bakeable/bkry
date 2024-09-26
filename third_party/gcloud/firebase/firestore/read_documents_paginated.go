package db

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
)

func sortDirection(sortDesc bool) firestore.Direction {
	if sortDesc {
		return firestore.Desc
	} else {
		return firestore.Asc
	}
}

func ReadDocumentsPaginated(collectionPath string, pageSize int, pageNumber int, sortBy string, sortDesc bool) ([]map[string]interface{}, error) {
	query := firestoreClient.Collection(collectionPath).OrderBy(sortBy, sortDirection(sortDesc)).Offset(pageSize * (pageNumber - 1)).Limit(pageSize)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, handleError(err, "Failed to read documents", &collectionPath, &[]string{
			"pageSize: " + strconv.Itoa(pageSize),
			"pageNumber: " + strconv.Itoa(pageNumber),
			"sortBy: " + sortBy,
			"sortDesc: " + strconv.FormatBool(sortDesc),
		})
	}

	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	return data, nil
}
