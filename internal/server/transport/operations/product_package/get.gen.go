package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
	repo "github.com/bakeable/bkry/internal/server/database/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Get(repository repo.IRepository, productPackageID string) product_package.ProductPackage {
	// Get ProductPackage
	entity, err := repository.GetProductPackage(productPackageID)
	if err != nil {
		panic(err)
	}

	// Process ProductPackage entity
	entity = afterGet(productPackageID, entity)

	// Return ProductPackage
	return entity
}