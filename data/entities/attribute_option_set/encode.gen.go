package attribute_option_set

// Encode converts a AttributeOptionSet struct to a map
func Encode(e AttributeOptionSet) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "label": e.Label,
        "options": e.Options,
    }
    return result
}



