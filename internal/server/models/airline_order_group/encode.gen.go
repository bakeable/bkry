package airline_order_group

// Encode converts a AirlineOrderGroup struct to a map
func Encode(e AirlineOrderGroup) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "afasOrderNumbers": e.AfasOrderNumbers,
        "_kind": e.Kind,
        "name": e.Name,
        "orders": e.Orders,
        "pricingTables": e.PricingTables,
        "settings": e.Settings,
    }
    return result
}



