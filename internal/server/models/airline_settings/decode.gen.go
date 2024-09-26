package airline_settings

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to AirlineSettings struct
func Decode(m interface{}) AirlineSettings {
	if m, ok := m.(map[string]interface{}); ok {
		return AirlineSettings{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			ClicheCosts: utils.DecodeFloat64(m["clicheCosts"], 295.0),
			Kind: utils.DecodeString(m["_kind"], "'AirlineSettings'"),
			MetersPerUnit: utils.DecodeInt(m["metersPerUnit"], 1000),
			RollLength: utils.DecodeInt(m["rollLength"], 1000),
			RollWidthDividerMap: utils.DecodeIntIntMap(m["rollWidthDividerMap"], map[int]int{175:5, 205:4, 210:4, 215:4, 220:4, 230:3, 265:3, 275:3, 300:2}),
			RollWidthOptions: utils.DecodeIntArray(m["rollWidthOptions"], []int{175, 205, 210, 215, 220, 230, 265, 275, 300}),
			SleeveCosts: utils.DecodeFloat64(m["sleeveCosts"], 650.0),
			SwitchCosts: utils.DecodeFloat64(m["switchCosts"], 200.0),
			UnitOptionSets: utils.DecodeIntArrayArray(m["unitOptionSets"], [][]int{{10, 15, 20, 25, 30, 35, 45, 60, 80, 100, 125, 150, 185, 215, 250, 300, 400, 500},{8, 12, 16, 20, 24, 32, 40, 52, 64, 84, 104, 128, 152, 192, 240, 300, 400, 500},{9, 12, 15, 18, 24, 30, 39, 48, 60, 75, 99, 126, 150, 198, 249, 300, 400, 500},{4, 6, 8, 10, 12, 16, 20, 28, 36, 48, 68, 88, 116, 152, 200, 300, 400, 500}}),
		}
	}

	return AirlineSettings{}
}

// DecodeAll converts a slice of maps to a slice of AirlineSettings struct
func DecodeAll(ms interface{}) []AirlineSettings {
	var entities []AirlineSettings

	if arr, ok := ms.([]map[string]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	if arr, ok := ms.([]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	return entities
}

