package attribute

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "Attribute"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Attribute) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "Attribute/[attributeID]"
func GetDocumentPath(attributeID string) string {
	output := documentPath
	output = strings.Replace(output, "[attributeID]", attributeID, -1)
	return output
}

func (entity Attribute) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Attribute) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity Attribute) GetID() string {
	return entity.ID
}
