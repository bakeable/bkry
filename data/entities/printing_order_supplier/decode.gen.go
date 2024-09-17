package printing_order_supplier

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to PrintingOrderSupplier struct
func Decode(m interface{}) PrintingOrderSupplier {
	if m, ok := m.(map[string]interface{}); ok {
		return PrintingOrderSupplier{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Code: utils.DecodeString(m["code"], ""),
			ContactName: utils.DecodeString(m["contactName"], ""),
			DeliveryType: PrintingOrderSupplierDeliveryType(utils.DecodeString(m["deliveryType"], "EMAIL")),
			Email: utils.DecodeString(m["email"], ""),
			Emails: utils.DecodeStringArray(m["emails"], []string{}),
			Kind: utils.DecodeString(m["_kind"], "'PrintingOrderSupplier'"),
			Name: utils.DecodeString(m["name"], ""),
			Skus: utils.DecodeStringArray(m["skus"], []string{}),
		}
	}

	return PrintingOrderSupplier{}
}

// DecodeAll converts a slice of maps to a slice of PrintingOrderSupplier struct
func DecodeAll(ms interface{}) []PrintingOrderSupplier {
	var entities []PrintingOrderSupplier

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

