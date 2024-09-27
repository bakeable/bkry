package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/internal/server/models/entities/delivery_entry"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) delivery_entry.DeliveryEntry {
	// Find DeliveryEntry
	entity, err := repository.FindDeliveryEntry(queries)
	if err != nil {
		panic(err)
	}

	// Process DeliveryEntry entity
	entity = afterFind(entity)

	// Return DeliveryEntry
	return entity
}