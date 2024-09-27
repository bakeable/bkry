package airline_pricing_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, airlinePricingID string) {
	// Perform before delete actions for AirlinePricing entity
	beforeDelete(airlinePricingID)

	// Delete AirlinePricing
	err := repository.DeleteAirlinePricing(airlinePricingID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for AirlinePricing entity
	afterDelete(airlinePricingID)
}