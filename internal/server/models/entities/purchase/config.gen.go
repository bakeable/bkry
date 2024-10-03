package purchase

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "users/[userID]/purchases"
func GetCollectionPath(userID string) string {
	output := collectionPath
	output = strings.Replace(output, "[userID]", userID, -1)
	return output
}

func (entity Purchase) GetCollectionPath() string {
	return GetCollectionPath(entity.UserID)
}

var documentPath = "users/[userID]/purchases/[purchaseID]"
func GetDocumentPath(userID string, purchaseID string) string {
	output := documentPath
	output = strings.Replace(output, "[userID]", userID, -1)
	output = strings.Replace(output, "[purchaseID]", purchaseID, -1)
	return output
}

func (entity Purchase) GetDocumentPath() string {
    return GetDocumentPath(entity.UserID,  entity.ID)
}

func (entity Purchase) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity Purchase) GetID() string {
	return entity.ID
}
