package batch_export

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to BatchExport struct
func Decode(m interface{}) BatchExport {
	if m, ok := m.(map[string]interface{}); ok {
		return BatchExport{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			ExportFromTimestamp: utils.DecodeInt64(m["exportFromTimestamp"], 0),
			Ids: utils.DecodeStringArray(m["ids"], []string{}),
			Kind: utils.DecodeString(m["_kind"], "Export"),
		}
	}

	return BatchExport{}
}

// DecodeAll converts a slice of maps to a slice of BatchExport struct
func DecodeAll(ms interface{}) []BatchExport {
	var entities []BatchExport

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

