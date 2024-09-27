package user_operations

import (
	user "github.com/bakeable/bkry/internal/server/models/entities/user"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, userID string) user.User {
	// Get User
	entity, err := repository.GetUser(userID)
	if err != nil {
		panic(err)
	}

	// Process User entity
	entity = afterGet(userID, entity)

	// Return User
	return entity
}