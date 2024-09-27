package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity attribute_option_set.AttributeOptionSet, editorID *string) *string {
	// Preprocess AttributeOptionSet
	entity = beforeSave(entity, editorID)

	// Save AttributeOptionSet
	id, err := repository.SaveAttributeOptionSet(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving AttributeOptionSet entity
	afterSave(entity, editorID)

	// Return AttributeOptionSet
	return &id
}