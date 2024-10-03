package question_bundle_operations

import (
	question_bundle "github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, queries []database.Query, pagination database.Pagination) []question_bundle.QuestionBundle {
	// Get QuestionBundles
	entities, err := repository.QueryQuestionBundles(queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process QuestionBundle entities
	entities = afterQuery(entities)

	// Return QuestionBundles
	return entities
}