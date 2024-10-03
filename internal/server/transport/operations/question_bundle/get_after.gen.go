package question_bundle_operations

import (
	question_bundle "github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
)

func afterGet(questionBundleID string, entity question_bundle.QuestionBundle) question_bundle.QuestionBundle {
	return entity
}