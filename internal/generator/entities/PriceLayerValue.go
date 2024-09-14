package entities

var priceLayerValueFields = []FieldConfig{
	{
		FieldName: "Attributes",
		FieldType: "map[string]string",
		JsonName: "attributes",
		Description: "The attributes that define the price layer value entry",
		DefaultValue: "map[string]string{}",
	},
}
