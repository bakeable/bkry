package entities

var cloudFunctionFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "CloudFunction",
	},
	{
		FieldName: "Name",
		FieldType: "string",
		JsonName: "name",
		Description: "The cloud function's name",
		DefaultValue: "",
	},
	{
		FieldName: "Description",
		FieldType: "string",
		JsonName: "description",
		Description: "The cloud function's description",
		DefaultValue: "",
	},
	{
		FieldName: "Region",
		FieldType: "string",
		JsonName: "region",
		Description: "The cloud function's region",
		DefaultValue: "",
	},
	{
		FieldName: "TriggerType",
		FieldType: "string",
		JsonName: "triggerType",
		Description: "The cloud function's trigger type",
		DefaultValue: "",
	},
	{
		FieldName: "TriggerValue",
		FieldType: "string",
		JsonName: "triggerValue",
		Description: "The cloud function's trigger value",
		DefaultValue: "",
	},
	{
		FieldName: "Runtime",
		FieldType: "string",
		JsonName: "runtime",
		Description: "The cloud function's runtime",
		DefaultValue: "",
	},
}
