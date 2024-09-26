package batch_export

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "BatchExport"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity BatchExport) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "BatchExport/[batchExportID]"
func GetDocumentPath(batchExportID string) string {
	output := documentPath
	output = strings.Replace(output, "[batchExportID]", batchExportID, -1)
	return output
}

func (entity BatchExport) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity BatchExport) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity BatchExport) GetID() string {
	return entity.ID
}
