package airline_order_group
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type AirlineOrderGroup struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The AFAS order numbers of the airline orders in the group
	AfasOrderNumbers []string `json:"afasOrderNumbers"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The name of the airline order group
	Name string `json:"name"`
	// The list of airline orders within the group
	Orders []airline_order.AirlineOrder `json:"orders"`
	// The pricing tables
	PricingTables []airline_pricing.AirlinePricing `json:"pricingTables"`
	// The airline settings
	Settings airline_settings.AirlineSettings `json:"settings"`
}








