package airline_pricing

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to AirlinePricing struct
func Decode(m interface{}) AirlinePricing {
	if m, ok := m.(map[string]interface{}); ok {
		return AirlinePricing{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Colors: utils.DecodeIntArray(m["colors"], []int{}),
			Description: utils.DecodeString(m["description"], ""),
			Dividers: utils.DecodeInt(m["dividers"], 0),
			Kind: utils.DecodeString(m["kind"], "AirlinePricing"),
			Margin: DecodeMargin(m["margin"]),
			Units: utils.DecodeIntArray(m["units"], []int{}),
			Values: DecodeAirlinePricingValueArray(m["values"]),
			Width: utils.DecodeFloat64(m["width"], 0.0),
			WidthType: utils.DecodeString(m["widthType"], "MILLIMETERS"),
		}
	}

	return AirlinePricing{}
}

// DecodeAll converts a slice of maps to a slice of AirlinePricing struct
func DecodeAll(ms interface{}) []AirlinePricing {
	var entities []AirlinePricing

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

// DecodeMargin converts a map to Margin struct
func DecodeMargin(m interface{}) Margin {
	if m == nil {
		return Margin{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Margin{
			Amount: utils.DecodeInt(val["amount"], 0),
			Percentage: utils.DecodeInt(val["percentage"], 0),
			PercentageType: utils.DecodeString(val["percentageType"], "profit-margin"),
			Staggered: utils.DecodeBool(val["staggered"], false),
			Type: utils.DecodeString(val["type"], "percentage"),
			Values: utils.DecodeInterfaceArray(val["values"], []interface{}{}),
		}
	}
	return Margin{}
}


// DecodeAirlinePricingValueArray converts a map to AirlinePricingValue struct
func DecodeAirlinePricingValueArray(m interface{}) []AirlinePricingValue {
	if m == nil {
		return []AirlinePricingValue{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []AirlinePricingValue
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, AirlinePricingValue{
				PricingAttributes: DecodePricingAttributes(v["attributes"]),
				PricingValue: DecodePricingValue(v["value"]),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []AirlinePricingValue
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, AirlinePricingValue{
					PricingAttributes: DecodePricingAttributes(val["attributes"]),
					PricingValue: DecodePricingValue(val["value"]),
				})
			}
		}
		return entities
	}
	return []AirlinePricingValue{}
}

// DecodePricingAttributes converts a map to PricingAttributes struct
func DecodePricingAttributes(m interface{}) PricingAttributes {
	if m == nil {
		return PricingAttributes{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return PricingAttributes{
			Colors: utils.DecodeInt(val["colors"], 0),
			Units: utils.DecodeInt(val["units"], 0),
		}
	}
	return PricingAttributes{}
}


// DecodePricingValue converts a map to PricingValue struct
func DecodePricingValue(m interface{}) PricingValue {
	if m == nil {
		return PricingValue{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return PricingValue{
			DecimalValue: utils.DecodeFloat64(val["decimalValue"], 0),
			FloatValue: utils.DecodeFloat64(val["floatValue"], 0),
			IntValue: utils.DecodeInt(val["intValue"], 0),
			TextValue: utils.DecodeString(val["textValue"], ""),
		}
	}
	return PricingValue{}
}



