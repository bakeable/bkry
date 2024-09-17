package attribute

import (
	"github.com/bakeable/bkry/data/entities/attribute_option"
	"github.com/bakeable/bkry/data/entities/attribute_option_set"
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to Attribute struct
func Decode(m interface{}) Attribute {
	if m, ok := m.(map[string]interface{}); ok {
		return Attribute{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Default: utils.DecodeBool(m["default"], true),
			DefaultOptions: DecodeDefaultOptionArray(m["defaultOptions"]),
			Description: utils.DecodeString(m["description"], ""),
			Included: utils.DecodeBool(m["included"], false),
			IncrementalPricing: AttributeIncrementalPricing(utils.DecodeString(m["incrementalPricing"], "standard")),
			Key: utils.DecodeString(m["key"], ""),
			Kind: utils.DecodeString(m["_kind"], "Attribute"),
			MaxOptionID: utils.DecodeInt(m["maxOptionId"], 0),
			Name: utils.DecodeString(m["name"], ""),
			Nid: utils.DecodeInt(m["nid"], 0),
			OptionSets: attribute_option_set.DecodeAll(m["optionSets"]),
			Optional: utils.DecodeBool(m["optional"], false),
			Options: attribute_option.DecodeAll(m["options"]),
			Order: utils.DecodeInt(m["order"], 0),
			Prefix: utils.DecodeString(m["prefix"], ""),
			PriceTypes: utils.DecodeStringArray(m["priceTypes"], "variable"),
			Surcharges: utils.DecodeMap(m["surcharges"], map[string]interface{}{}),
			Type: AttributeType(utils.DecodeString(m["type"], "text")),
		}
	}

	return Attribute{}
}

// DecodeAll converts a slice of maps to a slice of Attribute struct
func DecodeAll(ms interface{}) []Attribute {
	var entities []Attribute

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

// DecodeDefaultOptionArray converts a map to DefaultOption struct
func DecodeDefaultOptionArray(m interface{}) []DefaultOption {
	if m == nil {
		return []DefaultOption{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []DefaultOption
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, DefaultOption{
				ID: utils.DecodeString(v["id"], ""),
				Key: utils.DecodeString(v["key"], ""),
				Max: utils.DecodeFloat64(v["max"], 0),
				Min: utils.DecodeFloat64(v["min"], 0),
				Value: utils.DecodeString(v["value"], ""),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []DefaultOption
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, DefaultOption{
					ID: utils.DecodeString(val["id"], ""),
					Key: utils.DecodeString(val["key"], ""),
					Max: utils.DecodeFloat64(val["max"], 0),
					Min: utils.DecodeFloat64(val["min"], 0),
					Value: utils.DecodeString(val["value"], ""),
				})
			}
		}
		return entities
	}
	return []DefaultOption{}
}


