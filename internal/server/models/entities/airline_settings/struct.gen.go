package airline_settings
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type AirlineSettings struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The cliche costs for the airline settings
	ClicheCosts float64 `json:"clicheCosts"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The per units value for the airline settings
	MetersPerUnit int `json:"metersPerUnit"`
	// The roll length for the airline settings
	RollLength int `json:"rollLength"`
	// The roll width to divider map
	RollWidthDividerMap map[int]int `json:"rollWidthDividerMap"`
	// The roll width options
	RollWidthOptions []int `json:"rollWidthOptions"`
	// The sleeve costs for the airline settings
	SleeveCosts float64 `json:"sleeveCosts"`
	// The switch costs for the airline settings
	SwitchCosts float64 `json:"switchCosts"`
	// The unit options for the airline pricing
	UnitOptionSets [][]int `json:"unitOptionSets"`
}








