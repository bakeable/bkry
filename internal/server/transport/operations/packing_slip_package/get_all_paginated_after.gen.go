package packing_slip_package_operations

import (
	packing_slip_package "github.com/bakeable/bkry/internal/server/models/entities/packing_slip_package"
)


func afterGetAllPaginated(entities []packing_slip_package.PackingSlipPackage, pageSize int, pageNumber int, orderBy string, ascending bool) []packing_slip_package.PackingSlipPackage {
	return entities
}