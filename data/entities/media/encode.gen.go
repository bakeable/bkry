package media

// Encode converts a Media struct to a map
func Encode(e Media) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "contentType": e.ContentType,
        "description": e.Description,
        "extension": e.Extension,
        "filename": e.Filename,
        "mimeType": e.MimeType,
        "size": e.Size,
        "storagePath": e.StoragePath,
        "url": e.Url,
    }
    return result
}



