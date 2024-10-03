package product_package_operations

import (
	product_package "github.com/bakeable/bkry/internal/server/models/entities/product_package"
)

func beforeSave(entity product_package.ProductPackage, editorID *string) product_package.ProductPackage {
	// Return ProductPackage
	return entity
}