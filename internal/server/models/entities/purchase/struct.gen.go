package purchase
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Purchase struct {
	// The unique identifier for the Purchase
	ID string `json:"id"`
	// The audit information for the modification of the Purchase
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the Purchase
	Created auditinfo.AuditInfo `json:"created"`
	// The billing address for the purchase
	BillingAddress Address `json:"billing_address"`
	// The Id of the package being purchased
	PackageID string `json:"package_id"`
	// The method of payment
	PaymentMethod string `json:"payment_method"`
	// The price details of the purchase
	Price Price `json:"price"`
	// The status of the purchase
	Status Status `json:"status"`
	// The log of status changes for the purchase
	StatusLog []StatusLog `json:"status_log"`
	// The Id of the transaction
	TransactionID string `json:"transaction_id"`
	// The Id of the user making the purchase
	UserID string `json:"user_id"`
}




type Address struct {
	// The city of the billing address
	City string `json:"city"`
	// The country code of the billing address
	CountryCode string `json:"country_code"`
	// The first line of the billing address
	Line1 string `json:"line1"`
	// The second line of the billing address
	Line2 string `json:"line2"`
	// The postal code of the billing address
	PostalCode string `json:"postal_code"`
	// The state of the billing address
	State string `json:"state"`
	// The street of the billing address
	Street string `json:"street"`
}

type Price struct {
	// The currency of the purchase
	Currency string `json:"currency"`
	// The discount amount
	DiscountAmount float64 `json:"discount_amount"`
	// The final amount after tax and discounts
	FinalAmount float64 `json:"final_amount"`
	// The initial amount before tax and discounts
	InitialAmount float64 `json:"initial_amount"`
	// The tax details of the purchase
	Tax []Tax `json:"tax"`
	// The total tax amount
	TaxAmount float64 `json:"tax_amount"`
}

type Tax struct {
	// The tax amount
	Amount string `json:"amount"`
	// The tax code
	Code string `json:"code"`
	// The tax label
	Label string `json:"label"`
	// The tax percentage
	Percentage float64 `json:"percentage"`
}

type StatusLog struct {
	// The new status of the purchase
	NewStatus string `json:"new_status"`
	// The previous status of the purchase
	PrevStatus string `json:"prev_status"`
	// The timestamp when the status was changed
	Timestamp int64 `json:"timestamp"`
}



