package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []product_package.ProductPackage {
	// Get ProductPackages
	entities, err := repository.GetAllProductPackages()
	if err != nil {
		panic(err)
	}

	// Process ProductPackage entities
	entities = afterGetAll(entities)

	// Return ProductPackages
	return entities
}