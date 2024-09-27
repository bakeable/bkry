package printing_order_item

// Encode converts a PrintingOrderItem struct to a map
func Encode(e PrintingOrderItem) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "_kind": e.Kind,
        "sku": e.Sku,
        "supplierContactName": e.SupplierContactName,
        "supplierId": e.SupplierId,
        "supplierName": e.SupplierName,
    }
    return result
}



