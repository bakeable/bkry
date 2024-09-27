package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, attributeOptionSetID string) attribute_option_set.AttributeOptionSet {
	// Get AttributeOptionSet
	entity, err := repository.GetAttributeOptionSet(attributeOptionSetID)
	if err != nil {
		panic(err)
	}

	// Process AttributeOptionSet entity
	entity = afterGet(attributeOptionSetID, entity)

	// Return AttributeOptionSet
	return entity
}