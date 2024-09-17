package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/data/entities/price_configuration"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) price_configuration.PriceConfiguration {
	// Find PriceConfiguration
	entity, err := repository.FindPriceConfiguration(queries)
	if err != nil {
		panic(err)
	}

	// Process PriceConfiguration entity
	entity = afterFind(entity)

	// Return PriceConfiguration
	return entity
}