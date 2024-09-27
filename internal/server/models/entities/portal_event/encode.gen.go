package portal_event

// Encode converts a PortalEvent struct to a map
func Encode(e PortalEvent) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "entityId": e.EntityID,
        "entityName": e.EntityName,
        "_kind": e.Kind,
        "metadata": e.Metadata,
        "timestamp": e.Timestamp,
        "timezone": e.Timezone,
        "type": e.Type,
        "userId": e.UserId,
    }
    return result
}



