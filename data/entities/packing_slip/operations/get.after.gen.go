package packing_slip_operations

import (
	packing_slip "github.com/bakeable/bkry/data/entities/packing_slip"
)

func afterGet(packingSlipID string, entity packing_slip.PackingSlip) packing_slip.PackingSlip {
	return entity
}