package entities

var deliveryEntryFields = []FieldConfig{
	{
		FieldName: "_Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "DeliveryEntry",
	},
	{
		FieldName: "Date",
		FieldType: "string",
		JsonName: "date",
		Description: "Date of the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "OrderType",
		FieldType: "string",
		JsonName: "orderType",
		Description: "Type of order for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "OrderNumber",
		FieldType: "string",
		JsonName: "orderNumber",
		Description: "Order number for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "SerialNumber",
		FieldType: "int",
		JsonName: "serialNumber",
		Description: "Serial number for the delivery entry",
		DefaultValue: "0",
	},
	{
		FieldName: "Guid",
		FieldType: "string",
		JsonName: "guid",
		Description: "Globally unique identifier for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "CustomerId",
		FieldType: "string",
		JsonName: "customerId",
		Description: "Customer ID associated with the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "CustomerName",
		FieldType: "string",
		JsonName: "customerName",
		Description: "Name of the customer associated with the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "Examination",
		FieldType: "Examination",
		JsonName: "examination",
		Description: "The product's examination",
		IsChildStruct: true,
		Fields: []FieldConfig{
			{
				FieldName: "PriorityScore",
				FieldType: "int",
				JsonName: "priorityScore",
				Description: "The priority score of the examination",

			},
			{
				FieldName: "Priority",
				FieldType: "bool",
				JsonName: "priority",
				Description: "Whether the examination is a priority",
			},
			{
				FieldName: "Timeout",
				FieldType: "int",
				JsonName: "timeout",
				Description: "The timeout of the examination",
			},
			{
				FieldName: "Properties",
				FieldType: "map[string]interface{}",
				JsonName: "properties",
				Description: "External properties of the examination",
			},
		},
	},
	{
		FieldName: "Sku",
		FieldType: "string",
		JsonName: "sku",
		Description: "SKU (Stock Keeping Unit) for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "Description",
		FieldType: "string",
		JsonName: "description",
		Description: "Description of the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "ExaminationTasks",
		FieldType: "examination_task.ExaminationTask",
		JsonName: "examinationTasks",
		Description: "The examination tasks associated with the delivery entry",
		IsArray: true,
	},
	{
		FieldName: "ExtraDescription",
		FieldType: "string",
		JsonName: "extraDescription",
		Description: "Extra description for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "Product",
		FieldType: "product.Product",
		JsonName: "product",
		Description: "The product information",
	},
	{
		FieldName: "UnitType",
		FieldType: "string",
		JsonName: "unitType",
		Description: "Type of unit for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "Description2",
		FieldType: "string",
		JsonName: "description_2",
		Description: "Additional description for the delivery entry",
		DefaultValue: "",
	},
	{
		FieldName: "Units",
		FieldType: "int",
		JsonName: "units",
		Description: "Number of units in the delivery entry",
		DefaultValue: "0",
	},
	{
		FieldName: "UnitsPerPallet",
		FieldType: "int",
		JsonName: "unitsPerPallet",
		Description: "Number of units per pallet for the delivery entry",
		DefaultValue: "0",
	},
	{
		FieldName: "Warehouse",
		FieldType: "string",
		JsonName: "warehouse",
		Description: "Warehouse associated with the delivery entry",
		DefaultValue: "",
	},
}
