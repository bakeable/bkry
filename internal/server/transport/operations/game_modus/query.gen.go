package game_modus_operations

import (
	game_modus "github.com/bakeable/bkry/internal/server/models/entities/game_modus"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, queries []database.Query, pagination database.Pagination) []game_modus.GameModus {
	// Get GameModus
	entities, err := repository.QueryGameModus(queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process GameModus entities
	entities = afterQuery(entities)

	// Return GameModus
	return entities
}