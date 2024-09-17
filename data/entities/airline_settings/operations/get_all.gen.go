package airline_settings_operations

import (
	airline_settings "github.com/bakeable/bkry/data/entities/airline_settings"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []airline_settings.AirlineSettings {
	// Get AirlineSettings
	entities, err := repository.GetAllAirlineSettings()
	if err != nil {
		panic(err)
	}

	// Process AirlineSettings entities
	entities = afterGetAll(entities)

	// Return AirlineSettings
	return entities
}