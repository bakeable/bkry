package margin

type Margin struct {
	// The amount of margin
	Amount float64 `json:"amount"`
	// The percentage of margin
	Percentage float64 `json:"percentage"`
	// The type of percentage
	PercentageType string `json:"percentageType"`
	// Whether the margin is staggered
	Staggered bool `json:"staggered"`
	// The type of margin
	Type string `json:"type"`
	// The values of the margin
	Values []MarginValues `json:"values"`
}

type MarginValues struct {
	// The attributes of the margin
	Attributes map[string]interface{} `json:"attributes"`
	// The value of the margin
	Value MarginValue `json:"value"`
}

type MarginValue struct {
	// The text value of the margin
	TextValue string `json:"textValue"`
	// The integer value of the margin
	IntValue int `json:"intValue"`
	// The float value of the margin
	FloatValue float64 `json:"floatValue"`
	// The decimal value of the margin
	DecimalValue string `json:"decimalValue"`
}