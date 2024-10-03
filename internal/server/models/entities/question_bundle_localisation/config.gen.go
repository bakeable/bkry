package question_bundle_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "question_bundles/[questionBundleID]/question_bundle_localisations"
func GetCollectionPath(questionBundleID string) string {
	output := collectionPath
	output = strings.Replace(output, "[questionBundleID]", questionBundleID, -1)
	return output
}

func (entity QuestionBundleLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.QuestionBundleID)
}

var documentPath = "question_bundles/[questionBundleID]/question_bundle_localisations/[questionBundleLocalisationID]"
func GetDocumentPath(questionBundleID string, questionBundleLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[questionBundleID]", questionBundleID, -1)
	output = strings.Replace(output, "[questionBundleLocalisationID]", questionBundleLocalisationID, -1)
	return output
}

func (entity QuestionBundleLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.QuestionBundleID,  entity.ID)
}

func (entity QuestionBundleLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity QuestionBundleLocalisation) GetID() string {
	return entity.ID
}
