package question_bundle_localisation_operations

import (
	question_bundle_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []database.Query) []question_bundle_localisation.QuestionBundleLocalisation {
	// Query QuestionBundleLocalisations group
	entities, err := repository.QueryQuestionBundleLocalisationsGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process QuestionBundleLocalisation entities
	entities = afterQueryGroup(entities)

	// Return QuestionBundleLocalisations
	return entities
}