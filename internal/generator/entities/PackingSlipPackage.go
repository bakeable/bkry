package entities

var packingSlipPackageFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the datastore entity",
		DefaultValue: "PackingSlipPackage",
	},
	{
		FieldName: "Name",
		FieldType: "string",
		JsonName: "name",
		Description: "The name of the package",
		DefaultValue: "",
	},
	{
		FieldName: "Type",
		FieldType: "string",
		JsonName: "type",
		Description: "The type of the package",
		DefaultValue: "",
	},
	{
		FieldName: "TranssmartCode",
		FieldType: "TranssmartCode",
		JsonName: "transsmartCode",
		Description: "The code of the package in Transsmart",
		IsEnum: true,
		EnumValues: []string{"BOX", "BLOKPALLET", "EUROGESTAPELD", "EUROPALLET", "PALLETOVERIG", "WEGWERPGES", "WEGWERPPALLET"},
		DefaultValue: "BOX",
	},
	{
		FieldName: "Weight",
		FieldType: "dimension.Dimension",
		JsonName: "weight",
		Description: "The weight of the package",
	},
	{
		FieldName: "Dimensions",
		FieldType: "Dimensions",
		JsonName: "dimensions",
		Description: "The dimensions of the package",
		IsChildStruct: true,
		Fields: []FieldConfig{
			{
				FieldName: "Height",
				FieldType: "dimension.Dimension",
				JsonName: "height",
				Description: "The height of the package",
			},
			{
				FieldName: "Length",
				FieldType: "dimension.Dimension",
				JsonName: "length",
				Description: "The length of the package",
			},
			{
				FieldName: "Width",
				FieldType: "dimension.Dimension",
				JsonName: "width",
				Description: "The width of the package",
			},
		},
	},
}
