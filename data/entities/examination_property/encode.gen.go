package examination_property

// Encode converts a ExaminationProperty struct to a map
func Encode(e ExaminationProperty) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "acceptanceRegion": EncodeAcceptanceRegion(e.AcceptanceRegion),
        "description": e.Description,
        "examine": e.Examine,
        "info": e.Info,
        "inputType": e.InputType,
        "instructions": e.Instructions,
        "key": e.Key,
        "_kind": e.Kind,
        "name": e.Name,
        "productId": e.ProductId,
        "productSpecific": e.ProductSpecific,
        "required": e.Required,
        "testable": e.Testable,
        "type": e.Type,
        "unitType": e.UnitType,
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




