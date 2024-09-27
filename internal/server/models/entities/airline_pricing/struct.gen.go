package airline_pricing
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type AirlinePricing struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The number of colors for the airline pricing
	Colors []int `json:"colors"`
	// The description of the airline pricing
	Description string `json:"description"`
	// The number of dividers for the airline pricing
	Dividers int `json:"dividers"`
	// The kind of the object
	Kind string `json:"kind"`
	// The pricing table's margin
	Margin Margin `json:"margin"`
	// An array of units for the airline pricing
	Units []int `json:"units"`
	// An array of values for the airline pricing
	Values []AirlinePricingValue `json:"values"`
	// The width of the airline pricing
	Width float64 `json:"width"`
	// The type of width for the airline pricing
	WidthType string `json:"widthType"`
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


type AirlinePricingValue struct {
	// The attributes for the airline pricing
	PricingAttributes PricingAttributes `json:"attributes"`
	// The value for the airline pricing
	PricingValue PricingValue `json:"value"`
}
type PricingAttributes struct {
	// The colors for the airline pricing
	Colors int `json:"colors"`
	// The units for the airline pricing
	Units int `json:"units"`
}


type PricingValue struct {
	// The decimal value for the airline pricing
	DecimalValue float64 `json:"decimalValue"`
	// The float value for the airline pricing
	FloatValue float64 `json:"floatValue"`
	// The int value for the airline pricing
	IntValue int `json:"intValue"`
	// The text value for the airline pricing
	TextValue string `json:"textValue"`
}









