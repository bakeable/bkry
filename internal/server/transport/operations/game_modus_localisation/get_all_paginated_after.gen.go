package game_modus_localisation_operations

import (
	game_modus_localisation "github.com/bakeable/bkry/internal/server/models/entities/game_modus_localisation"
)


func afterGetAllPaginated(gameModusID string, entities []game_modus_localisation.GameModusLocalisation, pageSize int, pageNumber int, orderBy string, ascending bool) []game_modus_localisation.GameModusLocalisation {
	return entities
}