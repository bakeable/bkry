package packing_slip_package
import (
	"github.com/bakeable/bkry/data/general/dimension"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type PackingSlipPackage struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The dimensions of the package
	Dimensions Dimensions `json:"dimensions"`
	// The kind of the datastore entity
	Kind string `json:"_kind"`
	// The name of the package
	Name string `json:"name"`
	// The code of the package in Transsmart
	TranssmartCode TranssmartCode `json:"transsmartCode"`
	// The type of the package
	Type string `json:"type"`
	// The weight of the package
	Weight dimension.Dimension `json:"weight"`
}



type Dimensions struct {
	// The height of the package
	Height dimension.Dimension `json:"height"`
	// The length of the package
	Length dimension.Dimension `json:"length"`
	// The width of the package
	Width dimension.Dimension `json:"width"`
}



type TranssmartCode string




