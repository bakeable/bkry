package question_bundle_operations

import (
	question_bundle "github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
)

func beforeSave(entity question_bundle.QuestionBundle, editorID *string) question_bundle.QuestionBundle {
	// Return QuestionBundle
	return entity
}