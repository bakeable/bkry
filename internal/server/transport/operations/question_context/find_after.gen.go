package question_context_operations

import (
	question_context "github.com/bakeable/bkry/internal/server/models/entities/question_context"
)

func afterFind(entity question_context.QuestionContext) question_context.QuestionContext {
	return entity
}