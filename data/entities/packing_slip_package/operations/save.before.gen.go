package packing_slip_package_operations

import (
	packing_slip_package "github.com/bakeable/bkry/data/entities/packing_slip_package"
)

func beforeSave(entity packing_slip_package.PackingSlipPackage, editorID *string) packing_slip_package.PackingSlipPackage {
	// Return PackingSlipPackage
	return entity
}