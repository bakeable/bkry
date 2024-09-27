package cloud_function

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to CloudFunction struct
func Decode(m interface{}) CloudFunction {
	if m, ok := m.(map[string]interface{}); ok {
		return CloudFunction{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Description: utils.DecodeString(m["description"], ""),
			Kind: utils.DecodeString(m["_kind"], "CloudFunction"),
			Name: utils.DecodeString(m["name"], ""),
			Region: utils.DecodeString(m["region"], ""),
			Runtime: utils.DecodeString(m["runtime"], ""),
			TriggerType: utils.DecodeString(m["triggerType"], ""),
			TriggerValue: utils.DecodeString(m["triggerValue"], ""),
		}
	}

	return CloudFunction{}
}

// DecodeAll converts a slice of maps to a slice of CloudFunction struct
func DecodeAll(ms interface{}) []CloudFunction {
	var entities []CloudFunction

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

