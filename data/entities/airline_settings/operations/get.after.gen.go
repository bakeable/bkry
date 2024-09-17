package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/data/entities/airline_settings"
)

func afterGet(airlineSettingsID string, entity airline_settings.AirlineSettings) airline_settings.AirlineSettings {
	return entity
}