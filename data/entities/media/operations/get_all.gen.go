package media_operations

import (
	media "github.com/bakeable/bkry/data/entities/media"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []media.Media {
	// Get Media
	entities, err := repository.GetAllMedia()
	if err != nil {
		panic(err)
	}

	// Process Media entities
	entities = afterGetAll(entities)

	// Return Media
	return entities
}