package examination_settings
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type ExaminationSettings struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Standard timeframe from which to select tasks (in milliseconds)
	ExclusionTime int64 `json:"exclusionTime"`
	// Exclusion time frame
	ExclusionTimeFrame ExclusionTimeFrame `json:"exclusionTimeFrame"`
	// The kind of the object
	Kind string `json:"kind"`
	// Maximum number of open tasks
	MaxOpenTasks int `json:"maxOpenTasks"`
	// Maximum number of delivery entries to examine
	MaxProducts int `json:"maxProducts"`
	// Maximum number of properties to examine
	MaxProperties int `json:"maxProperties"`
	// Minimum number of delivery entries to examine
	MinProducts int `json:"minProducts"`
	// Minimum number of properties to examine
	MinProperties int `json:"minProperties"`
	// Properties to examine
	Properties []string `json:"properties"`
	// Timeframe from which to select delivery entries (in milliseconds)
	SelectionTime int64 `json:"selectionTime"`
	// Selection time frame
	SelectionTimeFrame SelectionTimeFrame `json:"selectionTimeFrame"`
}



type ExclusionTimeFrame struct {
	// Number of days in the exclusion time frame
	Days int `json:"days"`
	// Number of hours in the exclusion time frame
	Hours int `json:"hours"`
}


type SelectionTimeFrame struct {
	// Number of days in the selection time frame
	Days int `json:"days"`
	// Number of hours in the selection time frame
	Hours int `json:"hours"`
}







