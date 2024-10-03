package question_bundle

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "question_bundles"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity QuestionBundle) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "question_bundles/[questionBundleID]"
func GetDocumentPath(questionBundleID string) string {
	output := documentPath
	output = strings.Replace(output, "[questionBundleID]", questionBundleID, -1)
	return output
}

func (entity QuestionBundle) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity QuestionBundle) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity QuestionBundle) GetID() string {
	return entity.ID
}
