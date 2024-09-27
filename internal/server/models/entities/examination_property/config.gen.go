package examination_property

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "ExaminationProperty"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity ExaminationProperty) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "ExaminationProperty/[examinationPropertyID]"
func GetDocumentPath(examinationPropertyID string) string {
	output := documentPath
	output = strings.Replace(output, "[examinationPropertyID]", examinationPropertyID, -1)
	return output
}

func (entity ExaminationProperty) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity ExaminationProperty) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity ExaminationProperty) GetID() string {
	return entity.ID
}
