package quotation

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to Quotation struct
func Decode(m interface{}) Quotation {
	if m, ok := m.(map[string]interface{}); ok {
		return Quotation{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Customer: DecodeCustomer(m["customer"]),
			Kind: utils.DecodeString(m["_kind"], "Quotation"),
			LineItems: DecodeLineItemArray(m["lineItems"]),
			Name: utils.DecodeString(m["name"], ""),
			Notes: utils.DecodeString(m["notes"], ""),
			ReferralCode: utils.DecodeString(m["referralCode"], ""),
			Status: QuotationStatus(utils.DecodeString(m["status"], "concept")),
			StatusIndex: utils.DecodeInt(m["statusIndex"], 0),
			Totals: DecodeTotals(m["totals"]),
		}
	}

	return Quotation{}
}

// DecodeAll converts a slice of maps to a slice of Quotation struct
func DecodeAll(ms interface{}) []Quotation {
	var entities []Quotation

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

// DecodeCustomer converts a map to Customer struct
func DecodeCustomer(m interface{}) Customer {
	if m == nil {
		return Customer{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Customer{
			Address: utils.DecodeString(val["address"], ""),
			Email: utils.DecodeString(val["email"], ""),
			Name: utils.DecodeString(val["name"], ""),
			Phone: utils.DecodeString(val["phone"], ""),
		}
	}
	return Customer{}
}


// DecodeLineItemArray converts a map to LineItem struct
func DecodeLineItemArray(m interface{}) []LineItem {
	if m == nil {
		return []LineItem{}
	}
	if vArr, ok := m.([]map[string]interface{}); ok {
		var entities []LineItem
		for _, v := range vArr {
			if v == nil {
				continue
			}
			entities = append(entities, LineItem{
				Attributes: utils.DecodeStringMap(v["attributes"], map[string]string{}),
				FinalPrice: utils.DecodeFloat64(v["finalPrice"], 0.0),
				FixedCosts: utils.DecodeFloat64(v["fixedCosts"], 0.0),
				MarginAmount: utils.DecodeFloat64(v["marginAmount"], 0.0),
				MarginPercentage: utils.DecodeFloat64(v["marginPercentage"], 0.0),
				Product: DecodeProduct(v["product"]),
				TotalCosts: utils.DecodeFloat64(v["totalCosts"], 0.0),
				Units: utils.DecodeInt(v["units"], 0),
				VariableCosts: utils.DecodeFloat64(v["variableCosts"], 0.0),
			})
		}
		return entities
	}
	if vArr, ok := m.([]interface{}); ok {
		var entities []LineItem
		for _, v := range vArr {
			if v == nil {
				continue
			}
			if val, ok := v.(map[string]interface{}); ok {
				entities = append(entities, LineItem{
					Attributes: utils.DecodeStringMap(val["attributes"], map[string]string{}),
					FinalPrice: utils.DecodeFloat64(val["finalPrice"], 0.0),
					FixedCosts: utils.DecodeFloat64(val["fixedCosts"], 0.0),
					MarginAmount: utils.DecodeFloat64(val["marginAmount"], 0.0),
					MarginPercentage: utils.DecodeFloat64(val["marginPercentage"], 0.0),
					Product: DecodeProduct(val["product"]),
					TotalCosts: utils.DecodeFloat64(val["totalCosts"], 0.0),
					Units: utils.DecodeInt(val["units"], 0),
					VariableCosts: utils.DecodeFloat64(val["variableCosts"], 0.0),
				})
			}
		}
		return entities
	}
	return []LineItem{}
}

// DecodeProduct converts a map to Product struct
func DecodeProduct(m interface{}) Product {
	if m == nil {
		return Product{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Product{
			ID: utils.DecodeString(val["id"], ""),
			Name: utils.DecodeString(val["name"], ""),
			Sku: utils.DecodeString(val["sku"], ""),
		}
	}
	return Product{}
}



// DecodeTotals converts a map to Totals struct
func DecodeTotals(m interface{}) Totals {
	if m == nil {
		return Totals{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Totals{
			FinalPrice: utils.DecodeFloat64(val["finalPrice"], 0.0),
			MarginAmount: utils.DecodeFloat64(val["marginAmount"], 0.0),
			MarginPercentage: utils.DecodeFloat64(val["marginPercentage"], 0.0),
			TotalCosts: utils.DecodeFloat64(val["totalCosts"], 0.0),
		}
	}
	return Totals{}
}


