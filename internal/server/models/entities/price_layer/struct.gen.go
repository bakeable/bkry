package price_layer
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type PriceLayer struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The attributes associated with the variable price
	Attributes []string `json:"attributes"`
	// The description of the price layer
	Description string `json:"description"`
	// Whether the price layer is external
	External bool `json:"external"`
	// Whether the price layer should be included in the margin calculation
	IncludeInMarginCalculation bool `json:"includeInMarginCalculation"`
	// Whether the price layer should be included in the purchase price
	IncludeInPurchasePrice bool `json:"includeInPurchasePrice"`
	// The type of incremental pricing of the price layer
	IncrementalPricing AttributeIncrementalPricing `json:"incrementalPricing"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The name of the price layer
	Name string `json:"name"`
	// The options associated with the price layer
	Options []string `json:"options"`
	// The unit size of the variable price
	PerUnits int `json:"perUnits"`
	// The pricing type of the price layer
	PriceType PriceType `json:"priceType"`
	// How to round the unit sizes if they do not correspond to the unit size
	RoundingUnits RoundingUnits `json:"roundingUnits"`
	// Whether to save the price-layer to be used later
	SaveAsNew bool `json:"saveAsNew"`
	// The template of the price layer
	Template string `json:"template"`
	// The values of the variable price
	VariablePriceValues []VariablePriceValues `json:"values"`
}



type VariablePriceValues struct {
	// The attributes associated with the variable price
	Attributes map[string]interface{} `json:"attributes"`
	// The value of the variable price
	VariablePriceValue VariablePriceValue `json:"value"`
}
type VariablePriceValue struct {
	// The decimal value of the variable price
	DecimalValue float64 `json:"decimalValue"`
	// The float value of the variable price
	FloatValue float64 `json:"floatValue"`
	// The integer value of the variable price
	IntValue int `json:"intValue"`
	// The text value of the variable price
	TextValue string `json:"textValue"`
}





type AttributeIncrementalPricing string
type PriceType string
type RoundingUnits string




