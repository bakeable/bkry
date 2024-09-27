package packing_slip_operations

import (
	packing_slip "github.com/bakeable/bkry/internal/server/models/entities/packing_slip"
)

func beforeSave(entity packing_slip.PackingSlip, editorID *string) packing_slip.PackingSlip {
	// Return PackingSlip
	return entity
}