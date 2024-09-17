package packing_slip_package

// Encode converts a PackingSlipPackage struct to a map
func Encode(e PackingSlipPackage) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "dimensions": EncodeDimensions(e.Dimensions),
        "_kind": e.Kind,
        "name": e.Name,
        "transsmartCode": e.TranssmartCode,
        "type": e.Type,
        "weight": e.Weight,
    }
    return result
}
// EncodeDimensions converts a Dimensions struct to a map
func EncodeDimensions(e Dimensions) map[string]interface{} {
    return map[string]interface{}{
        "height": e.Height,
        "length": e.Length,
        "width": e.Width,
    }
}




