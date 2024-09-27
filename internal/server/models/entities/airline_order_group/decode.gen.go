package airline_order_group

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to AirlineOrderGroup struct
func Decode(m interface{}) AirlineOrderGroup {
	if m, ok := m.(map[string]interface{}); ok {
		return AirlineOrderGroup{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AfasOrderNumbers: utils.DecodeStringArray(m["afasOrderNumbers"], []string{}),
			Kind: utils.DecodeString(m["_kind"], "AirlineOrderGroup"),
			Name: utils.DecodeString(m["name"], ""),
			Orders: All(m["orders"]),
			PricingTables: All(m["pricingTables"]),
			Settings: (m["settings"]),
		}
	}

	return AirlineOrderGroup{}
}

// DecodeAll converts a slice of maps to a slice of AirlineOrderGroup struct
func DecodeAll(ms interface{}) []AirlineOrderGroup {
	var entities []AirlineOrderGroup

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

