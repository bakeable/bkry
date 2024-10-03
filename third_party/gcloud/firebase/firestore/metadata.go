package firestore

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
)

func onDocumentAdded(documentPath string) {
	// Get parent path
	info, err := extractInfo(documentPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = firestoreClient.Doc(info.ParentPath).Set(context.Background(), map[string]interface{}{
		"_metadata": map[string]interface{}{
			"document_count": map[string]interface{}{
				info.CollectionName: firestore.Increment(1),
			},
			"last_added": map[string]interface{}{
				info.CollectionName: firestore.ServerTimestamp,
			},
		},
	}, firestore.MergeAll)
}

func onDocumentRemoved(documentPath string) {
	// Get parent path
	info, err := extractInfo(documentPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = firestoreClient.Doc(info.ParentPath).Set(context.Background(), map[string]interface{}{
		"_metadata": map[string]interface{}{
			"document_count": map[string]interface{}{
				info.CollectionName: firestore.Increment(-1),
			},
			"last_deleted": map[string]interface{}{
				info.CollectionName: firestore.ServerTimestamp,
			},
		},
	}, firestore.MergeAll)
}

func onCollectionRemoved(collectionPath string) {
	// Get parent path
	info, err := extractInfo(collectionPath + "/id-not-used")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = firestoreClient.Doc(info.ParentPath).Set(context.Background(), map[string]interface{}{
		"_metadata": map[string]interface{}{
			"document_count": map[string]interface{}{
				info.CollectionName: 0,
			},
			"last_deleted": map[string]interface{}{
				info.CollectionName: firestore.ServerTimestamp,
			},
		},
	}, firestore.MergeAll)
}


type DocumentInfo struct {
	ParentPath string
	CollectionName string
	DocumentId string
}
func extractInfo(documentPath string) (DocumentInfo, error) {
	parts := strings.Split(documentPath, "/")

	if len(parts) < 2 {
		return DocumentInfo{}, fmt.Errorf("Invalid document path: %s", documentPath)
	} else if len(parts) == 2 {
		return DocumentInfo{
			ParentPath: "general/metadata",
			CollectionName: parts[0],
			DocumentId: parts[1],
		}, nil
	} else {
		return DocumentInfo{
			ParentPath: strings.Join(parts[:len(parts)-2], "/"),
			CollectionName: parts[len(parts)-2],
			DocumentId: parts[len(parts)-1],
		}, nil
	}
}