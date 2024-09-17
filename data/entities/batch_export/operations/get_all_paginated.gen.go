package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/data/entities/batch_export"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]batch_export.BatchExport, datastore.Pagination) {
	// Get BatchExports
	entities, pagination, err := repository.GetAllBatchExportsPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process BatchExport entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return BatchExports
	return entities, pagination
}