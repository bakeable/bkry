package question_context_localisation_operations

import (
	question_context_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
)

func afterSave(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) {}