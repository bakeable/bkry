package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/internal/server/models/entities/batch_export"
)


func afterGetAllPaginated(entities []batch_export.BatchExport, pageSize int, pageNumber int, orderBy string, ascending bool) []batch_export.BatchExport {
	return entities
}