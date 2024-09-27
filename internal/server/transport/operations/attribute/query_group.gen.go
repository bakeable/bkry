package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []datastore.Query) []attribute.Attribute {
	// Query Attributes group
	entities, err := repository.QueryAttributesGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process Attribute entities
	entities = afterQueryGroup(entities)

	// Return Attributes
	return entities
}