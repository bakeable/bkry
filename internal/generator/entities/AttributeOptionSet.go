package entities

var attributeOptionSetFields = []FieldConfig{
	{
		FieldName: "Label",
		FieldType: "string",
		JsonName: "label",
		Description: "The optionset's label",
		DefaultValue: "",
	},
	{
		FieldName: "Options",
		FieldType: "attribute_option.AttributeOption",
		JsonName: "options",
		IsArray: true,
		Description: "The optionset's options",
		DefaultValue: "[]attribute_option.AttributeOption",
	},
}
