package price_configuration
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type PriceConfiguration struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The additional costs
	AdditionalCosts float64 `json:"additional_costs"`
	// The attribute configuration for this price configuration
	Attributes map[string]string `json:"attributes"`
	// The currency used in the response
	Currency string `json:"currency"`
	// The final price
	FinalPrice float64 `json:"final_price"`
	// The fixed costs
	FixedCosts float64 `json:"fixed_costs"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The margin amount
	MarginAmount float64 `json:"margin_amount"`
	// The margin percentage
	MarginPercentage float64 `json:"margin_percentage"`
	// The priceConfiguration's corresponding product ID
	ProductId string `json:"product_id"`
	// The priceConfiguration's corresponding product SKU
	Sku string `json:"sku"`
	// The total costs
	TotalCosts float64 `json:"total_costs"`
	// The costs associated with a single unit
	UnitCosts float64 `json:"unit_costs"`
	// The costs associated with a single unit
	UnitMargin float64 `json:"unit_margin"`
	// The number of units
	Units int `json:"units"`
	// The costs associated with a single unit
	VariableCosts float64 `json:"variable_costs"`
}








