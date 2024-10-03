package question_context_operations

import (
	question_context "github.com/bakeable/bkry/internal/server/models/entities/question_context"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, questionContextID string) question_context.QuestionContext {
	// Get QuestionContext
	entity, err := repository.GetQuestionContext(questionContextID)
	if err != nil {
		panic(err)
	}

	// Process QuestionContext entity
	entity = afterGet(questionContextID, entity)

	// Return QuestionContext
	return entity
}