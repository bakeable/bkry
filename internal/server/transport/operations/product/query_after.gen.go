package product_operations

import (
	product "github.com/bakeable/bkry/internal/server/models/entities/product"
)


func afterQuery(entities []product.Product) []product.Product {
	return entities
}