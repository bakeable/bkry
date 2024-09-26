package entities

var apiKeyFields = []FieldConfig{
	{
		FieldName: "Kind",
		FieldType: "string",
		JsonName: "_kind",
		Description: "The kind of the object",
		DefaultValue: "ApiKey",
	},
	{
		FieldName: "publicKey",
		FieldType: "string",
		JsonName: "publicKey",
		Description: "The public key",
		DefaultValue: "",
	},
	{
		FieldName: "secretKey",
		FieldType: "string",
		JsonName: "secretKey",
		Description: "The secret key",
		DefaultValue: "",
	},
	{
		FieldName: "userId",
		FieldType: "string",
		JsonName: "userId",
		Description: "The user ID",
		DefaultValue: "",
	},
}
