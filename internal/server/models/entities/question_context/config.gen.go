package question_context

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "question_contexts"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity QuestionContext) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "question_contexts/[questionContextID]"
func GetDocumentPath(questionContextID string) string {
	output := documentPath
	output = strings.Replace(output, "[questionContextID]", questionContextID, -1)
	return output
}

func (entity QuestionContext) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity QuestionContext) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity QuestionContext) GetID() string {
	return entity.ID
}
