package airline_pricing

// Encode converts a AirlinePricing struct to a map
func Encode(e AirlinePricing) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "colors": e.Colors,
        "description": e.Description,
        "dividers": e.Dividers,
        "kind": e.Kind,
        "margin": EncodeMargin(e.Margin),
        "units": e.Units,
        "values": EncodeAirlinePricingValueArray(e.Values),
        "width": e.Width,
        "widthType": e.WidthType,
    }
    return result
}
// EncodeMargin converts a Margin struct to a map
func EncodeMargin(e Margin) map[string]interface{} {
    return map[string]interface{}{
        "amount": e.Amount,
        "percentage": e.Percentage,
        "percentageType": e.PercentageType,
        "staggered": e.Staggered,
        "type": e.Type,
        "values": e.Values,
    }
}

// EncodeAirlinePricingValueArray converts a AirlinePricingValue struct to a map
func EncodeAirlinePricingValueArray(e []AirlinePricingValue) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "attributes": EncodePricingAttributes(v.PricingAttributes),
            "value": EncodePricingValue(v.PricingValue),
        }
    }
    return result
}
// EncodePricingAttributes converts a PricingAttributes struct to a map
func EncodePricingAttributes(e PricingAttributes) map[string]interface{} {
    return map[string]interface{}{
        "colors": e.Colors,
        "units": e.Units,
    }
}

// EncodePricingValue converts a PricingValue struct to a map
func EncodePricingValue(e PricingValue) map[string]interface{} {
    return map[string]interface{}{
        "decimalValue": e.DecimalValue,
        "floatValue": e.FloatValue,
        "intValue": e.IntValue,
        "textValue": e.TextValue,
    }
}





