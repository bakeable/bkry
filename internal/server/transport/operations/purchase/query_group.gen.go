package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
	"github.com/bakeable/bkry/internal/server/database"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func QueryGroup(repository repo.IRepository, queries []database.Query) []purchase.Purchase {
	// Query Purchases group
	entities, err := repository.QueryPurchasesGroup(queries)
	if err != nil {
		panic(err)
	}

	// Process Purchase entities
	entities = afterQueryGroup(entities)

	// Return Purchases
	return entities
}