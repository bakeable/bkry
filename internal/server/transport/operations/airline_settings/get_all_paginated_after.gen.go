package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/internal/server/models/entities/airline_settings"
)


func afterGetAllPaginated(entities []airline_settings.AirlineSettings, pageSize int, pageNumber int, orderBy string, ascending bool) []airline_settings.AirlineSettings {
	return entities
}