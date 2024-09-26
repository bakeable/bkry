package airline_order

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to AirlineOrder struct
func Decode(m interface{}) AirlineOrder {
	if m, ok := m.(map[string]interface{}); ok {
		return AirlineOrder{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AfasOrderNumber: utils.DecodeString(m["afasOrderNumber"], ""),
			Colors: utils.DecodeInt(m["colors"], 0),
			CustomerId: utils.DecodeString(m["customerId"], ""),
			CustomerName: utils.DecodeString(m["customerName"], ""),
			Date: utils.DecodeString(m["date"], "new Date()"),
			DeliveryAddress: utils.DecodeString(m["deliveryAddress"], ""),
			Description: utils.DecodeString(m["description"], ""),
			Dividers: utils.DecodeInt(m["dividers"], 0),
			Guid: utils.DecodeString(m["guid"], ""),
			Inkooprelatie: utils.DecodeString(m["supplierId"], ""),
			IsGrouped: utils.DecodeBool(m["isGrouped"], false),
			Kind: utils.DecodeString(m["_kind"], "AirlineOrder"),
			Margin: DecodeMargin(m["margin"]),
			Nummer: utils.DecodeString(m["number"], ""),
			OrderNumber: utils.DecodeString(m["orderNumber"], ""),
			OriginalTotalCosts: utils.DecodeInt(m["originalTotalCosts"], 0),
			OriginalUnitCosts: utils.DecodeInt(m["originalUnitCosts"], 0),
			SerialNumber: utils.DecodeInt(m["serialNumber"], 0),
			Sku: utils.DecodeString(m["sku"], ""),
			Status: utils.DecodeString(m["status"], ""),
			TotalCosts: utils.DecodeFloat64(m["totalCosts"], 0.0),
			UnitCosts: utils.DecodeFloat64(m["unitCosts"], 0.0),
			UnitType: utils.DecodeString(m["unitType"], ""),
			Units: utils.DecodeInt(m["units"], 0),
			Warehouse: utils.DecodeString(m["warehouse"], ""),
			Width: utils.DecodeInt(m["width"], 0),
			WidthUnitType: utils.DecodeInt(m["widthUnitType"], 0),
		}
	}

	return AirlineOrder{}
}

// DecodeAll converts a slice of maps to a slice of AirlineOrder struct
func DecodeAll(ms interface{}) []AirlineOrder {
	var entities []AirlineOrder

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

// DecodeMargin converts a map to Margin struct
func DecodeMargin(m interface{}) Margin {
	if m == nil {
		return Margin{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Margin{
			Amount: utils.DecodeInt(val["amount"], 0),
			Percentage: utils.DecodeInt(val["percentage"], 0),
			PercentageType: utils.DecodeString(val["percentageType"], "profit-margin"),
			Staggered: utils.DecodeBool(val["staggered"], false),
			Type: utils.DecodeString(val["type"], "percentage"),
			Values: utils.DecodeInterfaceArray(val["values"], []interface{}{}),
		}
	}
	return Margin{}
}


