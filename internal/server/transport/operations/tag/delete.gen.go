package tag_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, tagID string) {
	// Perform before delete actions for Tag entity
	beforeDelete(tagID)

	// Delete Tag
	err := repository.DeleteTag(tagID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for Tag entity
	afterDelete(tagID)
}