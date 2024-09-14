package entities

var airlineSettingsFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "'AirlineSettings'",
	},
	{
		FieldName: "ClicheCosts",
		FieldType: "float64",
		JsonName: "clicheCosts",
		Description: "The cliche costs for the airline settings",
		DefaultValue: "295.0",
	},
	{
		FieldName: "MetersPerUnit",
		FieldType: "int",
		JsonName: "metersPerUnit",
		Description: "The per units value for the airline settings",
		DefaultValue: "1000",
	},
	{
		FieldName: "RollLength",
		FieldType: "int",
		JsonName: "rollLength",
		Description: "The roll length for the airline settings",
		DefaultValue: "1000",
	},
	{
		FieldName: "RollWidthOptions",
		FieldType: "[]int",
		JsonName: "rollWidthOptions",
		Description: "The roll width options",
		DefaultValue: "[]int{175, 205, 210, 215, 220, 230, 265, 275, 300}",
	},
	{
		FieldName: "RollWidthDividerMap",
		FieldType: "map[int]int",
		JsonName: "rollWidthDividerMap",
		Description: "The roll width to divider map",
		DefaultValue: "map[int]int{175:5, 205:4, 210:4, 215:4, 220:4, 230:3, 265:3, 275:3, 300:2}",
	},
	{
		FieldName: "SwitchCosts",
		FieldType: "float64",
		JsonName: "switchCosts",
		Description: "The switch costs for the airline settings",
		DefaultValue: "200.0",
	},
	{
		FieldName: "SleeveCosts",
		FieldType: "float64",
		JsonName: "sleeveCosts",
		Description: "The sleeve costs for the airline settings",
		DefaultValue: "650.0",
	},
	{
		FieldName: "UnitOptionSets",
		FieldType: "[][]int",
		JsonName: "unitOptionSets",
		Description: "The unit options for the airline pricing",
		DefaultValue: "[][]int{{10, 15, 20, 25, 30, 35, 45, 60, 80, 100, 125, 150, 185, 215, 250, 300, 400, 500},{8, 12, 16, 20, 24, 32, 40, 52, 64, 84, 104, 128, 152, 192, 240, 300, 400, 500},{9, 12, 15, 18, 24, 30, 39, 48, 60, 75, 99, 126, 150, 198, 249, 300, 400, 500},{4, 6, 8, 10, 12, 16, 20, 28, 36, 48, 68, 88, 116, 152, 200, 300, 400, 500}}",
	},
}
