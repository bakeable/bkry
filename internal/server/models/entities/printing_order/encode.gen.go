package printing_order

// Encode converts a PrintingOrder struct to a map
func Encode(e PrintingOrder) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "approved": e.Approved,
        "approvedTimestamp": e.ApprovedTimestamp,
        "artworkApproved": e.ArtworkApproved,
        "artworkIsUrl": e.ArtworkIsUrl,
        "autoCorrespondence": e.AutoCorrespondence,
        "companyName": e.CompanyName,
        "correspondence": EncodePrintingOrderCorrespondenceArray(e.Correspondence),
        "customerAddress": EncodeCustomerAddress(e.CustomerAddress),
        "customerEmail": e.CustomerEmail,
        "customerName": e.CustomerName,
        "description": e.Description,
        "emailCode": e.EmailCode,
        "files": EncodePrintingOrderFiles(e.Files),
        "isAirtableOrder": e.IsAirtableOrder,
        "itemIndex": e.ItemIndex,
        "_kind": e.Kind,
        "notes": e.Notes,
        "numberOfUnits": e.NumberOfUnits,
        "orderNumber": e.OrderNumber,
        "orderStatus": e.OrderStatus,
        "productOptions": EncodeProductOptionArray(e.ProductOptions),
        "quantity": e.Quantity,
        "quantityOrdered": e.QuantityOrdered,
        "rejected": e.Rejected,
        "rejectionReason": e.RejectionReason,
        "sku": e.Sku,
        "status": e.Status,
        "statusIndex": e.StatusIndex,
        "statusLog": EncodePrintingOrderStatusLogArray(e.StatusLog),
        "supplierContactName": e.SupplierContactName,
        "supplierId": e.SupplierId,
        "supplierName": e.SupplierName,
        "supplierReference": e.SupplierReference,
        "trackingUrl": e.TrackingUrl,
        "trackingUrlApproved": e.TrackingUrlApproved,
    }
    return result
}
// EncodePrintingOrderCorrespondenceArray converts a PrintingOrderCorrespondence struct to a map
func EncodePrintingOrderCorrespondenceArray(e []PrintingOrderCorrespondence) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "id": v.Id,
            "key": v.Key,
            "sent": v.Sent,
            "sentAt": v.SentAt,
            "statusLogIndex": v.StatusLogIndex,
            "timestamp": v.Timestamp,
        }
    }
    return result
}

// EncodeCustomerAddress converts a CustomerAddress struct to a map
func EncodeCustomerAddress(e CustomerAddress) map[string]interface{} {
    return map[string]interface{}{
        "city": e.City,
        "country": e.Country,
        "line": e.Line,
        "postalCode": e.PostalCode,
    }
}

// EncodePrintingOrderFiles converts a PrintingOrderFiles struct to a map
func EncodePrintingOrderFiles(e PrintingOrderFiles) map[string]interface{} {
    return map[string]interface{}{
        "artwork": e.Artwork,
        "media": e.Media,
        "orderConfirmation": e.OrderConfirmation,
    }
}

// EncodeProductOptionArray converts a ProductOption struct to a map
func EncodeProductOptionArray(e []ProductOption) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "id": v.Id,
            "key": v.Key,
            "value": v.Value,
        }
    }
    return result
}

// EncodePrintingOrderStatusLogArray converts a PrintingOrderStatusLog struct to a map
func EncodePrintingOrderStatusLogArray(e []PrintingOrderStatusLog) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "prevStatus": v.PrevStatus,
            "status": v.Status,
            "timestamp": v.Timestamp,
        }
    }
    return result
}




