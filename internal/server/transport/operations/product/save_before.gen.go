package product_operations

import (
	product "github.com/bakeable/bkry/data/entities/product"
)

func beforeSave(entity product.Product, editorID *string) product.Product {
	// Return Product
	return entity
}