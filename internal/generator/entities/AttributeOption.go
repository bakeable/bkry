package entities

var attributeOptionFields = []FieldConfig{
	{
		FieldName: "Key",
		FieldType: "string",
		JsonName: "key",
		Description: "The option's key",
		DefaultValue: "",
	},
	{
		FieldName: "Min",
		FieldType: "float64",
		JsonName: "min",
		Description: "The minimum value for the option",
		DefaultValue: "0",
	},
	{
		FieldName: "Max",
		FieldType: "float64",
		JsonName: "max",
		Description: "The maximum value for the option",
		DefaultValue: "0",
	},
	{
		FieldName: "Value",
		FieldType: "string",
		JsonName: "value",
		Description: "The option's value",
		DefaultValue: "",
	},
}