package attribute_option_set
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type AttributeOptionSet struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The optionset's label
	Label string `json:"label"`
	// The optionset's options
	Options []attribute_option.AttributeOption `json:"options"`
}








