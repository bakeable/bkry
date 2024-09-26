package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
)

func afterGet(productID string, entity product.Product) product.Product {
	return entity
}