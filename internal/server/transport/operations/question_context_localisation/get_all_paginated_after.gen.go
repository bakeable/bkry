package question_context_localisation_operations

import (
	question_context_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
)


func afterGetAllPaginated(questionContextID string, entities []question_context_localisation.QuestionContextLocalisation, pageSize int, pageNumber int, orderBy string, ascending bool) []question_context_localisation.QuestionContextLocalisation {
	return entities
}