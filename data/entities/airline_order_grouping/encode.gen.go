package airline_order_grouping

// Encode converts a AirlineOrderGrouping struct to a map
func Encode(e AirlineOrderGrouping) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "afasOrderNumbers": e.AfasOrderNumbers,
        "finalPrice": e.FinalPrice,
        "firstOrderDate": e.FirstOrderDate,
        "_kind": e.Kind,
        "marginAmount": e.MarginAmount,
        "marginPercentage": e.MarginPercentage,
        "name": e.Name,
        "orderDate": e.OrderDate,
        "orderGroups": e.OrderGroups,
        "orderIds": e.OrderIds,
        "originalTotalCosts": e.OriginalTotalCosts,
        "profitAmount": e.ProfitAmount,
        "profitPercentage": e.ProfitPercentage,
        "savedCosts": e.SavedCosts,
        "sentToAfas": e.SentToAfas,
        "sentToAfasTimestamp": e.SentToAfasTimestamp,
        "totalCosts": e.TotalCosts,
        "units": e.Units,
    }
    return result
}



