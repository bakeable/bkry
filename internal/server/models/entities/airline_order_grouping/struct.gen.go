package airline_order_grouping
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type AirlineOrderGrouping struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The returned order number from AFAS
	AfasOrderNumbers []string `json:"afasOrderNumbers"`
	// The final price of this order grouping
	FinalPrice float64 `json:"finalPrice"`
	// The date the first order was placed
	FirstOrderDate string `json:"firstOrderDate"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The total amount of margin
	MarginAmount float64 `json:"marginAmount"`
	// The total amount of margin
	MarginPercentage float64 `json:"marginPercentage"`
	// A reference name for the airline order grouping
	Name string `json:"name"`
	// The date the order grouping needs to be ordered
	OrderDate string `json:"orderDate"`
	// The list of order groups within the grouping
	OrderGroups []airline_order_group.AirlineOrderGroup `json:"orderGroups"`
	// The list of airline orders within the grouping
	OrderIds []string `json:"orderIds"`
	// The original costs when using separated orders
	OriginalTotalCosts float64 `json:"originalTotalCosts"`
	// The total profit amount of this order grouping
	ProfitAmount float64 `json:"profitAmount"`
	// The total profit percentage of this order grouping
	ProfitPercentage float64 `json:"profitPercentage"`
	// The costs saved when using this order grouping
	SavedCosts float64 `json:"savedCosts"`
	// Whether the order grouping has been sent to AFAS
	SentToAfas bool `json:"sentToAfas"`
	// The timestamp of when the order grouping was sent to AFAS
	SentToAfasTimestamp int64 `json:"sentToAfasTimestamp"`
	// The total costs of this order grouping
	TotalCosts float64 `json:"totalCosts"`
	// The total number of units of the airline order grouping
	Units int `json:"units"`
}








