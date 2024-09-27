package batch_export
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type BatchExport struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The timestamp from which the export should start
	ExportFromTimestamp int64 `json:"exportFromTimestamp"`
	// The IDs of the objects to export
	Ids []string `json:"ids"`
	// The kind of the object
	Kind string `json:"_kind"`
}








