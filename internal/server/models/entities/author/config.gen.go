package author

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "authors"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Author) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "authors/[authorID]"
func GetDocumentPath(authorID string) string {
	output := documentPath
	output = strings.Replace(output, "[authorID]", authorID, -1)
	return output
}

func (entity Author) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Author) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity Author) GetID() string {
	return entity.ID
}
