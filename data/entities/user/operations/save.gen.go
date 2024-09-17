package user_operations

import (
	user "github.com/bakeable/bkry/data/entities/user"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity user.User, editorID *string) *string {
	// Preprocess User
	entity = beforeSave(entity, editorID)

	// Save User
	id, err := repository.SaveUser(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving User entity
	afterSave(entity, editorID)

	// Return User
	return &id
}