package price_configuration

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to PriceConfiguration struct
func Decode(m interface{}) PriceConfiguration {
	if m, ok := m.(map[string]interface{}); ok {
		return PriceConfiguration{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AdditionalCosts: utils.DecodeFloat64(m["additional_costs"], 0),
			Attributes: utils.DecodeStringMap(m["attributes"], map[string]string{}),
			Currency: utils.DecodeString(m["currency"], "EUR"),
			FinalPrice: utils.DecodeFloat64(m["final_price"], 0),
			FixedCosts: utils.DecodeFloat64(m["fixed_costs"], 0),
			Kind: utils.DecodeString(m["_kind"], "PriceConfiguration"),
			MarginAmount: utils.DecodeFloat64(m["margin_amount"], 0),
			MarginPercentage: utils.DecodeFloat64(m["margin_percentage"], 0),
			ProductId: utils.DecodeString(m["product_id"], ""),
			Sku: utils.DecodeString(m["sku"], ""),
			TotalCosts: utils.DecodeFloat64(m["total_costs"], 0),
			UnitCosts: utils.DecodeFloat64(m["unit_costs"], 0),
			UnitMargin: utils.DecodeFloat64(m["unit_margin"], 0),
			Units: utils.DecodeInt(m["units"], 0),
			VariableCosts: utils.DecodeFloat64(m["variable_costs"], 0),
		}
	}

	return PriceConfiguration{}
}

// DecodeAll converts a slice of maps to a slice of PriceConfiguration struct
func DecodeAll(ms interface{}) []PriceConfiguration {
	var entities []PriceConfiguration

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

