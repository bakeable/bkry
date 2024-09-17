package quotation_operations

import (
	quotation "github.com/bakeable/bkry/data/entities/quotation"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []datastore.Query) []quotation.Quotation {
	// Query Quotations group
	entities, err := repository.QueryQuotationsGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process Quotation entities
	entities = afterQueryGroup(entities)

	// Return Quotations
	return entities
}