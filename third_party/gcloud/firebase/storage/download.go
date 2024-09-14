package firebase_storage

import (
	"context"
	"fmt"
	"io"
	"os"
)

func Download(filePath string) (string, error) {
	// Get bucket
	bucketName := os.Getenv("GO_FIREBASE_STORAGE_BUCKET")

	// Get bucket
	bucket, err := firebaseStorageClient.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error getting bucket", err)
		return "", err
	}

	// Get filepath
	ctx := context.Background()
	reader, err := bucket.Object(filePath).NewReader(ctx)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	// Read Data
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading data", err)
		return "", err
	}

	// Save to local tmp folder, create folders if nonexistent
	err = os.MkdirAll("/tmp" + filePath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folders", err)
		return "", err
	}

	// Write to file
	err = os.WriteFile("/tmp/" + filePath, data, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing file", err)
		return "", err
	}

	return "/tmp/" + filePath, nil
}