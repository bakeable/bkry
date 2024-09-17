package packing_slip
import (
	"github.com/bakeable/bkry/data/general/dimension"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type PackingSlip struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Name of the person who added this packing slip
	AddedBy string `json:"addedBy"`
	// Recipient's address
	Address Address `json:"address"`
	// The administration related to this packing slip
	Administration string `json:"administration"`
	// Airway bill number
	Awb string `json:"awb"`
	// Carrier code for the shipment
	CarrierCode CarrierCode `json:"carrierCode"`
	// List of combined order numbers
	CombinedOrderNumbers []string `json:"combinedOrderNumbers"`
	// List of combined packing slip IDs
	CombinedPackingSlipIds []string `json:"combinedPackingSlipIds"`
	// Name of the company
	CompanyName string `json:"companyName"`
	// Customer details
	Customer Customer `json:"customer"`
	// Date of the packing slip
	Date string `json:"date"`
	// Delivery details
	Delivery Delivery `json:"delivery"`
	// Description of the packing slip
	Description string `json:"description"`
	// Incoterm code for the shipment
	IncotermCode string `json:"incotermCode"`
	// Type of the packing slip
	Kind string `json:"_kind"`
	// ID of the main packing slip
	MainPackingSlipId string `json:"mainPackingSlipId"`
	// Flag indicating if the packing slip is marked as completed
	MarkedAsCompleted bool `json:"markedAsCompleted"`
	// Flag indicating if the packing slip is marked as printed
	MarkedAsPrinted bool `json:"markedAsPrinted"`
	// Flag indicating if the packing slip is marked as pushed
	MarkedAsPushed bool `json:"markedAsPushed"`
	// Additional note for the packing slip
	Note string `json:"note"`
	// Name of the person who created the order
	OrderCreatedBy string `json:"orderCreatedBy"`
	// Order number associated with the packing slip
	OrderNumber string `json:"orderNumber"`
	// Order processing details
	OrderProcessing string `json:"orderProcessing"`
	// Original total gross weight of the shipment
	OriginalTotalGrossWeight float64 `json:"originalTotalGrossWeight"`
	// List of packages in the shipment
	Packages []Package `json:"packages"`
	// Promised delivery date
	PromisedDeliveryDate string `json:"promisedDeliveryDate"`
	// Reference for the packing slip
	Reference string `json:"reference"`
	// Service details
	Service string `json:"service"`
	// Packing slip number
	SlipNumber string `json:"slipNumber"`
	// Status of the packing slip
	Status PackingSlipStatus `json:"status"`
	// The ordering index of the status
	StatusIndex int `json:"statusIndex"`
	// Name of the supplier contact
	SupplierContactName string `json:"supplierContactName"`
	// ID of the supplier
	SupplierID string `json:"supplierId"`
	// Name of the supplier
	SupplierName string `json:"supplierName"`
	// Total amount excluding VAT
	TotalAmountExcludingVAT float64 `json:"totalAmountExcludingVAT"`
	// Total gross weight
	TotalGrossWeight float64 `json:"totalGrossWeight"`
	// Total quantity
	TotalQuantity int `json:"totalQuantity"`
	// Tracking URL
	TrackingUrl string `json:"trackingUrl"`
}



type Address struct {
	// City of the address
	City string `json:"city"`
	// Country of the address
	Country string `json:"country"`
	// Description of the address
	Description string `json:"description"`
	// Formatted address string
	Formatted string `json:"formatted"`
	// House number of the address
	HouseNumber string `json:"houseNumber"`
	// Postal code of the address
	PostalCode string `json:"postalCode"`
	// Street of the address
	Street string `json:"street"`
}


type Customer struct {
	// Email of the customer
	Email string `json:"email"`
	// Name of the customer
	Name string `json:"name"`
	// Phone number of the customer
	PhoneNumber string `json:"phoneNumber"`
}


type Delivery struct {
	// Time after which the delivery is expected
	AfterTime AfterTime `json:"afterTime"`
	// Time before which the delivery is expected
	BeforeTime BeforeTime `json:"beforeTime"`
	// Date of the delivery
	Date string `json:"date"`
	// Notes about the delivery
	Notes string `json:"notes"`
}
type AfterTime struct {
	// Hour of the delivery time
	Hour int `json:"hour"`
	// Minute of the delivery time
	Minute int `json:"minute"`
}


type BeforeTime struct {
	// Hour of the delivery time
	Hour int `json:"hour"`
	// Minute of the delivery time
	Minute int `json:"minute"`
}




type Package struct {
	// Number of packages
	Amount int `json:"amount"`
	// Dimensions of the package
	Dimensions Dimensions `json:"dimensions"`
	// Transsmart code for the package
	TranssmartCode string `json:"transsmartCode"`
	// Type of the package
	Type string `json:"type"`
	// Weight of the package
	Weight dimension.Dimension `json:"weight"`
}
type Dimensions struct {
	// Height of the package
	Height dimension.Dimension `json:"height"`
	// Length of the package
	Length dimension.Dimension `json:"length"`
	// Width of the package
	Width dimension.Dimension `json:"width"`
}





type CarrierCode string
type PackingSlipStatus string




