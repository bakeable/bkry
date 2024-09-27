package printing_order_item

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to PrintingOrderItem struct
func Decode(m interface{}) PrintingOrderItem {
	if m, ok := m.(map[string]interface{}); ok {
		return PrintingOrderItem{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Kind: utils.DecodeString(m["_kind"], "PrintingOrderItem"),
			Sku: utils.DecodeString(m["sku"], ""),
			SupplierContactName: utils.DecodeString(m["supplierContactName"], ""),
			SupplierId: utils.DecodeString(m["supplierId"], ""),
			SupplierName: utils.DecodeString(m["supplierName"], ""),
		}
	}

	return PrintingOrderItem{}
}

// DecodeAll converts a slice of maps to a slice of PrintingOrderItem struct
func DecodeAll(ms interface{}) []PrintingOrderItem {
	var entities []PrintingOrderItem

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

