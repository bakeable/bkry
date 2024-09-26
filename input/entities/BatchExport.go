package entities

var exportFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "Export",
	},
	{
		FieldName: "ExportFromTimestamp",
		FieldType: "int64",
		JsonName: "exportFromTimestamp",
		Description: "The timestamp from which the export should start",
		DefaultValue: "0",
	},
	{
		FieldName: "Ids",
		FieldType: "string",
		JsonName: "ids",
		Description: "The IDs of the objects to export",
		IsArray: true,
	},
}
