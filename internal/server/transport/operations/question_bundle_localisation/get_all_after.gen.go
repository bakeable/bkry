package question_bundle_localisation_operations

import (
	question_bundle_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
)


func afterGetAll(questionBundleID string, entities []question_bundle_localisation.QuestionBundleLocalisation) []question_bundle_localisation.QuestionBundleLocalisation {
	return entities
}