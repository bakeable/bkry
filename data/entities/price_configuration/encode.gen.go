package price_configuration

// Encode converts a PriceConfiguration struct to a map
func Encode(e PriceConfiguration) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "additional_costs": e.AdditionalCosts,
        "attributes": e.Attributes,
        "currency": e.Currency,
        "final_price": e.FinalPrice,
        "fixed_costs": e.FixedCosts,
        "_kind": e.Kind,
        "margin_amount": e.MarginAmount,
        "margin_percentage": e.MarginPercentage,
        "product_id": e.ProductId,
        "sku": e.Sku,
        "total_costs": e.TotalCosts,
        "unit_costs": e.UnitCosts,
        "unit_margin": e.UnitMargin,
        "units": e.Units,
        "variable_costs": e.VariableCosts,
    }
    return result
}



