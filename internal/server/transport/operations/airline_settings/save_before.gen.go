package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/internal/server/models/entities/airline_settings"
)

func beforeSave(entity airline_settings.AirlineSettings, editorID *string) airline_settings.AirlineSettings {
	// Return AirlineSettings
	return entity
}