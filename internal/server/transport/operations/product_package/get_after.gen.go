package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
)

func afterGet(productPackageID string, entity product_package.ProductPackage) product_package.ProductPackage {
	return entity
}