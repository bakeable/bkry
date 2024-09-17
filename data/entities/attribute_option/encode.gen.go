package attribute_option

// Encode converts a AttributeOption struct to a map
func Encode(e AttributeOption) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "key": e.Key,
        "max": e.Max,
        "min": e.Min,
        "value": e.Value,
    }
    return result
}



