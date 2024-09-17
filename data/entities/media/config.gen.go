package media

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "Media"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Media) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "Media/[mediaID]"
func GetDocumentPath(mediaID string) string {
	output := documentPath
	output = strings.Replace(output, "[mediaID]", mediaID, -1)
	return output
}

func (entity Media) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Media) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity Media) GetID() string {
	return entity.ID
}
