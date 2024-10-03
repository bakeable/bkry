package question_context_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "question_contexts/[questionContextID]/question_context_localisations"
func GetCollectionPath(questionContextID string) string {
	output := collectionPath
	output = strings.Replace(output, "[questionContextID]", questionContextID, -1)
	return output
}

func (entity QuestionContextLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.QuestionContextID)
}

var documentPath = "question_contexts/[questionContextID]/question_context_localisations/[questionContextLocalisationID]"
func GetDocumentPath(questionContextID string, questionContextLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[questionContextID]", questionContextID, -1)
	output = strings.Replace(output, "[questionContextLocalisationID]", questionContextLocalisationID, -1)
	return output
}

func (entity QuestionContextLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.QuestionContextID,  entity.ID)
}

func (entity QuestionContextLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity QuestionContextLocalisation) GetID() string {
	return entity.ID
}
