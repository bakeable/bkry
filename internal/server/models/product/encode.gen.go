package product

// Encode converts a Product struct to a map
func Encode(e Product) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "attributeMap": e.AttributeMap,
        "attributes": e.Attributes,
        "categories": e.Categories,
        "configurablePrice": e.ConfigurablePrice,
        "examination": EncodeExamination(e.Examination),
        "examine": e.Examine,
        "family": e.Family,
        "_kind": e.Kind,
        "margin": e.Margin,
        "margins": EncodeMargins(e.Margins),
        "name": e.Name,
        "options": e.Options,
        "priceConfigurationsGenerated": e.PriceConfigurationGenerated,
        "priceConfigurationGeneratedTimestamp": e.PriceConfigurationGeneratedTimestamp,
        "priceLayerIds": e.PriceLayerIDs,
        "priceLayers": e.PriceLayers,
        "properties": e.Properties,
        "sku": e.Sku,
        "status": e.Status,
        "statusIndex": e.StatusIndex,
    }
    return result
}
// EncodeExamination converts a Examination struct to a map
func EncodeExamination(e Examination) map[string]interface{} {
    return map[string]interface{}{
        "priority": e.Priority,
        "priorityScore": e.PriorityScore,
        "properties": e.Properties,
        "timeout": e.Timeout,
    }
}

// EncodeMargins converts a Margins struct to a map
func EncodeMargins(e Margins) map[string]interface{} {
    return map[string]interface{}{
        "direct": e.Direct,
        "premium": e.Premium,
        "standard": e.Standard,
    }
}




