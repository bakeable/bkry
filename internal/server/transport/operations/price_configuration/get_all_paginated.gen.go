package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/internal/server/models/entities/price_configuration"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, datastore.Pagination) {
	// Get PriceConfigurations
	entities, pagination, err := repository.GetAllPriceConfigurationsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process PriceConfiguration entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return PriceConfigurations
	return entities, pagination
}