package question_context_operations

import (
	question_context "github.com/bakeable/bkry/internal/server/models/entities/question_context"
)

func afterSave(entity question_context.QuestionContext, editorID *string) {}