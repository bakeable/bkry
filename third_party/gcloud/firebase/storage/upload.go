package firebase_storage

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
)

func Upload(filePath string, data []byte) (string, error) {
	// Get bucket
	bucket, err := firebaseStorageClient.Bucket(os.Getenv("GO_FIREBASE_STORAGE_BUCKET"))
	if err != nil {
		return "", err
	}

	// Get filepath
	ctx := context.Background()
	writer := bucket.Object(filePath).NewWriter(ctx)
	defer writer.Close()

	// Write data to the object
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error writing data", err)
		return "", err
	}

	// Close the writer to flush the data to the object
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer", err)
		return "", err
	}

	// Make public
	err = bucket.Object(filePath).ACL().Set(ctx, storage.AllUsers, storage.RoleReader)
	if err != nil {
		fmt.Println("Error making file public", err)
		return "", err
	}

	// Get download url
	meta, err := bucket.Object(filePath).Attrs(ctx) 
	if err != nil {
		fmt.Println("Error getting metadata", err)
		return "", err
	}

	// Extract link
	fileUrl := meta.MediaLink

	return fileUrl, nil
}