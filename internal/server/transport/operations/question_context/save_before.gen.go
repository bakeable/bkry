package question_context_operations

import (
	question_context "github.com/bakeable/bkry/internal/server/models/entities/question_context"
)

func beforeSave(entity question_context.QuestionContext, editorID *string) question_context.QuestionContext {
	// Return QuestionContext
	return entity
}