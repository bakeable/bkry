package question_bundle_localisation_operations

import (
	question_bundle_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, questionBundleID string, queries []database.Query, pagination database.Pagination) []question_bundle_localisation.QuestionBundleLocalisation {
	// Get QuestionBundleLocalisations
	entities, err := repository.QueryQuestionBundleLocalisations(questionBundleID, queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process QuestionBundleLocalisation entities
	entities = afterQuery(questionBundleID, entities)

	// Return QuestionBundleLocalisations
	return entities
}