package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/data/entities/batch_export"
)

func afterFind(entity batch_export.BatchExport) batch_export.BatchExport {
	return entity
}