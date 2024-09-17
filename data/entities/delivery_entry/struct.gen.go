package delivery_entry
import (
	"github.com/bakeable/bkry/data/entities/examination_task"
	"github.com/bakeable/bkry/data/entities/product"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type DeliveryEntry struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Customer ID associated with the delivery entry
	CustomerId string `json:"customerId"`
	// Name of the customer associated with the delivery entry
	CustomerName string `json:"customerName"`
	// Date of the delivery entry
	Date string `json:"date"`
	// Description of the delivery entry
	Description string `json:"description"`
	// Additional description for the delivery entry
	Description2 string `json:"description_2"`
	// The product's examination
	Examination Examination `json:"examination"`
	// The examination tasks associated with the delivery entry
	ExaminationTasks []examination_task.ExaminationTask `json:"examinationTasks"`
	// Extra description for the delivery entry
	ExtraDescription string `json:"extraDescription"`
	// Globally unique identifier for the delivery entry
	Guid string `json:"guid"`
	// Order number for the delivery entry
	OrderNumber string `json:"orderNumber"`
	// Type of order for the delivery entry
	OrderType string `json:"orderType"`
	// The product information
	Product product.Product `json:"product"`
	// Serial number for the delivery entry
	SerialNumber int `json:"serialNumber"`
	// SKU (Stock Keeping Unit) for the delivery entry
	Sku string `json:"sku"`
	// Type of unit for the delivery entry
	UnitType string `json:"unitType"`
	// Number of units in the delivery entry
	Units int `json:"units"`
	// Number of units per pallet for the delivery entry
	UnitsPerPallet int `json:"unitsPerPallet"`
	// Warehouse associated with the delivery entry
	Warehouse string `json:"warehouse"`
	// The kind of the object
	_Kind string `json:"_kind"`
}



type Examination struct {
	// Whether the examination is a priority
	Priority bool `json:"priority"`
	// The priority score of the examination
	PriorityScore int `json:"priorityScore"`
	// External properties of the examination
	Properties map[string]interface{} `json:"properties"`
	// The timeout of the examination
	Timeout int `json:"timeout"`
}







