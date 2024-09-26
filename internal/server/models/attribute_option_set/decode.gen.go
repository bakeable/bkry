package attribute_option_set

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/entities/attribute_option"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to AttributeOptionSet struct
func Decode(m interface{}) AttributeOptionSet {
	if m, ok := m.(map[string]interface{}); ok {
		return AttributeOptionSet{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Label: utils.DecodeString(m["label"], ""),
			Options: attribute_option.DecodeAll(m["options"]),
		}
	}

	return AttributeOptionSet{}
}

// DecodeAll converts a slice of maps to a slice of AttributeOptionSet struct
func DecodeAll(ms interface{}) []AttributeOptionSet {
	var entities []AttributeOptionSet

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

