package game_modus_operations

import (
	game_modus "github.com/bakeable/bkry/internal/server/models/entities/game_modus"
)


func afterGetAllPaginated(entities []game_modus.GameModus, pageSize int, pageNumber int, orderBy string, ascending bool) []game_modus.GameModus {
	return entities
}