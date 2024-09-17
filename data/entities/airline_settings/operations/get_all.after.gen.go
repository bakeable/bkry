package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/data/entities/airline_settings"
)


func afterGetAll(entities []airline_settings.AirlineSettings) []airline_settings.AirlineSettings {
	return entities
}