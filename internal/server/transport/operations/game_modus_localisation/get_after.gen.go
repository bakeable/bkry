package game_modus_localisation_operations

import (
	game_modus_localisation "github.com/bakeable/bkry/internal/server/models/entities/game_modus_localisation"
)

func afterGet(gameModusID string, gameModusLocalisationID string, entity game_modus_localisation.GameModusLocalisation) game_modus_localisation.GameModusLocalisation {
	return entity
}