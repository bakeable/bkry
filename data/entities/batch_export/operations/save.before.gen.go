package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/data/entities/batch_export"
)

func beforeSave(entity batch_export.BatchExport, editorID *string) batch_export.BatchExport {
	// Return BatchExport
	return entity
}