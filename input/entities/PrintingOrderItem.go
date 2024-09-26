package entities

var printingOrderItemFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the entity",
		DefaultValue: "PrintingOrderItem",
	},
	{
		FieldName: "SupplierId",
		FieldType: "string",
		JsonName: "supplierId",
		Description: "The supplier ID",
		DefaultValue: "",
	},
	{
		FieldName: "SupplierName",
		FieldType: "string",
		JsonName: "supplierName",
		Description: "The supplier name",
		DefaultValue: "",
	},
	{
		FieldName: "SupplierContactName",
		FieldType: "string",
		JsonName: "supplierContactName",
		Description: "The supplier contact name",
		DefaultValue: "",
	},
	{
		FieldName: "Sku",
		FieldType: "string",
		JsonName: "sku",
		Description: "The sku associated with the supplier",
		DefaultValue: "",
	},
}
