package examination_settings

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "ExaminationSettings"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity ExaminationSettings) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "ExaminationSettings/[examinationSettingsID]"
func GetDocumentPath(examinationSettingsID string) string {
	output := documentPath
	output = strings.Replace(output, "[examinationSettingsID]", examinationSettingsID, -1)
	return output
}

func (entity ExaminationSettings) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity ExaminationSettings) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity ExaminationSettings) GetID() string {
	return entity.ID
}
