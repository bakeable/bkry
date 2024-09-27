package examination_task
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type ExaminationTask struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Whether the task is approved
	Approved bool `json:"approved"`
	// The person assigned to the task
	AssignedTo ExaminationTaskRole `json:"assignedTo"`
	// The change log entries for the task
	ChangeLog []ChangeLogEntry `json:"changeLog"`
	// The kind of the object
	Kind string `json:"kind"`
	// The observations for the task
	Observations map[string]interface{} `json:"observations"`
	// The ID of the product
	ProductId string `json:"productId"`
	// The name of the product
	ProductName string `json:"productName"`
	// The properties of the task
	Properties []Property `json:"properties"`
	// The SKU of the product
	Sku string `json:"sku"`
	// The status of the task
	Status ExaminationTaskStatus `json:"status"`
	// The ordering of the status
	StatusIndex int `json:"statusIndex"`
}



type ChangeLogEntry struct {
	// The action of the change log entry
	Action ExaminationTaskAction `json:"action"`
	// The status of the change log entry
	Status string `json:"status"`
	// The timestamp of the change log entry
	Timestamp string `json:"timestamp"`
	// The user of the change log entry
	User User `json:"user"`
}
type User struct {
	// The email of the user
	Email string `json:"email"`
	// The ID of the user
	ID string `json:"id"`
}



type ExaminationTaskAction string

type Property struct {
	// The acceptance region of the property
	AcceptanceRegion AcceptanceRegion `json:"acceptanceRegion"`
	// The description of the property
	Description string `json:"description"`
	// The expected value of the property
	ExpectedValue string `json:"expectedValue"`
	// The input type of the property
	InputType string `json:"inputType"`
	// The instructions of the property
	Instructions string `json:"instructions"`
	// The key of the property
	Key string `json:"key"`
	// The lower bound of the property
	LowerBound float64 `json:"lowerBound"`
	// The name of the property
	Name string `json:"name"`
	// Whether the property is product specific
	ProductSpecific bool `json:"productSpecific"`
	// Whether the property is required
	Required bool `json:"required"`
	// The type of the property
	Type string `json:"type"`
	// The unit type of the property
	UnitType string `json:"unitType"`
	// The upper bound of the property
	UpperBound float64 `json:"upperBound"`
}
type AcceptanceRegion struct {
	// The maximum value of the acceptance region
	Max float64 `json:"max"`
	// The minimum value of the acceptance region
	Min float64 `json:"min"`
	// Whether the acceptance region is symmetric
	Symmetric bool `json:"symmetric"`
	// The type of the acceptance region
	Type AcceptanceRegionType `json:"type"`
	// The value of the acceptance region
	Value float64 `json:"value"`
}

type AcceptanceRegionType string




type ExaminationTaskRole string
type ExaminationTaskStatus string




