package packing_slip_package

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/dimension"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to PackingSlipPackage struct
func Decode(m interface{}) PackingSlipPackage {
	if m, ok := m.(map[string]interface{}); ok {
		return PackingSlipPackage{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Dimensions: DecodeDimensions(m["dimensions"]),
			Kind: utils.DecodeString(m["_kind"], "PackingSlipPackage"),
			Name: utils.DecodeString(m["name"], ""),
			TranssmartCode: TranssmartCode(utils.DecodeString(m["transsmartCode"], "BOX")),
			Type: utils.DecodeString(m["type"], ""),
			Weight: dimension.Decode(m["weight"]),
		}
	}

	return PackingSlipPackage{}
}

// DecodeAll converts a slice of maps to a slice of PackingSlipPackage struct
func DecodeAll(ms interface{}) []PackingSlipPackage {
	var entities []PackingSlipPackage

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

// DecodeDimensions converts a map to Dimensions struct
func DecodeDimensions(m interface{}) Dimensions {
	if m == nil {
		return Dimensions{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Dimensions{
			Height: dimension.Decode(val["height"]),
			Length: dimension.Decode(val["length"]),
			Width: dimension.Decode(val["width"]),
		}
	}
	return Dimensions{}
}


