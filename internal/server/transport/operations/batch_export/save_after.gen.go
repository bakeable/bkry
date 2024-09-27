package batch_export_operations

import (
	batch_export "github.com/bakeable/bkry/internal/server/models/entities/batch_export"
)

func afterSave(entity batch_export.BatchExport, editorID *string) {}