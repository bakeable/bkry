package email
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Email struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The entity kind
	Kind string `json:"_kind"`
}








