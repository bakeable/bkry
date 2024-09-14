package entities

var printingOrderSupplierFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the entity",
		DefaultValue: "'PrintingOrderSupplier'",
	},
	{
		FieldName: "Code",
		FieldType: "string",
		JsonName: "code",
		Description: "The supplier code",
		DefaultValue: "",
	},
	{
		FieldName: "ContactName",
		FieldType: "string",
		JsonName: "contactName",
		Description: "Name of the supplier",
		DefaultValue: "",
	},
	{
		FieldName: "DeliveryType",
		FieldType: "PrintingOrderSupplierDeliveryType",
		JsonName: "deliveryType",
		Description: "The supplier delivery type",
		IsEnum: true,
		EnumValues: []string{"EMAIL"},
		DefaultValue: "EMAIL",
	},
	{
		FieldName: "Email",
		FieldType: "string",
		JsonName: "email",
		Description: "The main email of the supplier",
		DefaultValue: "",
	},
	{
		FieldName: "Emails",
		FieldType: "string",
		JsonName: "emails",
		IsArray: true,
		Description: "The other email addresses of the supplier",
		
	},
	{
		FieldName: "Name",
		FieldType: "string",
		JsonName: "name",
		Description: "The name of the supplier",
		DefaultValue: "",
	},
	{
		FieldName: "Skus",
		FieldType: "string",
		JsonName: "skus",
		IsArray: true,
		Description: "The sku's associated with the supplier",
		
	},
}
