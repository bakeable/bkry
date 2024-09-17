package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []product.Product {
	// Get Products
	entities, err := repository.GetAllProducts()
	if err != nil {
		panic(err)
	}

	// Process Product entities
	entities = afterGetAll(entities)

	// Return Products
	return entities
}