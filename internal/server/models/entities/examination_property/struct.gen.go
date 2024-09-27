package examination_property
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type ExaminationProperty struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The acceptance region of the property
	AcceptanceRegion AcceptanceRegion `json:"acceptanceRegion"`
	// A description of the property to be examined
	Description string `json:"description"`
	// Whether the property needs to be examined
	Examine bool `json:"examine"`
	// All additional information about the property
	Info map[string]interface{} `json:"info"`
	// The input type of the property
	InputType ExaminationInputType `json:"inputType"`
	// A set of instructions for the examiner of the property
	Instructions string `json:"instructions"`
	// The key of the property to be examined
	Key string `json:"key"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The name of the property to be examined
	Name string `json:"name"`
	// The ID of the product to which the property belongs
	ProductId string `json:"productId"`
	// Whether the property is product specific
	ProductSpecific bool `json:"productSpecific"`
	// Whether it is required to observe the property
	Required bool `json:"required"`
	// Whether the property is testable
	Testable bool `json:"testable"`
	// The type of the property to be examined
	Type string `json:"type"`
	// The unit type of the property to be examined
	UnitType string `json:"unitType"`
}



type AcceptanceRegion struct {
	// The maximum value of the acceptance region
	Max float64 `json:"max"`
	// The minimum value of the acceptance region
	Min float64 `json:"min"`
	// Whether the acceptance region is symmetric
	Symmetric bool `json:"symmetric"`
	// The type of the acceptance region
	Type string `json:"type"`
	// The value of the acceptance region
	Value float64 `json:"value"`
}



type ExaminationInputType string




