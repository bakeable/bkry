package packing_slip

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/dimension"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to PackingSlip struct
func Decode(m interface{}) PackingSlip {
	if m, ok := m.(map[string]interface{}); ok {
		return PackingSlip{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AddedBy: utils.DecodeString(m["addedBy"], ""),
			Address: DecodeAddress(m["address"]),
			Administration: utils.DecodeString(m["administration"], ""),
			Awb: utils.DecodeString(m["awb"], ""),
			CarrierCode: CarrierCode(utils.DecodeString(m["carrierCode"], "BRT")),
			CombinedOrderNumbers: utils.DecodeStringArray(m["combinedOrderNumbers"], []string{}),
			CombinedPackingSlipIds: utils.DecodeStringArray(m["combinedPackingSlipIds"], []string{}),
			CompanyName: utils.DecodeString(m["companyName"], ""),
			Customer: DecodeCustomer(m["customer"]),
			Date: utils.DecodeString(m["date"], ""),
			Delivery: DecodeDelivery(m["delivery"]),
			Description: utils.DecodeString(m["description"], ""),
			IncotermCode: utils.DecodeString(m["incotermCode"], ""),
			Kind: utils.DecodeString(m["_kind"], "PackingSlip"),
			MainPackingSlipId: utils.DecodeString(m["mainPackingSlipId"], ""),
			MarkedAsCompleted: utils.DecodeBool(m["markedAsCompleted"], false),
			MarkedAsPrinted: utils.DecodeBool(m["markedAsPrinted"], false),
			MarkedAsPushed: utils.DecodeBool(m["markedAsPushed"], false),
			Note: utils.DecodeString(m["note"], ""),
			OrderCreatedBy: utils.DecodeString(m["orderCreatedBy"], ""),
			OrderNumber: utils.DecodeString(m["orderNumber"], ""),
			OrderProcessing: utils.DecodeString(m["orderProcessing"], ""),
			OriginalTotalGrossWeight: utils.DecodeFloat64(m["originalTotalGrossWeight"], 0.0),
			Packages: DecodePackageArray(m["packages"]),
			PromisedDeliveryDate: utils.DecodeString(m["promisedDeliveryDate"], ""),
			Reference: utils.DecodeString(m["reference"], ""),
			Service: utils.DecodeString(m["service"], ""),
			SlipNumber: utils.DecodeString(m["slipNumber"], ""),
			Status: PackingSlipStatus(utils.DecodeString(m["status"], "shipment_not_created")),
			StatusIndex: utils.DecodeInt(m["statusIndex"], 0),
			SupplierContactName: utils.DecodeString(m["supplierContactName"], ""),
			SupplierID: utils.DecodeString(m["supplierId"], ""),
			SupplierName: utils.DecodeString(m["supplierName"], ""),
			TotalAmountExcludingVAT: utils.DecodeFloat64(m["totalAmountExcludingVAT"], 0.0),
			TotalGrossWeight: utils.DecodeFloat64(m["totalGrossWeight"], 0.0),
			TotalQuantity: utils.DecodeInt(m["totalQuantity"], 0),
			TrackingUrl: utils.DecodeString(m["trackingUrl"], ""),
		}
	}

	return PackingSlip{}
}

// DecodeAll converts a slice of maps to a slice of PackingSlip struct
func DecodeAll(ms interface{}) []PackingSlip {
	var entities []PackingSlip

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

// DecodeAddress converts a map to Address struct
func DecodeAddress(m interface{}) Address {
	if m == nil {
		return Address{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Address{
			City: utils.DecodeString(val["city"], ""),
			Country: utils.DecodeString(val["country"], ""),
			Description: utils.DecodeString(val["description"], ""),
			Formatted: utils.DecodeString(val["formatted"], ""),
			HouseNumber: utils.DecodeString(val["houseNumber"], ""),
			PostalCode: utils.DecodeString(val["postalCode"], ""),
			Street: utils.DecodeString(val["street"], ""),
		}
	}
	return Address{}
}


// DecodeCustomer converts a map to Customer struct
func DecodeCustomer(m interface{}) Customer {
	if m == nil {
		return Customer{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Customer{
			Email: utils.DecodeString(val["email"], ""),
			Name: utils.DecodeString(val["name"], ""),
			PhoneNumber: utils.DecodeString(val["phoneNumber"], ""),
		}
	}
	return Customer{}
}


// DecodeDelivery converts a map to Delivery struct
func DecodeDelivery(m interface{}) Delivery {
	if m == nil {
		return Delivery{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Delivery{
			AfterTime: DecodeAfterTime(val["afterTime"]),
			BeforeTime: DecodeBeforeTime(val["beforeTime"]),
			Date: utils.DecodeString(val["date"], ""),
			Notes: utils.DecodeString(val["notes"], ""),
		}
	}
	return Delivery{}
}

// DecodeAfterTime converts a map to AfterTime struct
func DecodeAfterTime(m interface{}) AfterTime {
	if m == nil {
		return AfterTime{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return AfterTime{
			Hour: utils.DecodeInt(val["hour"], 9),
			Minute: utils.DecodeInt(val["minute"], 0),
		}
	}
	return AfterTime{}
}


// DecodeBeforeTime converts a map to BeforeTime struct
func DecodeBeforeTime(m interface{}) BeforeTime {
	if m == nil {
		return BeforeTime{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return BeforeTime{
			Hour: utils.DecodeInt(val["hour"], 17),
			Minute: utils.DecodeInt(val["minute"], 0),
		}
	}
	return BeforeTime{}
}



// DecodePackageArray converts a map to Package struct
func DecodePackageArray(m interface{}) []Package {
	if m == nil {
		return []Package{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []Package
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, Package{
				Amount: utils.DecodeInt(v["amount"], 0),
				Dimensions: DecodeDimensions(v["dimensions"]),
				TranssmartCode: utils.DecodeString(v["transsmartCode"], ""),
				Type: utils.DecodeString(v["type"], ""),
				Weight: dimension.Decode(v["weight"]),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []Package
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, Package{
					Amount: utils.DecodeInt(val["amount"], 0),
					Dimensions: DecodeDimensions(val["dimensions"]),
					TranssmartCode: utils.DecodeString(val["transsmartCode"], ""),
					Type: utils.DecodeString(val["type"], ""),
					Weight: dimension.Decode(val["weight"]),
				})
			}
		}
		return entities
	}
	return []Package{}
}

// DecodeDimensions converts a map to Dimensions struct
func DecodeDimensions(m interface{}) Dimensions {
	if m == nil {
		return Dimensions{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Dimensions{
			Height: dimension.Decode(val["height"]),
			Length: dimension.Decode(val["length"]),
			Width: dimension.Decode(val["width"]),
		}
	}
	return Dimensions{}
}



