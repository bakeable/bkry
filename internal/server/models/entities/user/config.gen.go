package user

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "users"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity User) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "users/[userID]"
func GetDocumentPath(userID string) string {
	output := documentPath
	output = strings.Replace(output, "[userID]", userID, -1)
	return output
}

func (entity User) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity User) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity User) GetID() string {
	return entity.ID
}
