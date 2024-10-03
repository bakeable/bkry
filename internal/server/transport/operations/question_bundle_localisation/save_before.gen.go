package question_bundle_localisation_operations

import (
	question_bundle_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
)

func beforeSave(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) question_bundle_localisation.QuestionBundleLocalisation {
	// Return QuestionBundleLocalisation
	return entity
}