package portal_event

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to PortalEvent struct
func Decode(m interface{}) PortalEvent {
	if m, ok := m.(map[string]interface{}); ok {
		return PortalEvent{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			EntityID: utils.DecodeString(m["entityId"], ""),
			EntityName: utils.DecodeString(m["entityName"], ""),
			Kind: utils.DecodeString(m["_kind"], "PortalEvent"),
			Metadata: utils.DecodeMap(m["metadata"], map[string]interface{}{}),
			Timestamp: utils.DecodeInt(m["timestamp"], 0),
			Timezone: utils.DecodeInt(m["timezone"], 0),
			Type: PortalEventType(utils.DecodeString(m["type"], "Update")),
			UserId: utils.DecodeString(m["userId"], ""),
		}
	}

	return PortalEvent{}
}

// DecodeAll converts a slice of maps to a slice of PortalEvent struct
func DecodeAll(ms interface{}) []PortalEvent {
	var entities []PortalEvent

	if arr, ok := ms.([]map[string]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	if arr, ok := ms.([]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	return entities
}

