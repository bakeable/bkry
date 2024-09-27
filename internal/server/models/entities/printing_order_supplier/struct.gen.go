package printing_order_supplier
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type PrintingOrderSupplier struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The supplier code
	Code string `json:"code"`
	// Name of the supplier
	ContactName string `json:"contactName"`
	// The supplier delivery type
	DeliveryType PrintingOrderSupplierDeliveryType `json:"deliveryType"`
	// The main email of the supplier
	Email string `json:"email"`
	// The other email addresses of the supplier
	Emails []string `json:"emails"`
	// The kind of the entity
	Kind string `json:"_kind"`
	// The name of the supplier
	Name string `json:"name"`
	// The sku's associated with the supplier
	Skus []string `json:"skus"`
}




type PrintingOrderSupplierDeliveryType string




