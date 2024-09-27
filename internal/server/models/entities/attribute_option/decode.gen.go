package attribute_option

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to AttributeOption struct
func Decode(m interface{}) AttributeOption {
	if m, ok := m.(map[string]interface{}); ok {
		return AttributeOption{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Key: utils.DecodeString(m["key"], ""),
			Max: utils.DecodeFloat64(m["max"], 0),
			Min: utils.DecodeFloat64(m["min"], 0),
			Value: utils.DecodeString(m["value"], ""),
		}
	}

	return AttributeOption{}
}

// DecodeAll converts a slice of maps to a slice of AttributeOption struct
func DecodeAll(ms interface{}) []AttributeOption {
	var entities []AttributeOption

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

