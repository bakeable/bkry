package email

// Encode converts a Email struct to a map
func Encode(e Email) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "_kind": e.Kind,
    }
    return result
}



