package airline_order_grouping_operations

import (
	airline_order_grouping "github.com/bakeable/bkry/data/entities/airline_order_grouping"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) airline_order_grouping.AirlineOrderGrouping {
	// Find AirlineOrderGrouping
	entity, err := repository.FindAirlineOrderGrouping(queries)
	if err != nil {
		panic(err)
	}

	// Process AirlineOrderGrouping entity
	entity = afterFind(entity)

	// Return AirlineOrderGrouping
	return entity
}