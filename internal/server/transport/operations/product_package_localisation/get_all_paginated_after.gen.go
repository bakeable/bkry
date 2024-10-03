package product_package_localisation_operations

import (
	product_package_localisation "github.com/bakeable/bkry/internal/server/models/entities/product_package_localisation"
)


func afterGetAllPaginated(productPackageID string, entities []product_package_localisation.ProductPackageLocalisation, pageSize int, pageNumber int, orderBy string, ascending bool) []product_package_localisation.ProductPackageLocalisation {
	return entities
}