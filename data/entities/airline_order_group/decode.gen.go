package airline_order_group

import (
	"github.com/bakeable/bkry/data/entities/airline_order"
	"github.com/bakeable/bkry/data/entities/airline_pricing"
	"github.com/bakeable/bkry/data/entities/airline_settings"
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
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
			Orders: airline_order.DecodeAll(m["orders"]),
			PricingTables: airline_pricing.DecodeAll(m["pricingTables"]),
			Settings: airline_settings.Decode(m["settings"]),
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

