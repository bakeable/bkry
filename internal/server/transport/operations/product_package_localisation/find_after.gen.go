package product_package_localisation_operations

import (
	product_package_localisation "github.com/bakeable/bkry/internal/server/models/entities/product_package_localisation"
)

func afterFind(productPackageID string, entity product_package_localisation.ProductPackageLocalisation) product_package_localisation.ProductPackageLocalisation {
	return entity
}