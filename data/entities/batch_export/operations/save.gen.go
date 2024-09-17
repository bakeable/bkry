package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/data/entities/batch_export"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Save(repository repo.IRepository, entity batch_export.BatchExport, editorID *string) *string {
	// Preprocess BatchExport
	entity = beforeSave(entity, editorID)

	// Save BatchExport
	id, err := repository.SaveBatchExport(entity, editorID)
	if err != nil {
		panic(err)
	}

	// Set ID on entity
	entity.ID = id

	// Take care of business logic after saving BatchExport entity
	afterSave(entity, editorID)

	// Return BatchExport
	return &id
}