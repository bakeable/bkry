package category

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "categories"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Category) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "categories/[categoryID]"
func GetDocumentPath(categoryID string) string {
	output := documentPath
	output = strings.Replace(output, "[categoryID]", categoryID, -1)
	return output
}

func (entity Category) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Category) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity Category) GetID() string {
	return entity.ID
}
