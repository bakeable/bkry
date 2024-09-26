package printing_order
import (
	"github.com/bakeable/bkry/data/entities/media"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type PrintingOrder struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Flag indicating if the printing order has been approved
	Approved bool `json:"approved"`
	// The timestamp of the approved
	ApprovedTimestamp int64 `json:"approvedTimestamp"`
	// Flag indicating whether the artwork is approved
	ArtworkApproved bool `json:"artworkApproved"`
	// Flag indicating whether the artwork is a URL or a file
	ArtworkIsUrl bool `json:"artworkIsUrl"`
	// Flag indicating whether to automatically send correspondence
	AutoCorrespondence bool `json:"autoCorrespondence"`
	// The name of the company associated with the printing order
	CompanyName string `json:"companyName"`
	// The correspondence tracking
	Correspondence []PrintingOrderCorrespondence `json:"correspondence"`
	// The address of the customer
	CustomerAddress CustomerAddress `json:"customerAddress"`
	// Email of the customer
	CustomerEmail string `json:"customerEmail"`
	// Name of the customer
	CustomerName string `json:"customerName"`
	// Description of the printing order
	Description string `json:"description"`
	// The 10 digit e-mail code
	EmailCode string `json:"emailCode"`
	// The files associated with the printing order
	Files PrintingOrderFiles `json:"files"`
	// Whether the printing order is from Airtable
	IsAirtableOrder bool `json:"isAirtableOrder"`
	// The item index of the printing order
	ItemIndex int `json:"itemIndex"`
	// Type of the printing order
	Kind string `json:"_kind"`
	// Notes of the printing order
	Notes string `json:"notes"`
	// Number of units ordered per quantity
	NumberOfUnits int `json:"numberOfUnits"`
	// The order number of the printing order
	OrderNumber string `json:"orderNumber"`
	// The status of the order in Magento
	OrderStatus string `json:"orderStatus"`
	// The options of the product
	ProductOptions []ProductOption `json:"productOptions"`
	// Quantity of the printing order
	Quantity int `json:"quantity"`
	// Quantity ordered of the printing order
	QuantityOrdered int `json:"quantityOrdered"`
	// Whether the sample was rejected
	Rejected bool `json:"rejected"`
	// The reason for rejection
	RejectionReason string `json:"rejectionReason"`
	// SKU of the printing order
	Sku string `json:"sku"`
	// Status of the printing order
	Status PrintingOrderStatus `json:"status"`
	// The ordering index of the status
	StatusIndex int `json:"statusIndex"`
	// The log of previous statuses
	StatusLog []PrintingOrderStatusLog `json:"statusLog"`
	// Name of the supplier contact associated with the printing order
	SupplierContactName string `json:"supplierContactName"`
	// ID of the supplier associated with the printing order
	SupplierId string `json:"supplierId"`
	// Name of the supplier associated with the printing order
	SupplierName string `json:"supplierName"`
	// The supplier reference
	SupplierReference string `json:"supplierReference"`
	// Tracking URL of the printing order
	TrackingUrl string `json:"trackingUrl"`
	// Whether the tracking URL is approved
	TrackingUrlApproved bool `json:"trackingUrlApproved"`
}



type PrintingOrderCorrespondence struct {
	// ID of the correspondence
	Id string `json:"id"`
	// Key of the correspondence
	Key PrintingOrderCorrespondenceTemplateKey `json:"key"`
	// Flag indicating if the correspondence was sent
	Sent bool `json:"sent"`
	// Timestamp when the correspondence was sent
	SentAt int64 `json:"sentAt"`
	// The index of the status log change that triggered the correspondence
	StatusLogIndex int `json:"statusLogIndex"`
	// Timestamp of the correspondence
	Timestamp int64 `json:"timestamp"`
}

type PrintingOrderCorrespondenceTemplateKey string

type CustomerAddress struct {
	// City of the address
	City string `json:"city"`
	// Country of the address
	Country string `json:"country"`
	// Address line
	Line string `json:"line"`
	// Postal code of the address
	PostalCode string `json:"postalCode"`
}


type PrintingOrderFiles struct {
	// The artwork file
	Artwork media.Media `json:"artwork"`
	// The media file
	Media media.Media `json:"media"`
	// The order confirmation file
	OrderConfirmation media.Media `json:"orderConfirmation"`
}


type ProductOption struct {
	// ID of the product option
	Id string `json:"id"`
	// Key of the product option
	Key string `json:"key"`
	// Value of the product option
	Value string `json:"value"`
}


type PrintingOrderStatusLog struct {
	// Previous status of the printing order
	PrevStatus PrintingOrderStatus `json:"prevStatus"`
	// Status of the printing order
	Status PrintingOrderStatus `json:"status"`
	// Timestamp of the status change
	Timestamp int64 `json:"timestamp"`
}



type PrintingOrderStatus string




