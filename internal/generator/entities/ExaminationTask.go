package entities

var examinationTaskFields = []FieldConfig{
	{
		FieldName:    "Kind",
		FieldType:    "string",
		JsonName:     "kind",
		Description:  "The kind of the object",
		DefaultValue: "ExaminationTask",
	},
	{
		FieldName: "Approved",
		FieldType: "bool",
		JsonName:  "approved",
		Description: "Whether the task is approved",
	},
	{
		FieldName: "AssignedTo",
		FieldType: "ExaminationTaskRole",
		IsEnum: true,
		EnumValues: []string{"admin", "controller", "observer"},
		DefaultValue: "observer",
		JsonName:  "assignedTo",
		Description: "The person assigned to the task",
	},
	{
		FieldName: "ChangeLog",
		FieldType: "ChangeLogEntry",
		JsonName:  "changeLog",
		Description: "The change log entries for the task",
		IsArray: true,
		IsChildStruct: true,
		Fields: changeLogEntryFields,
	},
	{
		FieldName: "Observations",
		FieldType: "map[string]interface{}",
		JsonName:  "observations",
		Description: "The observations for the task",
	},
	{
		FieldName: "ProductId",
		FieldType: "string",
		JsonName:  "productId",
		Description: "The ID of the product",
	},
	{
		FieldName: "ProductName",
		FieldType: "string",
		JsonName:  "productName",
		Description: "The name of the product",
	},
	{
		FieldName: "Properties",
		FieldType: "Property",
		JsonName:  "properties",
		Description: "The properties of the task",
		IsArray: true,
		IsChildStruct: true,
		Fields: propertyFields,
	},
	{
		FieldName: "Sku",
		FieldType: "string",
		JsonName:  "sku",
		Description: "The SKU of the product",
	},
	{
		FieldName: "Status",
		FieldType: "ExaminationTaskStatus",
		IsEnum: true,
		EnumValues: []string{"concept", "pending", "failed", "success", "done", "archived"},
		DefaultValue: "concept",
		JsonName:  "status",
		Description: "The status of the task",
	},
	{
		FieldName: "StatusIndex",
		FieldType: "int",
		DefaultValue: "0",
		JsonName:  "statusIndex",
		Description: "The ordering of the status",
	},
}

var changeLogEntryFields = []FieldConfig{
	{
		FieldName:   "Action",
		FieldType:   "ExaminationTaskAction",
		IsEnum:      true,
		EnumValues:  []string{"created", "updated", "observed", "solved", "approved", "archived"},
		DefaultValue: "created",
		JsonName:    "action",
		Description: "The action of the change log entry",
	},
	{
		FieldName:   "Status",
		FieldType:   "string",
		JsonName:    "status",
		Description: "The status of the change log entry",
	},
	{
		FieldName:   "Timestamp",
		FieldType:   "string",
		JsonName:    "timestamp",
		Description: "The timestamp of the change log entry",
	},
	{
		FieldName:   "User",
		FieldType:   "User",
		JsonName:    "user",
		Description: "The user of the change log entry",
		IsChildStruct: true,
		Fields: examinationTaskUserFields,
	},
}

var examinationTaskUserFields = []FieldConfig{
	{
		FieldName:   "Email",
		FieldType:   "string",
		JsonName:    "email",
		Description: "The email of the user",
	},
	{
		FieldName:   "ID",
		FieldType:   "string",
		JsonName:    "id",
		Description: "The ID of the user",
	},
}

var propertyFields = []FieldConfig{
	{
		FieldName:   "AcceptanceRegion",
		FieldType:   "AcceptanceRegion",
		JsonName:    "acceptanceRegion",
		Description: "The acceptance region of the property",
		IsChildStruct: true,
		Fields: acceptanceRegionFields,
	},
	{
		FieldName:   "Description",
		FieldType:   "string",
		JsonName:    "description",
		Description: "The description of the property",
	},
	{
		FieldName:   "ExpectedValue",
		FieldType:   "string",
		JsonName:    "expectedValue",
		Description: "The expected value of the property",
	},
	{
		FieldName:   "InputType",
		FieldType:   "string",
		JsonName:    "inputType",
		Description: "The input type of the property",
	},
	{
		FieldName:   "Instructions",
		FieldType:   "string",
		JsonName:    "instructions",
		Description: "The instructions of the property",
	},
	{
		FieldName:   "Key",
		FieldType:   "string",
		JsonName:    "key",
		Description: "The key of the property",
	},
	{
		FieldName:   "LowerBound",
		FieldType:   "float64",
		JsonName:    "lowerBound",
		Description: "The lower bound of the property",
	},
	{
		FieldName:   "Name",
		FieldType:   "string",
		JsonName:    "name",
		Description: "The name of the property",
	},
	{
		FieldName:   "ProductSpecific",
		FieldType:   "bool",
		JsonName:    "productSpecific",
		Description: "Whether the property is product specific",
	},
	{
		FieldName:   "Required",
		FieldType:   "bool",
		JsonName:    "required",
		Description: "Whether the property is required",
	},
	{
		FieldName:   "Type",
		FieldType:   "string",
		JsonName:    "type",
		Description: "The type of the property",
	},
	{
		FieldName:   "UnitType",
		FieldType:   "string",
		JsonName:    "unitType",
		Description: "The unit type of the property",
	},
	{
		FieldName:   "UpperBound",
		FieldType:   "float64",
		JsonName:    "upperBound",
		Description: "The upper bound of the property",
	},
}
		

var acceptanceRegionFields = []FieldConfig{
	{
		FieldName:   "Min",
		FieldType:   "float64",
		JsonName:    "min",
		Description: "The minimum value of the acceptance region",
	},
	{
		FieldName:   "Max",
		FieldType:   "float64",
		JsonName:    "max",
		Description: "The maximum value of the acceptance region",
	},
	{
		FieldName:   "Symmetric",
		FieldType:   "bool",
		JsonName:    "symmetric",
		Description: "Whether the acceptance region is symmetric",
	},
	{
		FieldName:   "Type",
		FieldType:   "AcceptanceRegionType",
		JsonName:    "type",
		IsEnum: 	true,
		EnumValues: []string{"range", "percentage"},
		DefaultValue: "percentage",
		Description: "The type of the acceptance region",
	},
	{
		FieldName:   "Value",
		FieldType:   "float64",
		JsonName:    "value",
		Description: "The value of the acceptance region",
	},
}
