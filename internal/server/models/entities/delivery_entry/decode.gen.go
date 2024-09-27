package delivery_entry

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to DeliveryEntry struct
func Decode(m interface{}) DeliveryEntry {
	if m, ok := m.(map[string]interface{}); ok {
		return DeliveryEntry{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			CustomerId: utils.DecodeString(m["customerId"], ""),
			CustomerName: utils.DecodeString(m["customerName"], ""),
			Date: utils.DecodeString(m["date"], ""),
			Description: utils.DecodeString(m["description"], ""),
			Description2: utils.DecodeString(m["description_2"], ""),
			Examination: DecodeExamination(m["examination"]),
			ExaminationTasks: All(m["examinationTasks"]),
			ExtraDescription: utils.DecodeString(m["extraDescription"], ""),
			Guid: utils.DecodeString(m["guid"], ""),
			OrderNumber: utils.DecodeString(m["orderNumber"], ""),
			OrderType: utils.DecodeString(m["orderType"], ""),
			Product: (m["product"]),
			SerialNumber: utils.DecodeInt(m["serialNumber"], 0),
			Sku: utils.DecodeString(m["sku"], ""),
			UnitType: utils.DecodeString(m["unitType"], ""),
			Units: utils.DecodeInt(m["units"], 0),
			UnitsPerPallet: utils.DecodeInt(m["unitsPerPallet"], 0),
			Warehouse: utils.DecodeString(m["warehouse"], ""),
			_Kind: utils.DecodeString(m["_kind"], "DeliveryEntry"),
		}
	}

	return DeliveryEntry{}
}

// DecodeAll converts a slice of maps to a slice of DeliveryEntry struct
func DecodeAll(ms interface{}) []DeliveryEntry {
	var entities []DeliveryEntry

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

// DecodeExamination converts a map to Examination struct
func DecodeExamination(m interface{}) Examination {
	if m == nil {
		return Examination{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Examination{
			Priority: utils.DecodeBool(val["priority"], false),
			PriorityScore: utils.DecodeInt(val["priorityScore"], 0),
			Properties: utils.DecodeMap(val["properties"], nil),
			Timeout: utils.DecodeInt(val["timeout"], 0),
		}
	}
	return Examination{}
}


