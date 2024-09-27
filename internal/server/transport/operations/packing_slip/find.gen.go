package packing_slip_operations

import (
	packing_slip "github.com/bakeable/bkry/internal/server/models/entities/packing_slip"
	repo "github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) packing_slip.PackingSlip {
	// Find PackingSlip
	entity, err := repository.FindPackingSlip(queries)
	if err != nil {
		panic(err)
	}

	// Process PackingSlip entity
	entity = afterFind(entity)

	// Return PackingSlip
	return entity
}