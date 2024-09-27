package api_key

// Encode converts a ApiKey struct to a map
func Encode(e ApiKey) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "_kind": e.Kind,
        "publicKey": e.publicKey,
        "secretKey": e.secretKey,
        "userId": e.userId,
    }
    return result
}



