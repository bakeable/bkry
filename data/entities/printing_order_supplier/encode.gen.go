package printing_order_supplier

// Encode converts a PrintingOrderSupplier struct to a map
func Encode(e PrintingOrderSupplier) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "code": e.Code,
        "contactName": e.ContactName,
        "deliveryType": e.DeliveryType,
        "email": e.Email,
        "emails": e.Emails,
        "_kind": e.Kind,
        "name": e.Name,
        "skus": e.Skus,
    }
    return result
}



