package question_context_localisation_operations

import (
	question_context_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
)

func afterFind(questionContextID string, entity question_context_localisation.QuestionContextLocalisation) question_context_localisation.QuestionContextLocalisation {
	return entity
}