package category_localisation_operations

import (
	category_localisation "github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []database.Query) []category_localisation.CategoryLocalisation {
	// Query CategoryLocalisations group
	entities, err := repository.QueryCategoryLocalisationsGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process CategoryLocalisation entities
	entities = afterQueryGroup(entities)

	// Return CategoryLocalisations
	return entities
}