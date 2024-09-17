package email_operations

import (
	email "github.com/bakeable/bkry/data/entities/email"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) email.Email {
	// Find Email
	entity, err := repository.FindEmail(queries)
	if err != nil {
		panic(err)
	}

	// Process Email entity
	entity = afterFind(entity)

	// Return Email
	return entity
}