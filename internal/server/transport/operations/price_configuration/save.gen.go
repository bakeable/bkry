package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/internal/server/models/entities/price_configuration"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity price_configuration.PriceConfiguration, editorID *string) *string {
	// Preprocess PriceConfiguration
	entity = beforeSave(entity, editorID)

	// Save PriceConfiguration
	id, err := repository.SavePriceConfiguration(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving PriceConfiguration entity
	afterSave(entity, editorID)

	// Return PriceConfiguration
	return &id
}