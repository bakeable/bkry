package price_layer

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to PriceLayer struct
func Decode(m interface{}) PriceLayer {
	if m, ok := m.(map[string]interface{}); ok {
		return PriceLayer{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Attributes: utils.DecodeStringArray(m["attributes"], []string{}),
			Description: utils.DecodeString(m["description"], ""),
			External: utils.DecodeBool(m["external"], false),
			IncludeInMarginCalculation: utils.DecodeBool(m["includeInMarginCalculation"], false),
			IncludeInPurchasePrice: utils.DecodeBool(m["includeInPurchasePrice"], false),
			IncrementalPricing: AttributeIncrementalPricing(utils.DecodeString(m["incrementalPricing"], "standard")),
			Kind: utils.DecodeString(m["_kind"], "PriceLayer"),
			Name: utils.DecodeString(m["name"], ""),
			Options: utils.DecodeStringArray(m["options"], []string{}),
			PerUnits: utils.DecodeInt(m["perUnits"], 1),
			PriceType: PriceType(utils.DecodeString(m["priceType"], "variable")),
			RoundingUnits: RoundingUnits(utils.DecodeString(m["roundingUnits"], "ratio")),
			SaveAsNew: utils.DecodeBool(m["saveAsNew"], false),
			Template: utils.DecodeString(m["template"], ""),
			VariablePriceValues: DecodeVariablePriceValuesArray(m["values"]),
		}
	}

	return PriceLayer{}
}

// DecodeAll converts a slice of maps to a slice of PriceLayer struct
func DecodeAll(ms interface{}) []PriceLayer {
	var entities []PriceLayer

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

// DecodeVariablePriceValuesArray converts a map to VariablePriceValues struct
func DecodeVariablePriceValuesArray(m interface{}) []VariablePriceValues {
	if m == nil {
		return []VariablePriceValues{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []VariablePriceValues
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, VariablePriceValues{
				Attributes: utils.DecodeMap(v["attributes"], nil),
				VariablePriceValue: DecodeVariablePriceValue(v["value"]),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []VariablePriceValues
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, VariablePriceValues{
					Attributes: utils.DecodeMap(val["attributes"], nil),
					VariablePriceValue: DecodeVariablePriceValue(val["value"]),
				})
			}
		}
		return entities
	}
	return []VariablePriceValues{}
}

// DecodeVariablePriceValue converts a map to VariablePriceValue struct
func DecodeVariablePriceValue(m interface{}) VariablePriceValue {
	if m == nil {
		return VariablePriceValue{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return VariablePriceValue{
			DecimalValue: utils.DecodeFloat64(val["decimalValue"], 0.0),
			FloatValue: utils.DecodeFloat64(val["floatValue"], 0.0),
			IntValue: utils.DecodeInt(val["intValue"], 0),
			TextValue: utils.DecodeString(val["textValue"], ""),
		}
	}
	return VariablePriceValue{}
}



