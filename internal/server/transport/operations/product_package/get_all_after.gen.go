package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
)


func afterGetAll(entities []product_package.ProductPackage) []product_package.ProductPackage {
	return entities
}