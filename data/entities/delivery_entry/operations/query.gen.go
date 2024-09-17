package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/data/entities/delivery_entry"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, queries []datastore.Query, pagination datastore.Pagination) []delivery_entry.DeliveryEntry {
	// Get DeliveryEntries
	entities, err := repository.QueryDeliveryEntries(queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process DeliveryEntry entities
	entities = afterQuery(entities)

	// Return DeliveryEntries
	return entities
}