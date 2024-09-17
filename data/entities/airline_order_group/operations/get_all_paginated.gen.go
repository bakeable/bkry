package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/data/entities/airline_order_group"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, datastore.Pagination) {
	// Get AirlineOrderGroups
	entities, pagination, err := repository.GetAllAirlineOrderGroupsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process AirlineOrderGroup entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return AirlineOrderGroups
	return entities, pagination
}