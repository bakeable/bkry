package examination_settings

// Encode converts a ExaminationSettings struct to a map
func Encode(e ExaminationSettings) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "exclusionTime": e.ExclusionTime,
        "exclusionTimeFrame": EncodeExclusionTimeFrame(e.ExclusionTimeFrame),
        "kind": e.Kind,
        "maxOpenTasks": e.MaxOpenTasks,
        "maxProducts": e.MaxProducts,
        "maxProperties": e.MaxProperties,
        "minProducts": e.MinProducts,
        "minProperties": e.MinProperties,
        "properties": e.Properties,
        "selectionTime": e.SelectionTime,
        "selectionTimeFrame": EncodeSelectionTimeFrame(e.SelectionTimeFrame),
    }
    return result
}
// EncodeExclusionTimeFrame converts a ExclusionTimeFrame struct to a map
func EncodeExclusionTimeFrame(e ExclusionTimeFrame) map[string]interface{} {
    return map[string]interface{}{
        "days": e.Days,
        "hours": e.Hours,
    }
}

// EncodeSelectionTimeFrame converts a SelectionTimeFrame struct to a map
func EncodeSelectionTimeFrame(e SelectionTimeFrame) map[string]interface{} {
    return map[string]interface{}{
        "days": e.Days,
        "hours": e.Hours,
    }
}




