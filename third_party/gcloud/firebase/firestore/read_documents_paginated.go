package firestore

import (
	"context"
	"log"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/bakeable/bkry/internal/server/database"
)

func sortDirection(sortDesc bool) firestore.Direction {
	if sortDesc {
		return firestore.Asc
	} else {
		return firestore.Desc
	}
}

func ReadDocumentsPaginated(collectionPath string, pagination database.Pagination) ([]map[string]interface{}, database.Pagination, error) {
	log.Println("Reading documents from collection", collectionPath, "with pagination", pagination)
	query := firestoreClient.Collection(collectionPath).OrderBy(pagination.OrderBy, sortDirection(pagination.Ascending)).Offset(pagination.PageSize * (pagination.PageNumber - 1)).Limit(pagination.PageSize)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, pagination, handleError(err, "Failed to read documents", &collectionPath, &[]string{
			"pageSize: " + strconv.Itoa(pagination.PageSize),
			"pageNumber: " + strconv.Itoa(pagination.PageNumber),
			"orderBy: " + pagination.OrderBy,
			"descending: " + strconv.FormatBool(pagination.Ascending),
		})
	}
	var data []map[string]interface{}
	for _, doc := range docs {
		data = append(data, doc.Data())
	}

	// Get collection name
	collectionPathParts := strings.Split(collectionPath, "/")
	collectionName := collectionPathParts[len(collectionPathParts)-1]

	// Get parent reference
	parentRef := firestoreClient.Collection(collectionPath).Parent
	if parentRef == nil {
		parentRef = firestoreClient.Doc("general/metadata") // Root collection
	}

	// Get document
	doc, _ := parentRef.Get(context.Background())

	// Get total count
	d := doc.Data()
	if _metadata, ok := d["_metadata"]; ok {
		if metadata, ok := _metadata.(map[string]interface{}); ok {
			if documentCounts, ok := metadata["document_count"]; ok {
				if documentCounts, ok := documentCounts.(map[string]interface{}); ok {
					if count, ok := documentCounts[collectionName]; ok {
						pagination.Count = int(count.(int64))
					}
				}
			}
		}
	}

	// Calculate total pages
	pagination.TotalPages = int(float64(pagination.Count) / float64(pagination.PageSize))

	return data, pagination, nil
}
