package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, productID string) product.Product {
	// Get Product
	entity, err := repository.GetProduct(productID)
	if err != nil {
		panic(err)
	}

	// Process Product entity
	entity = afterGet(productID, entity)

	// Return Product
	return entity
}