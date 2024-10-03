package author
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Author struct {
	// The unique identifier for the Author
	ID string `json:"id"`
	// The audit information for the modification of the Author
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the Author
	Created auditinfo.AuditInfo `json:"created"`
	// The display name of the author
	DisplayName string `json:"display_name"`
}






