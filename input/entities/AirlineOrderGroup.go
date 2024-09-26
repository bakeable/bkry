package entities

var airlineOrderGroupFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "AirlineOrderGroup",
	},
	{
		FieldName: "Name",
		FieldType: "string",
		JsonName: "name",
		Description: "The name of the airline order group",
		DefaultValue: "",
	},
	{
		FieldName: "Orders",
		FieldType: "airline_order.AirlineOrder",
		JsonName: "orders",
		Description: "The list of airline orders within the group",
		IsArray: true,
	},
	{
		FieldName: "PricingTables",
		FieldType: "airline_pricing.AirlinePricing",
		JsonName: "pricingTables",
		Description: "The pricing tables",
		IsArray: true,
	},
	{
		FieldName: "Settings",
		FieldType: "airline_settings.AirlineSettings",
		JsonName: "settings",
		Description: "The airline settings",
	},
	{
		FieldName: "AfasOrderNumbers",
		FieldType: "string",
		JsonName: "afasOrderNumbers",
		Description: "The AFAS order numbers of the airline orders in the group",
		DefaultValue: "[]string{}",
		IsArray: true,
	},
}
