package attribute_option
import "github.com/bakeable/bkry/data/general/auditinfo"

type AttributeOption struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The option's key
	Key string `json:"key"`
	// The maximum value for the option
	Max float64 `json:"max"`
	// The minimum value for the option
	Min float64 `json:"min"`
	// The option's value
	Value string `json:"value"`
}








