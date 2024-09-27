package quotation
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Quotation struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The customer's information
	Customer Customer `json:"customer"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The quotation's added products
	LineItems []LineItem `json:"lineItems"`
	// The quotation's name
	Name string `json:"name"`
	// Any notes made on the quotation
	Notes string `json:"notes"`
	// A supplied referral code
	ReferralCode string `json:"referralCode"`
	// The quotation's status
	Status QuotationStatus `json:"status"`
	// The ordering index of the status
	StatusIndex int `json:"statusIndex"`
	// The quotation totals
	Totals Totals `json:"totals"`
}



type Customer struct {
	// The customer's address
	Address string `json:"address"`
	// The customer's email
	Email string `json:"email"`
	// The customer's name
	Name string `json:"name"`
	// The customer's phone number
	Phone string `json:"phone"`
}


type LineItem struct {
	// The quotation's attributes
	Attributes map[string]string `json:"attributes"`
	// The final price of the quotation
	FinalPrice float64 `json:"finalPrice"`
	// The fixed costs of the quotation
	FixedCosts float64 `json:"fixedCosts"`
	// The margin amount of the quotation
	MarginAmount float64 `json:"marginAmount"`
	// The margin percentage of the quotation
	MarginPercentage float64 `json:"marginPercentage"`
	// The products information
	Product Product `json:"product"`
	// The total costs of the quotation
	TotalCosts float64 `json:"totalCosts"`
	// The number of units in the quotation
	Units int `json:"units"`
	// The variable costs of the quotation
	VariableCosts float64 `json:"variableCosts"`
}
type Product struct {
	// The product's id
	ID string `json:"id"`
	// The product's name
	Name string `json:"name"`
	// The product's name
	Sku string `json:"sku"`
}




type Totals struct {
	// The final price of the quotation
	FinalPrice float64 `json:"finalPrice"`
	// The margin amount of the quotation
	MarginAmount float64 `json:"marginAmount"`
	// The margin percentage of the quotation
	MarginPercentage float64 `json:"marginPercentage"`
	// The total costs of the quotation
	TotalCosts float64 `json:"totalCosts"`
}



type QuotationStatus string




