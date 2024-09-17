package cloud_function

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "CloudFunction"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity CloudFunction) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "CloudFunction/[cloudFunctionID]"
func GetDocumentPath(cloudFunctionID string) string {
	output := documentPath
	output = strings.Replace(output, "[cloudFunctionID]", cloudFunctionID, -1)
	return output
}

func (entity CloudFunction) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity CloudFunction) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity CloudFunction) GetID() string {
	return entity.ID
}
