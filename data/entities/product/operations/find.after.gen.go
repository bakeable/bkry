package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
)

func afterFind(entity product.Product) product.Product {
	return entity
}