package question_context_operations

import (
	question_context "github.com/bakeable/bkry/internal/server/models/entities/question_context"
)


func afterGetAllPaginated(entities []question_context.QuestionContext, pageSize int, pageNumber int, orderBy string, ascending bool) []question_context.QuestionContext {
	return entities
}