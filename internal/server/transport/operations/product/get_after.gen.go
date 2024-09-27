package product_operations

import (
	product "github.com/bakeable/bkry/internal/server/models/entities/product"
)

func afterGet(productID string, entity product.Product) product.Product {
	return entity
}