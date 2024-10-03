package question_bundle_localisation_operations

import (
	question_bundle_localisation "github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, questionBundleID string, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, database.Pagination) {
	// Get QuestionBundleLocalisations
	entities, pagination, err := repository.GetAllQuestionBundleLocalisationsPaginated(questionBundleID, pagination)
	if err != nil {
		panic(err)
	}

	// Process QuestionBundleLocalisation entities
	entities = afterGetAllPaginated(questionBundleID, entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return QuestionBundleLocalisations
	return entities, pagination
}