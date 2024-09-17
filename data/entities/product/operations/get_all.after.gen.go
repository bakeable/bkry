package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
)


func afterGetAll(entities []product.Product) []product.Product {
	return entities
}