package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []attribute_option_set.AttributeOptionSet {
	// Get AttributeOptionSets
	entities, err := repository.GetAllAttributeOptionSets()
	if err != nil {
		panic(err)
	}

	// Process AttributeOptionSet entities
	entities = afterGetAll(entities)

	// Return AttributeOptionSets
	return entities
}