package media

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "media"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Media) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "media/[mediaID]"
func GetDocumentPath(mediaID string) string {
	output := documentPath
	output = strings.Replace(output, "[mediaID]", mediaID, -1)
	return output
}

func (entity Media) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Media) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity Media) GetID() string {
	return entity.ID
}
