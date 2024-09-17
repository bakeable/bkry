package packing_slip_package_operations

import (
	packing_slip_package "github.com/bakeable/bkry/data/entities/packing_slip_package"
)

func afterGet(packingSlipPackageID string, entity packing_slip_package.PackingSlipPackage) packing_slip_package.PackingSlipPackage {
	return entity
}