package user

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "User"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity User) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "User/[userID]"
func GetDocumentPath(userID string) string {
	output := documentPath
	output = strings.Replace(output, "[userID]", userID, -1)
	return output
}

func (entity User) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity User) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity User) GetID() string {
	return entity.ID
}
