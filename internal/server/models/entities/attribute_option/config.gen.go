package attribute_option

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AttributeOption"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AttributeOption) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AttributeOption/[attributeOptionID]"
func GetDocumentPath(attributeOptionID string) string {
	output := documentPath
	output = strings.Replace(output, "[attributeOptionID]", attributeOptionID, -1)
	return output
}

func (entity AttributeOption) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AttributeOption) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AttributeOption) GetID() string {
	return entity.ID
}
