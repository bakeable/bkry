package airline_order

// Encode converts a AirlineOrder struct to a map
func Encode(e AirlineOrder) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "afasOrderNumber": e.AfasOrderNumber,
        "colors": e.Colors,
        "customerId": e.CustomerId,
        "customerName": e.CustomerName,
        "date": e.Date,
        "deliveryAddress": e.DeliveryAddress,
        "description": e.Description,
        "dividers": e.Dividers,
        "guid": e.Guid,
        "supplierId": e.Inkooprelatie,
        "isGrouped": e.IsGrouped,
        "_kind": e.Kind,
        "margin": EncodeMargin(e.Margin),
        "number": e.Nummer,
        "orderNumber": e.OrderNumber,
        "originalTotalCosts": e.OriginalTotalCosts,
        "originalUnitCosts": e.OriginalUnitCosts,
        "serialNumber": e.SerialNumber,
        "sku": e.Sku,
        "status": e.Status,
        "totalCosts": e.TotalCosts,
        "unitCosts": e.UnitCosts,
        "unitType": e.UnitType,
        "units": e.Units,
        "warehouse": e.Warehouse,
        "width": e.Width,
        "widthUnitType": e.WidthUnitType,
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




