package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/internal/server/models/entities/airline_order_group"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, airlineOrderGroupID string) airline_order_group.AirlineOrderGroup {
	// Get AirlineOrderGroup
	entity, err := repository.GetAirlineOrderGroup(airlineOrderGroupID)
	if err != nil {
		panic(err)
	}

	// Process AirlineOrderGroup entity
	entity = afterGet(airlineOrderGroupID, entity)

	// Return AirlineOrderGroup
	return entity
}