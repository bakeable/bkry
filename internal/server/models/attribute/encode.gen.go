package attribute

// Encode converts a Attribute struct to a map
func Encode(e Attribute) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "default": e.Default,
        "defaultOptions": EncodeDefaultOptionArray(e.DefaultOptions),
        "description": e.Description,
        "included": e.Included,
        "incrementalPricing": e.IncrementalPricing,
        "key": e.Key,
        "_kind": e.Kind,
        "maxOptionId": e.MaxOptionID,
        "name": e.Name,
        "nid": e.Nid,
        "optionSets": e.OptionSets,
        "optional": e.Optional,
        "options": e.Options,
        "order": e.Order,
        "prefix": e.Prefix,
        "priceTypes": e.PriceTypes,
        "surcharges": e.Surcharges,
        "type": e.Type,
    }
    return result
}
// EncodeDefaultOptionArray converts a DefaultOption struct to a map
func EncodeDefaultOptionArray(e []DefaultOption) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "id": v.ID,
            "key": v.Key,
            "max": v.Max,
            "min": v.Min,
            "value": v.Value,
        }
    }
    return result
}




