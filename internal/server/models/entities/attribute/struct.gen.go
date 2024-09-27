package attribute
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Attribute struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Whether an attribute is selected by default
	Default bool `json:"default"`
	// The default options for this attribute
	DefaultOptions []DefaultOption `json:"defaultOptions"`
	// A description of the attribute
	Description string `json:"description"`
	// Whether the attribute is included
	Included bool `json:"included"`
	// Whether graduated pricing applies
	IncrementalPricing AttributeIncrementalPricing `json:"incrementalPricing"`
	// The attribute's key
	Key string `json:"key"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The maximum option ID
	MaxOptionID int `json:"maxOptionId"`
	// The attribute's name
	Name string `json:"name"`
	// The attribute's numerical identifier
	Nid int `json:"nid"`
	// The attribute's option sets
	OptionSets []attribute_option_set.AttributeOptionSet `json:"optionSets"`
	// Whether the attribute is optional
	Optional bool `json:"optional"`
	// All options within the attribute
	Options []attribute_option.AttributeOption `json:"options"`
	// A numerical representation of the order in which the attribute should be displayed
	Order int `json:"order"`
	// The attribute's prefix
	Prefix string `json:"prefix"`
	// The associated price types
	PriceTypes []AttributePriceType `json:"priceTypes"`
	// The attribute's surcharges
	Surcharges map[string]interface{} `json:"surcharges"`
	Type AttributeType `json:"type"`
}



type DefaultOption struct {
	// Unique identifier for the option
	ID string `json:"id"`
	// The option's key
	Key string `json:"key"`
	// The maximum value for the option
	Max float64 `json:"max"`
	// The minimum value for the option
	Min float64 `json:"min"`
	// The option's value
	Value string `json:"value"`
}



type AttributeIncrementalPricing string
type AttributePriceType string
type AttributeType string




