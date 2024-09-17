package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/data/entities/airline_order_group"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity airline_order_group.AirlineOrderGroup, editorID *string) *string {
	// Preprocess AirlineOrderGroup
	entity = beforeSave(entity, editorID)

	// Save AirlineOrderGroup
	id, err := repository.SaveAirlineOrderGroup(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving AirlineOrderGroup entity
	afterSave(entity, editorID)

	// Return AirlineOrderGroup
	return &id
}