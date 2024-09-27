package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/internal/server/models/entities/airline_settings"
)


func afterGetAll(entities []airline_settings.AirlineSettings) []airline_settings.AirlineSettings {
	return entities
}