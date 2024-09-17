package examination_settings

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to ExaminationSettings struct
func Decode(m interface{}) ExaminationSettings {
	if m, ok := m.(map[string]interface{}); ok {
		return ExaminationSettings{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			ExclusionTime: utils.DecodeInt64(m["exclusionTime"], 2592000000),
			ExclusionTimeFrame: DecodeExclusionTimeFrame(m["exclusionTimeFrame"]),
			Kind: utils.DecodeString(m["kind"], "'ExaminationSettings'"),
			MaxOpenTasks: utils.DecodeInt(m["maxOpenTasks"], 20),
			MaxProducts: utils.DecodeInt(m["maxProducts"], 10),
			MaxProperties: utils.DecodeInt(m["maxProperties"], 5),
			MinProducts: utils.DecodeInt(m["minProducts"], 1),
			MinProperties: utils.DecodeInt(m["minProperties"], 1),
			Properties: utils.DecodeStringArray(m["properties"], []string{}),
			SelectionTime: utils.DecodeInt64(m["selectionTime"], 604800000),
			SelectionTimeFrame: DecodeSelectionTimeFrame(m["selectionTimeFrame"]),
		}
	}

	return ExaminationSettings{}
}

// DecodeAll converts a slice of maps to a slice of ExaminationSettings struct
func DecodeAll(ms interface{}) []ExaminationSettings {
	var entities []ExaminationSettings

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

// DecodeExclusionTimeFrame converts a map to ExclusionTimeFrame struct
func DecodeExclusionTimeFrame(m interface{}) ExclusionTimeFrame {
	if m == nil {
		return ExclusionTimeFrame{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return ExclusionTimeFrame{
			Days: utils.DecodeInt(val["days"], 0),
			Hours: utils.DecodeInt(val["hours"], 0),
		}
	}
	return ExclusionTimeFrame{}
}


// DecodeSelectionTimeFrame converts a map to SelectionTimeFrame struct
func DecodeSelectionTimeFrame(m interface{}) SelectionTimeFrame {
	if m == nil {
		return SelectionTimeFrame{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return SelectionTimeFrame{
			Days: utils.DecodeInt(val["days"], 0),
			Hours: utils.DecodeInt(val["hours"], 0),
		}
	}
	return SelectionTimeFrame{}
}


