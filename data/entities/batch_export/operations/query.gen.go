package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/data/entities/batch_export"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Query(repository repo.IRepository, queries []datastore.Query, pagination datastore.Pagination) []batch_export.BatchExport {
	// Get BatchExports
	entities, err := repository.QueryBatchExports(queries, pagination)
	if err != nil {
		panic(err)
	}

	// Process BatchExport entities
	entities = afterQuery(entities)

	// Return BatchExports
	return entities
}