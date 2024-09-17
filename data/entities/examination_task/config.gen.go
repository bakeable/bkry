package examination_task

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "ExaminationTask"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity ExaminationTask) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "ExaminationTask/[examinationTaskID]"
func GetDocumentPath(examinationTaskID string) string {
	output := documentPath
	output = strings.Replace(output, "[examinationTaskID]", examinationTaskID, -1)
	return output
}

func (entity ExaminationTask) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity ExaminationTask) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity ExaminationTask) GetID() string {
	return entity.ID
}
