package cloud_function
import "github.com/bakeable/bkry/data/general/auditinfo"

type CloudFunction struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The cloud function's description
	Description string `json:"description"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The cloud function's name
	Name string `json:"name"`
	// The cloud function's region
	Region string `json:"region"`
	// The cloud function's runtime
	Runtime string `json:"runtime"`
	// The cloud function's trigger type
	TriggerType string `json:"triggerType"`
	// The cloud function's trigger value
	TriggerValue string `json:"triggerValue"`
}








