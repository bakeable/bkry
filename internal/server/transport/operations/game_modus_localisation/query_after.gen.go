package game_modus_localisation_operations

import (
	game_modus_localisation "github.com/bakeable/bkry/internal/server/models/entities/game_modus_localisation"
)


func afterQuery(gameModusID string, entities []game_modus_localisation.GameModusLocalisation) []game_modus_localisation.GameModusLocalisation {
	return entities
}