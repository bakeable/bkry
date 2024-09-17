package examination_task

// Encode converts a ExaminationTask struct to a map
func Encode(e ExaminationTask) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "approved": e.Approved,
        "assignedTo": e.AssignedTo,
        "changeLog": EncodeChangeLogEntryArray(e.ChangeLog),
        "kind": e.Kind,
        "observations": e.Observations,
        "productId": e.ProductId,
        "productName": e.ProductName,
        "properties": EncodePropertyArray(e.Properties),
        "sku": e.Sku,
        "status": e.Status,
        "statusIndex": e.StatusIndex,
    }
    return result
}
// EncodeChangeLogEntryArray converts a ChangeLogEntry struct to a map
func EncodeChangeLogEntryArray(e []ChangeLogEntry) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "action": v.Action,
            "status": v.Status,
            "timestamp": v.Timestamp,
            "user": EncodeUser(v.User),
        }
    }
    return result
}
// EncodeUser converts a User struct to a map
func EncodeUser(e User) map[string]interface{} {
    return map[string]interface{}{
        "email": e.Email,
        "id": e.ID,
    }
}


// EncodePropertyArray converts a Property struct to a map
func EncodePropertyArray(e []Property) []map[string]interface{} {
    result := make([]map[string]interface{}, len(e))
    for i, v := range e {
        result[i] = map[string]interface{}{
            "acceptanceRegion": EncodeAcceptanceRegion(v.AcceptanceRegion),
            "description": v.Description,
            "expectedValue": v.ExpectedValue,
            "inputType": v.InputType,
            "instructions": v.Instructions,
            "key": v.Key,
            "lowerBound": v.LowerBound,
            "name": v.Name,
            "productSpecific": v.ProductSpecific,
            "required": v.Required,
            "type": v.Type,
            "unitType": v.UnitType,
            "upperBound": v.UpperBound,
        }
    }
    return result
}
// EncodeAcceptanceRegion converts a AcceptanceRegion struct to a map
func EncodeAcceptanceRegion(e AcceptanceRegion) map[string]interface{} {
    return map[string]interface{}{
        "max": e.Max,
        "min": e.Min,
        "symmetric": e.Symmetric,
        "type": e.Type,
        "value": e.Value,
    }
}





