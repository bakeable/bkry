package delivery_entry

// Encode converts a DeliveryEntry struct to a map
func Encode(e DeliveryEntry) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "customerId": e.CustomerId,
        "customerName": e.CustomerName,
        "date": e.Date,
        "description": e.Description,
        "description_2": e.Description2,
        "examination": EncodeExamination(e.Examination),
        "examinationTasks": e.ExaminationTasks,
        "extraDescription": e.ExtraDescription,
        "guid": e.Guid,
        "orderNumber": e.OrderNumber,
        "orderType": e.OrderType,
        "product": e.Product,
        "serialNumber": e.SerialNumber,
        "sku": e.Sku,
        "unitType": e.UnitType,
        "units": e.Units,
        "unitsPerPallet": e.UnitsPerPallet,
        "warehouse": e.Warehouse,
        "_kind": e._Kind,
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




