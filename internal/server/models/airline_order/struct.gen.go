package airline_order
import "github.com/bakeable/bkry/data/general/auditinfo"

type AirlineOrder struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The AFAS order number of the airline order
	AfasOrderNumber string `json:"afasOrderNumber"`
	// The number of colors of the airline order
	Colors int `json:"colors"`
	// The ID of the customer associated with the airline order
	CustomerId string `json:"customerId"`
	// The name of the customer associated with the airline order
	CustomerName string `json:"customerName"`
	// The date of the airline order
	Date string `json:"date"`
	// The delivery address of the order
	DeliveryAddress string `json:"deliveryAddress"`
	// The description of the airline order
	Description string `json:"description"`
	// The number of dividers of the airline order
	Dividers int `json:"dividers"`
	// The unique identifier of the airline order
	Guid string `json:"guid"`
	// The ID of the supplier associated with the airline order
	Inkooprelatie string `json:"supplierId"`
	// Whether the airline order has already been grouped
	IsGrouped bool `json:"isGrouped"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The order margin
	Margin Margin `json:"margin"`
	// The number of the airline order
	Nummer string `json:"number"`
	// The order number of the airline order
	OrderNumber string `json:"orderNumber"`
	// The original total costs (when separately ordered)
	OriginalTotalCosts int `json:"originalTotalCosts"`
	// The original unit costs (when separately ordered)
	OriginalUnitCosts int `json:"originalUnitCosts"`
	// The serial number of the airline order
	SerialNumber int `json:"serialNumber"`
	// The SKU (Stock Keeping Unit) of the airline order
	Sku string `json:"sku"`
	// The status of the airline order
	Status string `json:"status"`
	// The total costs of the airline order
	TotalCosts float64 `json:"totalCosts"`
	// The unit costs of the airline order
	UnitCosts float64 `json:"unitCosts"`
	// The unit type of the airline order
	UnitType string `json:"unitType"`
	// The number of units in the airline order
	Units int `json:"units"`
	// The warehouse of the airline order
	Warehouse string `json:"warehouse"`
	// The width of the ordered airline foil
	Width int `json:"width"`
	// The width unit type of the ordered airline foil
	WidthUnitType int `json:"widthUnitType"`
}



type Margin struct {
	// The amount of margin
	Amount int `json:"amount"`
	// The percentage of margin
	Percentage int `json:"percentage"`
	// The type of percentage
	PercentageType string `json:"percentageType"`
	// Whether the margin is staggered
	Staggered bool `json:"staggered"`
	// The type of margin
	Type string `json:"type"`
	// The values of the margin
	Values []interface{} `json:"values"`
}







