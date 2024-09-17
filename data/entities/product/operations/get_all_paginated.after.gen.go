package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
)


func afterGetAllPaginated(entities []product.Product, pageSize int, pageNumber int, orderBy string, ascending bool) []product.Product {
	return entities
}