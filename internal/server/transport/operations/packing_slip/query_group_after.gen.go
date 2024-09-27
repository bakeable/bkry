package packing_slip_operations

import (
	packing_slip "github.com/bakeable/bkry/internal/server/models/entities/packing_slip"
)


func afterQueryGroup(entities []packing_slip.PackingSlip) []packing_slip.PackingSlip {
	return entities
}