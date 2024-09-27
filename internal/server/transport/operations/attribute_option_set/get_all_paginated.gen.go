package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, datastore.Pagination) {
	// Get AttributeOptionSets
	entities, pagination, err := repository.GetAllAttributeOptionSetsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process AttributeOptionSet entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return AttributeOptionSets
	return entities, pagination
}