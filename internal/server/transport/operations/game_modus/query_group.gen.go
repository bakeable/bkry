package game_modus_operations

import (
	game_modus "github.com/bakeable/bkry/internal/server/models/entities/game_modus"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []database.Query) []game_modus.GameModus {
	// Query GameModus group
	entities, err := repository.QueryGameModusGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process GameModus entities
	entities = afterQueryGroup(entities)

	// Return GameModus
	return entities
}