package airline_settings

// Encode converts a AirlineSettings struct to a map
func Encode(e AirlineSettings) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "clicheCosts": e.ClicheCosts,
        "_kind": e.Kind,
        "metersPerUnit": e.MetersPerUnit,
        "rollLength": e.RollLength,
        "rollWidthDividerMap": e.RollWidthDividerMap,
        "rollWidthOptions": e.RollWidthOptions,
        "sleeveCosts": e.SleeveCosts,
        "switchCosts": e.SwitchCosts,
        "unitOptionSets": e.UnitOptionSets,
    }
    return result
}



