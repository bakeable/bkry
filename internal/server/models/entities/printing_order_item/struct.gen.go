package printing_order_item
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type PrintingOrderItem struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The kind of the entity
	Kind string `json:"_kind"`
	// The sku associated with the supplier
	Sku string `json:"sku"`
	// The supplier contact name
	SupplierContactName string `json:"supplierContactName"`
	// The supplier ID
	SupplierId string `json:"supplierId"`
	// The supplier name
	SupplierName string `json:"supplierName"`
}








