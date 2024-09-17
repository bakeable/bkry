package email

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "Email"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Email) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "Email/[emailID]"
func GetDocumentPath(emailID string) string {
	output := documentPath
	output = strings.Replace(output, "[emailID]", emailID, -1)
	return output
}

func (entity Email) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Email) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity Email) GetID() string {
	return entity.ID
}
