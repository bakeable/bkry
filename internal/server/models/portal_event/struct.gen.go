package portal_event
import "github.com/bakeable/bkry/data/general/auditinfo"

type PortalEvent struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The entity ID
	EntityID string `json:"entityId"`
	// The entity name
	EntityName string `json:"entityName"`
	// The entity kind
	Kind string `json:"_kind"`
	// The metadata of the event
	Metadata map[string]interface{} `json:"metadata"`
	// The timestamp of the event
	Timestamp int `json:"timestamp"`
	// The timezone of the event
	Timezone int `json:"timezone"`
	// The type of the event
	Type PortalEventType `json:"type"`
	// The user triggering the event
	UserId string `json:"userId"`
}




type PortalEventType string




