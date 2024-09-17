package price_layer

// Encode converts a PriceLayer struct to a map
func Encode(e PriceLayer) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "attributes": e.Attributes,
        "description": e.Description,
        "external": e.External,
        "includeInMarginCalculation": e.IncludeInMarginCalculation,
        "includeInPurchasePrice": e.IncludeInPurchasePrice,
        "incrementalPricing": e.IncrementalPricing,
        "_kind": e.Kind,
        "name": e.Name,
        "options": e.Options,
        "perUnits": e.PerUnits,
        "priceType": e.PriceType,
        "roundingUnits": e.RoundingUnits,
        "saveAsNew": e.SaveAsNew,
        "template": e.Template,
        "values": EncodeVariablePriceValuesArray(e.VariablePriceValues),
    }
    return result
}
// EncodeVariablePriceValuesArray converts a VariablePriceValues struct to a map
func EncodeVariablePriceValuesArray(e []VariablePriceValues) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "attributes": v.Attributes,
            "value": EncodeVariablePriceValue(v.VariablePriceValue),
        }
    }
    return result
}
// EncodeVariablePriceValue converts a VariablePriceValue struct to a map
func EncodeVariablePriceValue(e VariablePriceValue) map[string]interface{} {
    return map[string]interface{}{
        "decimalValue": e.DecimalValue,
        "floatValue": e.FloatValue,
        "intValue": e.IntValue,
        "textValue": e.TextValue,
    }
}





