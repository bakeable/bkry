package packing_slip

// Encode converts a PackingSlip struct to a map
func Encode(e PackingSlip) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "addedBy": e.AddedBy,
        "address": EncodeAddress(e.Address),
        "administration": e.Administration,
        "awb": e.Awb,
        "carrierCode": e.CarrierCode,
        "combinedOrderNumbers": e.CombinedOrderNumbers,
        "combinedPackingSlipIds": e.CombinedPackingSlipIds,
        "companyName": e.CompanyName,
        "customer": EncodeCustomer(e.Customer),
        "date": e.Date,
        "delivery": EncodeDelivery(e.Delivery),
        "description": e.Description,
        "incotermCode": e.IncotermCode,
        "_kind": e.Kind,
        "mainPackingSlipId": e.MainPackingSlipId,
        "markedAsCompleted": e.MarkedAsCompleted,
        "markedAsPrinted": e.MarkedAsPrinted,
        "markedAsPushed": e.MarkedAsPushed,
        "note": e.Note,
        "orderCreatedBy": e.OrderCreatedBy,
        "orderNumber": e.OrderNumber,
        "orderProcessing": e.OrderProcessing,
        "originalTotalGrossWeight": e.OriginalTotalGrossWeight,
        "packages": EncodePackageArray(e.Packages),
        "promisedDeliveryDate": e.PromisedDeliveryDate,
        "reference": e.Reference,
        "service": e.Service,
        "slipNumber": e.SlipNumber,
        "status": e.Status,
        "statusIndex": e.StatusIndex,
        "supplierContactName": e.SupplierContactName,
        "supplierId": e.SupplierID,
        "supplierName": e.SupplierName,
        "totalAmountExcludingVAT": e.TotalAmountExcludingVAT,
        "totalGrossWeight": e.TotalGrossWeight,
        "totalQuantity": e.TotalQuantity,
        "trackingUrl": e.TrackingUrl,
    }
    return result
}
// EncodeAddress converts a Address struct to a map
func EncodeAddress(e Address) map[string]interface{} {
    return map[string]interface{}{
        "city": e.City,
        "country": e.Country,
        "description": e.Description,
        "formatted": e.Formatted,
        "houseNumber": e.HouseNumber,
        "postalCode": e.PostalCode,
        "street": e.Street,
    }
}

// EncodeCustomer converts a Customer struct to a map
func EncodeCustomer(e Customer) map[string]interface{} {
    return map[string]interface{}{
        "email": e.Email,
        "name": e.Name,
        "phoneNumber": e.PhoneNumber,
    }
}

// EncodeDelivery converts a Delivery struct to a map
func EncodeDelivery(e Delivery) map[string]interface{} {
    return map[string]interface{}{
        "afterTime": EncodeAfterTime(e.AfterTime),
        "beforeTime": EncodeBeforeTime(e.BeforeTime),
        "date": e.Date,
        "notes": e.Notes,
    }
}
// EncodeAfterTime converts a AfterTime struct to a map
func EncodeAfterTime(e AfterTime) map[string]interface{} {
    return map[string]interface{}{
        "hour": e.Hour,
        "minute": e.Minute,
    }
}

// EncodeBeforeTime converts a BeforeTime struct to a map
func EncodeBeforeTime(e BeforeTime) map[string]interface{} {
    return map[string]interface{}{
        "hour": e.Hour,
        "minute": e.Minute,
    }
}


// EncodePackageArray converts a Package struct to a map
func EncodePackageArray(e []Package) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "amount": v.Amount,
            "dimensions": EncodeDimensions(v.Dimensions),
            "transsmartCode": v.TranssmartCode,
            "type": v.Type,
            "weight": v.Weight,
        }
    }
    return result
}
// EncodeDimensions converts a Dimensions struct to a map
func EncodeDimensions(e Dimensions) map[string]interface{} {
    return map[string]interface{}{
        "height": e.Height,
        "length": e.Length,
        "width": e.Width,
    }
}





