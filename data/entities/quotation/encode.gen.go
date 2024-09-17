package quotation

// Encode converts a Quotation struct to a map
func Encode(e Quotation) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "customer": EncodeCustomer(e.Customer),
        "_kind": e.Kind,
        "lineItems": EncodeLineItemArray(e.LineItems),
        "name": e.Name,
        "notes": e.Notes,
        "referralCode": e.ReferralCode,
        "status": e.Status,
        "statusIndex": e.StatusIndex,
        "totals": EncodeTotals(e.Totals),
    }
    return result
}
// EncodeCustomer converts a Customer struct to a map
func EncodeCustomer(e Customer) map[string]interface{} {
    return map[string]interface{}{
        "address": e.Address,
        "email": e.Email,
        "name": e.Name,
        "phone": e.Phone,
    }
}

// EncodeLineItemArray converts a LineItem struct to a map
func EncodeLineItemArray(e []LineItem) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "attributes": v.Attributes,
            "finalPrice": v.FinalPrice,
            "fixedCosts": v.FixedCosts,
            "marginAmount": v.MarginAmount,
            "marginPercentage": v.MarginPercentage,
            "product": EncodeProduct(v.Product),
            "totalCosts": v.TotalCosts,
            "units": v.Units,
            "variableCosts": v.VariableCosts,
        }
    }
    return result
}
// EncodeProduct converts a Product struct to a map
func EncodeProduct(e Product) map[string]interface{} {
    return map[string]interface{}{
        "id": e.ID,
        "name": e.Name,
        "sku": e.Sku,
    }
}


// EncodeTotals converts a Totals struct to a map
func EncodeTotals(e Totals) map[string]interface{} {
    return map[string]interface{}{
        "finalPrice": e.FinalPrice,
        "marginAmount": e.MarginAmount,
        "marginPercentage": e.MarginPercentage,
        "totalCosts": e.TotalCosts,
    }
}




