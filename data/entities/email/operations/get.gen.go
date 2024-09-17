package email_operations

import (
	email "github.com/bakeable/bkry/data/entities/email"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, emailID string) email.Email {
	// Get Email
	entity, err := repository.GetEmail(emailID)
	if err != nil {
		panic(err)
	}

	// Process Email entity
	entity = afterGet(emailID, entity)

	// Return Email
	return entity
}