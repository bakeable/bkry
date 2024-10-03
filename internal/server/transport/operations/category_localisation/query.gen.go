package category_localisation_operations

import (
	category_localisation "github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, categoryID string, queries []database.Query, pagination database.Pagination) []category_localisation.CategoryLocalisation {
	// Get CategoryLocalisations
	entities, err := repository.QueryCategoryLocalisations(categoryID, queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process CategoryLocalisation entities
	entities = afterQuery(categoryID, entities)

	// Return CategoryLocalisations
	return entities
}