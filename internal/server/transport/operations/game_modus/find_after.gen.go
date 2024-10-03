package game_modus_operations

import (
	game_modus "github.com/bakeable/bkry/internal/server/models/entities/game_modus"
)

func afterFind(entity game_modus.GameModus) game_modus.GameModus {
	return entity
}