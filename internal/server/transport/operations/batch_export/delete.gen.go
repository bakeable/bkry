package batch_export_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, batchExportID string) {
	// Perform before delete actions for BatchExport entity
	beforeDelete(batchExportID)

	// Delete BatchExport
	err := repository.DeleteBatchExport(batchExportID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for BatchExport entity
	afterDelete(batchExportID)
}