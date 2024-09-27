package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/internal/server/models/entities/batch_export"
)


func afterQueryGroup(entities []batch_export.BatchExport) []batch_export.BatchExport {
	return entities
}