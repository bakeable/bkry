package batch_export

// Encode converts a BatchExport struct to a map
func Encode(e BatchExport) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "exportFromTimestamp": e.ExportFromTimestamp,
        "ids": e.Ids,
        "_kind": e.Kind,
    }
    return result
}



