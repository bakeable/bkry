package printing_order

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/entities/media"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to PrintingOrder struct
func Decode(m interface{}) PrintingOrder {
	if m, ok := m.(map[string]interface{}); ok {
		return PrintingOrder{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Approved: utils.DecodeBool(m["approved"], false),
			ApprovedTimestamp: utils.DecodeInt64(m["approvedTimestamp"], 0),
			ArtworkApproved: utils.DecodeBool(m["artworkApproved"], false),
			ArtworkIsUrl: utils.DecodeBool(m["artworkIsUrl"], false),
			AutoCorrespondence: utils.DecodeBool(m["autoCorrespondence"], true),
			CompanyName: utils.DecodeString(m["companyName"], ""),
			Correspondence: DecodePrintingOrderCorrespondenceArray(m["correspondence"]),
			CustomerAddress: DecodeCustomerAddress(m["customerAddress"]),
			CustomerEmail: utils.DecodeString(m["customerEmail"], ""),
			CustomerName: utils.DecodeString(m["customerName"], ""),
			Description: utils.DecodeString(m["description"], ""),
			EmailCode: utils.DecodeString(m["emailCode"], ""),
			Files: DecodePrintingOrderFiles(m["files"]),
			IsAirtableOrder: utils.DecodeBool(m["isAirtableOrder"], false),
			ItemIndex: utils.DecodeInt(m["itemIndex"], 0),
			Kind: utils.DecodeString(m["_kind"], "PrintingOrder"),
			Notes: utils.DecodeString(m["notes"], ""),
			NumberOfUnits: utils.DecodeInt(m["numberOfUnits"], 0),
			OrderNumber: utils.DecodeString(m["orderNumber"], ""),
			OrderStatus: utils.DecodeString(m["orderStatus"], ""),
			ProductOptions: DecodeProductOptionArray(m["productOptions"]),
			Quantity: utils.DecodeInt(m["quantity"], 0),
			QuantityOrdered: utils.DecodeInt(m["quantityOrdered"], 0),
			Rejected: utils.DecodeBool(m["rejected"], false),
			RejectionReason: utils.DecodeString(m["rejectionReason"], ""),
			Sku: utils.DecodeString(m["sku"], ""),
			Status: PrintingOrderStatus(utils.DecodeString(m["status"], "new")),
			StatusIndex: utils.DecodeInt(m["statusIndex"], 0),
			StatusLog: DecodePrintingOrderStatusLogArray(m["statusLog"]),
			SupplierContactName: utils.DecodeString(m["supplierContactName"], ""),
			SupplierId: utils.DecodeString(m["supplierId"], ""),
			SupplierName: utils.DecodeString(m["supplierName"], ""),
			SupplierReference: utils.DecodeString(m["supplierReference"], ""),
			TrackingUrl: utils.DecodeString(m["trackingUrl"], ""),
			TrackingUrlApproved: utils.DecodeBool(m["trackingUrlApproved"], false),
		}
	}

	return PrintingOrder{}
}

// DecodeAll converts a slice of maps to a slice of PrintingOrder struct
func DecodeAll(ms interface{}) []PrintingOrder {
	var entities []PrintingOrder

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

// DecodePrintingOrderCorrespondenceArray converts a map to PrintingOrderCorrespondence struct
func DecodePrintingOrderCorrespondenceArray(m interface{}) []PrintingOrderCorrespondence {
	if m == nil {
		return []PrintingOrderCorrespondence{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []PrintingOrderCorrespondence
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, PrintingOrderCorrespondence{
				Id: utils.DecodeString(v["id"], ""),
				Key: PrintingOrderCorrespondenceTemplateKey(utils.DecodeString(v["key"], "CUSTOMER_APPROVE_SAMPLE")),
				Sent: utils.DecodeBool(v["sent"], false),
				SentAt: utils.DecodeInt64(v["sentAt"], 0),
				StatusLogIndex: utils.DecodeInt(v["statusLogIndex"], 0),
				Timestamp: utils.DecodeInt64(v["timestamp"], 0),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []PrintingOrderCorrespondence
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, PrintingOrderCorrespondence{
					Id: utils.DecodeString(val["id"], ""),
					Key: PrintingOrderCorrespondenceTemplateKey(utils.DecodeString(val["key"], "CUSTOMER_APPROVE_SAMPLE")),
					Sent: utils.DecodeBool(val["sent"], false),
					SentAt: utils.DecodeInt64(val["sentAt"], 0),
					StatusLogIndex: utils.DecodeInt(val["statusLogIndex"], 0),
					Timestamp: utils.DecodeInt64(val["timestamp"], 0),
				})
			}
		}
		return entities
	}
	return []PrintingOrderCorrespondence{}
}


// DecodeCustomerAddress converts a map to CustomerAddress struct
func DecodeCustomerAddress(m interface{}) CustomerAddress {
	if m == nil {
		return CustomerAddress{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return CustomerAddress{
			City: utils.DecodeString(val["city"], ""),
			Country: utils.DecodeString(val["country"], "NL"),
			Line: utils.DecodeString(val["line"], ""),
			PostalCode: utils.DecodeString(val["postalCode"], ""),
		}
	}
	return CustomerAddress{}
}


// DecodePrintingOrderFiles converts a map to PrintingOrderFiles struct
func DecodePrintingOrderFiles(m interface{}) PrintingOrderFiles {
	if m == nil {
		return PrintingOrderFiles{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return PrintingOrderFiles{
			Artwork: media.Decode(val["artwork"]),
			Media: media.Decode(val["media"]),
			OrderConfirmation: media.Decode(val["orderConfirmation"]),
		}
	}
	return PrintingOrderFiles{}
}


// DecodeProductOptionArray converts a map to ProductOption struct
func DecodeProductOptionArray(m interface{}) []ProductOption {
	if m == nil {
		return []ProductOption{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []ProductOption
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, ProductOption{
				Id: utils.DecodeString(v["id"], ""),
				Key: utils.DecodeString(v["key"], ""),
				Value: utils.DecodeString(v["value"], ""),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []ProductOption
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, ProductOption{
					Id: utils.DecodeString(val["id"], ""),
					Key: utils.DecodeString(val["key"], ""),
					Value: utils.DecodeString(val["value"], ""),
				})
			}
		}
		return entities
	}
	return []ProductOption{}
}


// DecodePrintingOrderStatusLogArray converts a map to PrintingOrderStatusLog struct
func DecodePrintingOrderStatusLogArray(m interface{}) []PrintingOrderStatusLog {
	if m == nil {
		return []PrintingOrderStatusLog{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []PrintingOrderStatusLog
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, PrintingOrderStatusLog{
				PrevStatus: PrintingOrderStatus(utils.DecodeString(v["prevStatus"], "new")),
				Status: PrintingOrderStatus(utils.DecodeString(v["status"], "new")),
				Timestamp: utils.DecodeInt64(v["timestamp"], 0),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []PrintingOrderStatusLog
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, PrintingOrderStatusLog{
					PrevStatus: PrintingOrderStatus(utils.DecodeString(val["prevStatus"], "new")),
					Status: PrintingOrderStatus(utils.DecodeString(val["status"], "new")),
					Timestamp: utils.DecodeInt64(val["timestamp"], 0),
				})
			}
		}
		return entities
	}
	return []PrintingOrderStatusLog{}
}


