package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, queries []database.Query, pagination database.Pagination) []tag.Tag {
	// Get Tags
	entities, err := repository.QueryTags(queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process Tag entities
	entities = afterQuery(entities)

	// Return Tags
	return entities
}