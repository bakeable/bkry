package airline_order_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, airlineOrderID string) {
	// Perform before delete actions for AirlineOrder entity
	beforeDelete(airlineOrderID)

	// Delete AirlineOrder
	err := repository.DeleteAirlineOrder(airlineOrderID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for AirlineOrder entity
	afterDelete(airlineOrderID)
}