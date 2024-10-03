package question_bundle_operations

import (
	question_bundle "github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
)

func afterSave(entity question_bundle.QuestionBundle, editorID *string) {}