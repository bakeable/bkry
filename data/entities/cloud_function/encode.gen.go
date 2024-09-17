package cloud_function

// Encode converts a CloudFunction struct to a map
func Encode(e CloudFunction) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "description": e.Description,
        "_kind": e.Kind,
        "name": e.Name,
        "region": e.Region,
        "runtime": e.Runtime,
        "triggerType": e.TriggerType,
        "triggerValue": e.TriggerValue,
    }
    return result
}



