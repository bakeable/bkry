package examination_task

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to ExaminationTask struct
func Decode(m interface{}) ExaminationTask {
	if m, ok := m.(map[string]interface{}); ok {
		return ExaminationTask{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Approved: utils.DecodeBool(m["approved"], false),
			AssignedTo: ExaminationTaskRole(utils.DecodeString(m["assignedTo"], "observer")),
			ChangeLog: DecodeChangeLogEntryArray(m["changeLog"]),
			Kind: utils.DecodeString(m["kind"], "ExaminationTask"),
			Observations: utils.DecodeMap(m["observations"], nil),
			ProductId: utils.DecodeString(m["productId"], ""),
			ProductName: utils.DecodeString(m["productName"], ""),
			Properties: DecodePropertyArray(m["properties"]),
			Sku: utils.DecodeString(m["sku"], ""),
			Status: ExaminationTaskStatus(utils.DecodeString(m["status"], "concept")),
			StatusIndex: utils.DecodeInt(m["statusIndex"], 0),
		}
	}

	return ExaminationTask{}
}

// DecodeAll converts a slice of maps to a slice of ExaminationTask struct
func DecodeAll(ms interface{}) []ExaminationTask {
	var entities []ExaminationTask

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

// DecodeChangeLogEntryArray converts a map to ChangeLogEntry struct
func DecodeChangeLogEntryArray(m interface{}) []ChangeLogEntry {
	if m == nil {
		return []ChangeLogEntry{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []ChangeLogEntry
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, ChangeLogEntry{
				Action: ExaminationTaskAction(utils.DecodeString(v["action"], "created")),
				Status: utils.DecodeString(v["status"], ""),
				Timestamp: utils.DecodeString(v["timestamp"], ""),
				User: DecodeUser(v["user"]),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []ChangeLogEntry
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, ChangeLogEntry{
					Action: ExaminationTaskAction(utils.DecodeString(val["action"], "created")),
					Status: utils.DecodeString(val["status"], ""),
					Timestamp: utils.DecodeString(val["timestamp"], ""),
					User: DecodeUser(val["user"]),
				})
			}
		}
		return entities
	}
	return []ChangeLogEntry{}
}

// DecodeUser converts a map to User struct
func DecodeUser(m interface{}) User {
	if m == nil {
		return User{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return User{
			Email: utils.DecodeString(val["email"], ""),
			ID: utils.DecodeString(val["id"], ""),
		}
	}
	return User{}
}



// DecodePropertyArray converts a map to Property struct
func DecodePropertyArray(m interface{}) []Property {
	if m == nil {
		return []Property{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []Property
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, Property{
				AcceptanceRegion: DecodeAcceptanceRegion(v["acceptanceRegion"]),
				Description: utils.DecodeString(v["description"], ""),
				ExpectedValue: utils.DecodeString(v["expectedValue"], ""),
				InputType: utils.DecodeString(v["inputType"], ""),
				Instructions: utils.DecodeString(v["instructions"], ""),
				Key: utils.DecodeString(v["key"], ""),
				LowerBound: utils.DecodeFloat64(v["lowerBound"], 0),
				Name: utils.DecodeString(v["name"], ""),
				ProductSpecific: utils.DecodeBool(v["productSpecific"], false),
				Required: utils.DecodeBool(v["required"], false),
				Type: utils.DecodeString(v["type"], ""),
				UnitType: utils.DecodeString(v["unitType"], ""),
				UpperBound: utils.DecodeFloat64(v["upperBound"], 0),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []Property
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, Property{
					AcceptanceRegion: DecodeAcceptanceRegion(val["acceptanceRegion"]),
					Description: utils.DecodeString(val["description"], ""),
					ExpectedValue: utils.DecodeString(val["expectedValue"], ""),
					InputType: utils.DecodeString(val["inputType"], ""),
					Instructions: utils.DecodeString(val["instructions"], ""),
					Key: utils.DecodeString(val["key"], ""),
					LowerBound: utils.DecodeFloat64(val["lowerBound"], 0),
					Name: utils.DecodeString(val["name"], ""),
					ProductSpecific: utils.DecodeBool(val["productSpecific"], false),
					Required: utils.DecodeBool(val["required"], false),
					Type: utils.DecodeString(val["type"], ""),
					UnitType: utils.DecodeString(val["unitType"], ""),
					UpperBound: utils.DecodeFloat64(val["upperBound"], 0),
				})
			}
		}
		return entities
	}
	return []Property{}
}

// DecodeAcceptanceRegion converts a map to AcceptanceRegion struct
func DecodeAcceptanceRegion(m interface{}) AcceptanceRegion {
	if m == nil {
		return AcceptanceRegion{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return AcceptanceRegion{
			Max: utils.DecodeFloat64(val["max"], 0),
			Min: utils.DecodeFloat64(val["min"], 0),
			Symmetric: utils.DecodeBool(val["symmetric"], false),
			Type: AcceptanceRegionType(utils.DecodeString(val["type"], "percentage")),
			Value: utils.DecodeFloat64(val["value"], 0),
		}
	}
	return AcceptanceRegion{}
}



