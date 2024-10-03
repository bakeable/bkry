package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
)


func afterGetAllPaginated(entities []product_package.ProductPackage, pageSize int, pageNumber int, orderBy string, ascending bool) []product_package.ProductPackage {
	return entities
}