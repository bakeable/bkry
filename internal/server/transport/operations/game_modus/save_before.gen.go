package game_modus_operations

import (
	game_modus "github.com/bakeable/bkry/internal/server/models/entities/game_modus"
)

func beforeSave(entity game_modus.GameModus, editorID *string) game_modus.GameModus {
	// Return GameModus
	return entity
}