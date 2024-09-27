package user_profile_operations

import (
	user_profile "github.com/bakeable/bkry/internal/server/models/entities/user_profile"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []datastore.Query) []user_profile.UserProfile {
	// Query UserProfiles group
	entities, err := repository.QueryUserProfilesGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process UserProfile entities
	entities = afterQueryGroup(entities)

	// Return UserProfiles
	return entities
}