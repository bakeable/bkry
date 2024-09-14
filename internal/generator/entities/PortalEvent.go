package entities

var portalEventFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The entity kind",
		DefaultValue: "PortalEvent",
	},
	{
		FieldName: "Timestamp",
		FieldType: "int",
		JsonName: "timestamp",
		Description: "The timestamp of the event",
	},
	{
		FieldName: "Timezone",
		FieldType: "int",
		JsonName: "timezone",
		Description: "The timezone of the event",
	},
	{
		FieldName: "UserId",
		FieldType: "string",
		JsonName: "userId",
		Description: "The user triggering the event",
	},
	{
		FieldName: "Type",
		FieldType: "PortalEventType",
		JsonName: "type",
		IsEnum: true,
		EnumValues: []string{
			"Created", "Updated", "Deleted", 
			"New", 
			"Push", "Print", "PrintLabel", 
			"EmailSent", "EmailReceived", 
			"Error",
			"StatusUpdate",
		},
		Description: "The type of the event",
		DefaultValue: "Update",
	},
	{
		FieldName: "Metadata",
		FieldType: "map[string]interface{}",
		JsonName: "metadata",
		Description: "The metadata of the event",
		DefaultValue: "{}",
	},
	{
		FieldName: "EntityName",
		FieldType: "string",
		JsonName: "entityName",
		Description: "The entity name",
	},
	{
		FieldName: "EntityID",
		FieldType: "string",
		JsonName: "entityId",
		Description: "The entity ID",
	},
}
