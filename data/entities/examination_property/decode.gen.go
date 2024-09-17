package examination_property

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to ExaminationProperty struct
func Decode(m interface{}) ExaminationProperty {
	if m, ok := m.(map[string]interface{}); ok {
		return ExaminationProperty{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AcceptanceRegion: DecodeAcceptanceRegion(m["acceptanceRegion"]),
			Description: utils.DecodeString(m["description"], ""),
			Examine: utils.DecodeBool(m["examine"], true),
			Info: utils.DecodeMap(m["info"], nil),
			InputType: ExaminationInputType(utils.DecodeString(m["inputType"], "value")),
			Instructions: utils.DecodeString(m["instructions"], ""),
			Key: utils.DecodeString(m["key"], ""),
			Kind: utils.DecodeString(m["_kind"], "ExaminationProperty"),
			Name: utils.DecodeString(m["name"], ""),
			ProductId: utils.DecodeString(m["productId"], ""),
			ProductSpecific: utils.DecodeBool(m["productSpecific"], false),
			Required: utils.DecodeBool(m["required"], true),
			Testable: utils.DecodeBool(m["testable"], true),
			Type: utils.DecodeString(m["type"], ""),
			UnitType: utils.DecodeString(m["unitType"], ""),
		}
	}

	return ExaminationProperty{}
}

// DecodeAll converts a slice of maps to a slice of ExaminationProperty struct
func DecodeAll(ms interface{}) []ExaminationProperty {
	var entities []ExaminationProperty

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

// DecodeAcceptanceRegion converts a map to AcceptanceRegion struct
func DecodeAcceptanceRegion(m interface{}) AcceptanceRegion {
	if m == nil {
		return AcceptanceRegion{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return AcceptanceRegion{
			Max: utils.DecodeFloat64(val["max"], 0.0),
			Min: utils.DecodeFloat64(val["min"], 0.0),
			Symmetric: utils.DecodeBool(val["symmetric"], false),
			Type: utils.DecodeString(val["type"], ""),
			Value: utils.DecodeFloat64(val["value"], 0.0),
		}
	}
	return AcceptanceRegion{}
}


