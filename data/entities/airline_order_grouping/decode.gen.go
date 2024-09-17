package airline_order_grouping

import (
	"github.com/bakeable/bkry/data/entities/airline_order_group"
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to AirlineOrderGrouping struct
func Decode(m interface{}) AirlineOrderGrouping {
	if m, ok := m.(map[string]interface{}); ok {
		return AirlineOrderGrouping{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AfasOrderNumbers: utils.DecodeStringArray(m["afasOrderNumbers"], []string{}),
			FinalPrice: utils.DecodeFloat64(m["finalPrice"], 0),
			FirstOrderDate: utils.DecodeString(m["firstOrderDate"], "time.Now()"),
			Kind: utils.DecodeString(m["_kind"], "AirlineOrderGrouping"),
			MarginAmount: utils.DecodeFloat64(m["marginAmount"], 0),
			MarginPercentage: utils.DecodeFloat64(m["marginPercentage"], 0),
			Name: utils.DecodeString(m["name"], ""),
			OrderDate: utils.DecodeString(m["orderDate"], "time.Now()"),
			OrderGroups: airline_order_group.DecodeAll(m["orderGroups"]),
			OrderIds: utils.DecodeStringArray(m["orderIds"], []string{}),
			OriginalTotalCosts: utils.DecodeFloat64(m["originalTotalCosts"], 0),
			ProfitAmount: utils.DecodeFloat64(m["profitAmount"], 0),
			ProfitPercentage: utils.DecodeFloat64(m["profitPercentage"], 0),
			SavedCosts: utils.DecodeFloat64(m["savedCosts"], 0),
			SentToAfas: utils.DecodeBool(m["sentToAfas"], false),
			SentToAfasTimestamp: utils.DecodeInt64(m["sentToAfasTimestamp"], 0),
			TotalCosts: utils.DecodeFloat64(m["totalCosts"], 0),
			Units: utils.DecodeInt(m["units"], 0),
		}
	}

	return AirlineOrderGrouping{}
}

// DecodeAll converts a slice of maps to a slice of AirlineOrderGrouping struct
func DecodeAll(ms interface{}) []AirlineOrderGrouping {
	var entities []AirlineOrderGrouping

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

