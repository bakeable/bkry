package question_context_localisation_operations

import (
	question_context_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
)

func afterGet(questionContextID string, questionContextLocalisationID string, entity question_context_localisation.QuestionContextLocalisation) question_context_localisation.QuestionContextLocalisation {
	return entity
}