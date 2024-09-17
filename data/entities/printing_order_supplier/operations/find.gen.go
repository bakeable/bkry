package printing_order_supplier_operations

import (
	printing_order_supplier "github.com/bakeable/bkry/data/entities/printing_order_supplier"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Find(repository repo.IRepository, queries []datastore.Query) printing_order_supplier.PrintingOrderSupplier {
	// Find PrintingOrderSupplier
	entity, err := repository.FindPrintingOrderSupplier(queries)
	if err != nil {
		panic(err)
	}

	// Process PrintingOrderSupplier entity
	entity = afterFind(entity)

	// Return PrintingOrderSupplier
	return entity
}