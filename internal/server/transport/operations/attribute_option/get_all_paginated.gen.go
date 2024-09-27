package attribute_option_operations

import (
	attribute_option "github.com/bakeable/bkry/internal/server/models/entities/attribute_option"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]attribute_option.AttributeOption, datastore.Pagination) {
	// Get AttributeOptions
	entities, pagination, err := repository.GetAllAttributeOptionsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process AttributeOption entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return AttributeOptions
	return entities, pagination
}