package product_operations

import (
	product "github.com/bakeable/bkry/internal/server/models/entities/product"
)


func afterQueryGroup(entities []product.Product) []product.Product {
	return entities
}