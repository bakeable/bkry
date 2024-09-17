package attribute_option_set

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AttributeOptionSet"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AttributeOptionSet) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AttributeOptionSet/[attributeOptionSetID]"
func GetDocumentPath(attributeOptionSetID string) string {
	output := documentPath
	output = strings.Replace(output, "[attributeOptionSetID]", attributeOptionSetID, -1)
	return output
}

func (entity AttributeOptionSet) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AttributeOptionSet) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AttributeOptionSet) GetID() string {
	return entity.ID
}
