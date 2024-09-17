package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []datastore.Query) []airline_order.AirlineOrder {
	// Query AirlineOrders group
	entities, err := repository.QueryAirlineOrdersGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process AirlineOrder entities
	entities = afterQueryGroup(entities)

	// Return AirlineOrders
	return entities
}