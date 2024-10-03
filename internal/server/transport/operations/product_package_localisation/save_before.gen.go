package product_package_localisation_operations

import (
	product_package_localisation "github.com/bakeable/bkry/internal/server/models/entities/product_package_localisation"
)

func beforeSave(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) product_package_localisation.ProductPackageLocalisation {
	// Return ProductPackageLocalisation
	return entity
}