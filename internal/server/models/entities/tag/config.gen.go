package tag

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "tags"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Tag) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "tags/[tagID]"
func GetDocumentPath(tagID string) string {
	output := documentPath
	output = strings.Replace(output, "[tagID]", tagID, -1)
	return output
}

func (entity Tag) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Tag) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity Tag) GetID() string {
	return entity.ID
}
