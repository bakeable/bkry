package user_profile

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "UserProfile"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity UserProfile) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "UserProfile/[userProfileID]"
func GetDocumentPath(userProfileID string) string {
	output := documentPath
	output = strings.Replace(output, "[userProfileID]", userProfileID, -1)
	return output
}

func (entity UserProfile) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity UserProfile) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity UserProfile) GetID() string {
	return entity.ID
}
