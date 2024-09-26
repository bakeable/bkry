package api_key

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "ApiKey"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity ApiKey) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "ApiKey/[apiKeyID]"
func GetDocumentPath(apiKeyID string) string {
	output := documentPath
	output = strings.Replace(output, "[apiKeyID]", apiKeyID, -1)
	return output
}

func (entity ApiKey) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity ApiKey) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity ApiKey) GetID() string {
	return entity.ID
}
