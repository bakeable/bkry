package airline_settings_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, airlineSettingsID string) {
	// Perform before delete actions for AirlineSettings entity
	beforeDelete(airlineSettingsID)

	// Delete AirlineSettings
	err := repository.DeleteAirlineSettings(airlineSettingsID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for AirlineSettings entity
	afterDelete(airlineSettingsID)
}