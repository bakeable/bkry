package packing_slip_operations

import (
	packing_slip "github.com/bakeable/bkry/internal/server/models/entities/packing_slip"
)


func afterGetAllPaginated(entities []packing_slip.PackingSlip, pageSize int, pageNumber int, orderBy string, ascending bool) []packing_slip.PackingSlip {
	return entities
}